package main

import (
	"Tarea2/DataNode/datanode"
	"Tarea2/NameNode/namenode"
	"context"
	"io"
	"log"
	"math"
	"math/rand"
	"net"
	"time"

	"google.golang.org/grpc"
)

type ServerDataNode struct {
	integer int32
}

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

	// return nil
}

func (sdn *ServerDataNode) UploadBook(incomestream datanode.DataNodeHandler_UploadBookServer) error {
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
		//_ = fileName
		file := in.Content

		// if _, err12 := os.Stat("/libros"); os.IsNotExist(err12) {
		// 	errFolder := os.Mkdir("libros", 0755)
		// 	fmt.Printf("Carpeta Creada")
		// 	if errFolder != nil {
		// 		log.Fatal(err)
		// 	}
		// }
		// filePath := "libros/" + nombreParte
		// filePart := ioutil.WriteFile(filePath, file, 0644)

		chunkFiles = append(chunkFiles, chunkFilesList{chunkFile: file, chunkName: nombreParte})

		// if filePart != nil {
		// 	log.Fatal(filePart)
		// }

		/*
			Preparación respuesta a cliente.
		*/
		log.Printf("Archivo %s recibido", nombreParte)

	}

	// Gestionar envío de propuesta.
	cantidadpartes := int32(len(chunkFiles))
	propuesta := GenerarPropuesta(cantidadpartes)

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9443", grpc.WithInsecure())

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

	_ = propuestaNueva

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

func GenerarPropuesta(cantidadpartes int32) *namenode.Propuesta {
	var number int32 = cantidadpartes
	a := makeRange(0, int32(math.Floor(float64(number)/float64(3))))
	b := makeRange(int32(math.Floor(float64(number)/float64(3))), int32(math.Floor(2*float64(number)/float64(3))))
	c := makeRange(int32(math.Floor(2*float64(number)/float64(3))), number)

	return &namenode.Propuesta{
		Datanode1:      a,
		Datanode2:      b,
		Datanode3:      c,
		CantidadPartes: cantidadpartes,
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

// func getOutboundIP() net.IP {
// 	conn, err := net.Dial("udp", "8.8.8.8:80")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer conn.Close()

// 	localAddr := conn.LocalAddr().(*net.UDPAddr)

// 	return localAddr.IP
// }

// func servirServidor(wg *sync.WaitGroup, datanodeServer *datanode.DatanodeServer, puerto string) {
// 	lis, err := net.Listen("tcp", ":"+puerto)
// 	if err != nil {
// 		log.Fatalf("Failed to listen on port %s: %v", puerto, err)
// 	}
// 	grpcServer := grpc.NewServer()

// 	datanode.RegisterDatanodeServiceServer(grpcServer, datanodeServer)

// 	if err := grpcServer.Serve(lis); err != nil {
// 		log.Fatalf("Failed to serve gRPC server over port %s: %v", puerto, err)
// 	}
// }

func main() {

	server := ServerDataNode{}
	puerto := ":9444"
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
