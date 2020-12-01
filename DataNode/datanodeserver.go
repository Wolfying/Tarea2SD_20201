package main

import (
	"Tarea2/DataNode/datanode"
	"Tarea2/NameNode/namenode"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"net"
	"os"
	"time"

	"google.golang.org/grpc"
)

// ServerDataNode ...
type ServerDataNode struct {
	integer int32
}

// Ping ...
func (sdn *ServerDataNode) Ping(incomeping datanode.DataNodeHandler_PingServer) error {
	for {
		in, err := incomeping.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		_ = in
		/*
			Preparación respuesta al ping.
		*/
		randnumber := rand.Intn(100)

		acepta := true
		if randnumber > 90 {
			acepta = false
		}

		if err := incomeping.Send(&datanode.Ping{
			CanReceive: acepta,
			Message:    ""}); err != nil {
			return incomeping.Send(&datanode.Ping{Message: err.Error(),
				CanReceive: false})
		}

	}
}

// UploadBook ...
func (sdn *ServerDataNode) UploadBook(incomestream datanode.DataNodeHandler_UploadBookServer) error {

	for {
		in, err := incomestream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		nombreParte := in.ChunkName
		file := in.Content

		if _, err12 := os.Stat("/libros"); os.IsNotExist(err12) {
			errFolder := os.Mkdir("libros", 0755)
			fmt.Printf("Carpeta Creada")
			if errFolder != nil {
				log.Fatal(err)
			}
		}
		filePath := "libros/" + nombreParte
		filePart := ioutil.WriteFile(filePath, file, 0644)

		if filePart != nil {
			log.Fatal(filePart)
		}

		log.Printf("Archivo %s recibido", nombreParte)
	}
	/*
		Preparación respuesta a cliente.
	*/

	if err := incomestream.Send(&datanode.Response{
		Message: "Todos los archivos recibidos",
		Status:  datanode.StatusCode_Success}); err != nil {
		return err
	}

	return nil
}

type chunkFilesList struct {
	chunkFile []byte
	chunkName string
}

// UploadFile ...
func (sdn *ServerDataNode) UploadFile(incomestream datanode.DataNodeHandler_UploadFileServer) error {
	var fileName string
	var chunkFiles []chunkFilesList

	for {
		in, err := incomestream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		nombreParte := in.ChunkName
		fileName = in.FileName
		file := in.Content

		// if _, err12 := os.Stat("/libros"); os.IsNotExist(err12) {
		// 	errFolder := os.Mkdir("libros", 0755)
		// 	fmt.Printf("Carpeta Creada")
		// 	if errFolder != nil {
		// 		log.Fatal(err)
		// 	}
		// }
		chunkFiles = append(chunkFiles, chunkFilesList{chunkFile: file, chunkName: nombreParte})

		/*
			Preparación respuesta a cliente.
		*/
		log.Printf("Archivo %s recibido", nombreParte)
	}

	// Gestionar envío de propuesta.
	cantidadpartes := int32(len(chunkFiles))
	propuesta := GenerarPropuesta(cantidadpartes, fileName)

	var conn *grpc.ClientConn
	conn, err := grpc.Dial("dist144:8080", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Conexion fallida: %s", err)
	}

	defer conn.Close()

	link := namenode.NewNameNodeHandlerClient(conn)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	response, err := link.ManejarPropuesta(ctx)

	if err != nil {
		log.Printf("Conexion fallida: %s", err)
	}

	waitc := make(chan *namenode.Propuesta)

	go func() {
		for {
			in, err := response.Recv()
			if err == io.EOF {
				close(waitc)
				return
			}

			if err != nil {
				log.Fatalf("Error al recibir un mensaje: %v", err)
			}
			log.Printf("El server retorna el siguiente mensaje: %v", in.Status)
			waitc <- in
		}
	}()

	if err := response.Send(propuesta); err != nil {
		log.Fatalf("Failed to send a note: %v", err)
	}

	propuestaNueva := <-waitc

	datanode1 := propuestaNueva.Datanode1 //datanode actual
	datanode2 := propuestaNueva.Datanode2
	datanode3 := propuestaNueva.Datanode3

	for _, filenumber := range datanode1 {
		// guardar archivos en datanode
		parte := chunkFiles[filenumber]
		filePath := "libros/" + parte.chunkName
		filePart := ioutil.WriteFile(filePath, parte.chunkFile, 0644)

		if filePart != nil {
			log.Fatal(filePart)
		}

	}
	var chunkstosend []chunkFilesList
	for _, filenumber := range datanode2 {
		// mandar archivos a datanode 2
		chunkstosend = append(chunkstosend, chunkFiles[filenumber])
	}
	sendChunkList(chunkstosend, "dist142:8080")
	chunkstosend = []chunkFilesList{}
	for _, filenumber := range datanode3 {
		// mandar archivos a datanode 2
		chunkstosend = append(chunkstosend, chunkFiles[filenumber])
	}
	sendChunkList(chunkstosend, "dist143:8080")
	//

	// _ = propuestaNueva

	//Se envia la respuesta
	if err := incomestream.Send(&datanode.Response{
		Message: fileName,
		Status:  datanode.StatusCode_Success}); err != nil {
		return err
		// return incomestream.Send(&datanode.Response{Message: err.Error(),
		// 	Status: datanode.StatusCode_InternalError})
	}

	return nil
}

func sendChunkList(chunks []chunkFilesList, datanodeip string) *datanode.Response {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(datanodeip, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Conexion fallida: %s", err)
	}

	defer conn.Close()

	link := datanode.NewDataNodeHandlerClient(conn)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	response, err := link.UploadFile(ctx)

	if err != nil {
		log.Printf("Conexion fallida: %s", err)
	}

	waitc := make(chan struct{})

	go func() {
		for {
			in, err := response.Recv()
			if err == io.EOF {
				close(waitc)
				return
			}

			if err != nil {
				log.Fatalf("Error al recibir un mensaje: %v", err)
			}
			log.Printf("El server retorna el siguiente mensaje: %v", in.Message)
		}
	}()

	var mensaje datanode.Chunk
	var i uint64
	for _, filenumber := range chunks {
		nombreParte := filenumber.chunkName
		chunkBytes := filenumber.chunkFile

		mensaje = datanode.Chunk{
			Content:   chunkBytes,
			ChunkName: nombreParte,
			FileName:  nombreParte,
			ChunkPos:  i,
		}

		if err := response.Send(&mensaje); err != nil {
			log.Fatalf("Failed to send a note: %v", err)
		}
	}

	return nil
}

// GenerarPropuesta ...
func GenerarPropuesta(cantidadpartes int32, nombreLibro string) *namenode.Propuesta {
	var number int32 = cantidadpartes
	a := makeRange(0, int32(math.Floor(float64(number)/float64(3))))
	b := makeRange(int32(math.Floor(float64(number)/float64(3))), int32(math.Floor(2*float64(number)/float64(3))))
	c := makeRange(int32(math.Floor(2*float64(number)/float64(3))), number)

	return &namenode.Propuesta{
		Datanode1:      a,
		Datanode2:      b,
		Datanode3:      c,
		CantidadPartes: cantidadpartes,
		Status:         namenode.PropuestaStatus_Aprobado,
		NombreLibro:    nombreLibro,
	}
}

func makeRange(min, max int32) []int32 {
	a := make([]int32, max-min)
	for i := range a {
		a[i] = min + int32(i)
	}
	return a
}

// DownloadFile ...
func (sdn *ServerDataNode) DownloadFile(incomestream datanode.DataNodeHandler_DownloadFileServer) error {

	return nil
}

func main() {

	server := ServerDataNode{}
	puerto := ":8080"
	lis, err := net.Listen("tcp", puerto)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", puerto, err)
	}
	grpcServer := grpc.NewServer()

	datanode.RegisterDataNodeHandlerServer(grpcServer, &server)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port %s: %v", puerto, err)
	}

}
