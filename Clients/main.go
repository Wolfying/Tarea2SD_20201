package main

import (
	splitter "Tarea2/Clients/Splitter"
	"Tarea2/DataNode/datanode"
	"Tarea2/NameNode/namenode"
	"bufio"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"google.golang.org/grpc"
)

func uploadLibro() bool {
	var conn *grpc.ClientConn
	var reader *bufio.Reader
	conn, err := grpc.Dial("dist141:8080", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Conexion fallida: %s", err)
	}

	defer conn.Close()

	link := datanode.NewDataNodeHandlerClient(conn)

	ctx, err1 := context.WithTimeout(context.Background(), 100*time.Second)

	// preguntar centralizado o distribuido
	var response datanode.DataNodeHandler_UploadFileClient

	exito := 0
	for exito == 0 {
		fmt.Println("1) Funcionamiento algoritmo Centralizado")
		fmt.Println("2) Funcionamiento algoritmo Distribuido")
		reader = bufio.NewReader(os.Stdin)
		fmt.Printf("Ingrese opcion a utilizar: ")
		opcion, _ := reader.ReadString('\n')
		opcion = strings.TrimSuffix(opcion, "\n")
		opcion = strings.TrimSuffix(opcion, "\r")
		if opcion == "1" {
			response, err = link.UploadFile(ctx)
			exito = 1
		} else if opcion == "2" {
			response, err = link.DistUploadFile(ctx)
			exito = 1
		} else {
			fmt.Println("Ingrese opcion una opción válida: ")
		}
	}

	_ = err1
	if err != nil {
		log.Printf("Conexion fallida al conectarse al datanode para mandar archivo: %s", err)
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
	var cantidadPartes uint64
	cantidadPartes = 0
	var file string
	for cantidadPartes == 0 {
		fmt.Printf("Ingrese nombre del archivo: ")
		file, _ = reader.ReadString('\n')
		file = strings.TrimSuffix(file, "\n")
		file = strings.TrimSuffix(file, "\r")

		cantidadPartes = splitter.Splitter(file)
		if cantidadPartes == 0 {
			fmt.Println("El archivo ingresado no existe, ingrese uno nuevo.")
		}
	}
	var mensaje datanode.Chunk
	var i uint64
	for i = 0; i < cantidadPartes; i++ {
		nombreParte := file + "_" + strconv.FormatUint(i, 10)
		chunkBytes, errBytes := ioutil.ReadFile(nombreParte)
		if errBytes != nil {
			fmt.Print(errBytes)

		}
		os.Remove(nombreParte)
		mensaje = datanode.Chunk{
			Content:   chunkBytes,
			ChunkName: nombreParte,
			FileName:  file,
			ChunkPos:  i,
		}
		if err := response.Send(&mensaje); err != nil {
			log.Fatalf("Failed to send a note: %v", err)
		}

		// time.Sleep(1 * time.Second)

	}

	response.CloseSend()
	<-waitc
	return true
}

func retrieveChunkList(parte int64, libro string, datanodeip string) []byte {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(datanodeip, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Conexion fallida: %s", err)
	}

	defer conn.Close()

	link := datanode.NewDataNodeHandlerClient(conn)

	ctx, err1 := context.WithTimeout(context.Background(), 10*time.Second)
	response, err := link.DownloadFile(ctx)
	_ = err1
	if err != nil {
		log.Printf("Conexion fallida: %s", err)
	}

	waitc := make(chan *datanode.ChunkResponse)

	go func() {
		for {
			in, err := response.Recv()
			if err == io.EOF {
				waitc <- in
				close(waitc)
				return
			}

			if err != nil {
				log.Fatalf("Error al recibir la respuesta del estado del archivo: %v", err)
			}
			log.Printf("El server retorna el siguiente mensaje: %v", in.Message)
		}
	}()

	mensaje := datanode.FileInfo{
		FileName: libro,
		ChunkPos: uint64(parte),
	}

	if err := response.Send(&mensaje); err != nil {
		log.Fatalf("Failed to send a note: %v", err)
	}
	response.CloseSend()
	infoparte := <-waitc

	log.Println(infoparte.GetContent())
	return infoparte.Content
}

func descagarLibro() {
	var conn *grpc.ClientConn
	var reader *bufio.Reader
	conn, err := grpc.Dial("dist144:8080", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Conexion fallida: %s", err)
	}

	defer conn.Close()

	link := namenode.NewNameNodeHandlerClient(conn)

	ctx, err1 := context.WithTimeout(context.Background(), 10*time.Second)

	// preguntar centralizado o distribuido
	var response namenode.NameNodeHandler_SolicitarLibrosClient
	response, err = link.SolicitarLibros(ctx)
	_ = err1
	if err != nil {
		log.Printf("Conexion fallida al conectarse al datanode para mandar archivo: %s", err)
	}

	waitc := make(chan []namenode.Propuesta)
	var libros []namenode.Propuesta

	go func() {
		for {
			in, err := response.Recv()
			if err == io.EOF {
				waitc <- libros
				close(waitc)
				return
			}

			if err != nil {
				log.Fatalf("Error al recibir un mensaje: %v", err)
			}
			log.Printf("El server retorna el siguiente mensaje: %v", in.Status)
			libros = append(libros, *in)
		}
	}()

	mensaje := namenode.Peticion{}
	if err := response.Send(&mensaje); err != nil {
		log.Fatalf("Failed to send a note: %v", err)
	}
	response.CloseSend()
	libros = <-waitc

	for pos, libro := range libros {
		nombreLibro := libro.NombreLibro
		log.Printf("%d) %s", pos, nombreLibro)
	}
	exito := 0
	var poslibro int
	for exito == 0 {
		reader = bufio.NewReader(os.Stdin)
		fmt.Printf("Ingrese el Numero del libro a descargar: ")
		opcion, _ := reader.ReadString('\n')
		opcion = strings.TrimSuffix(opcion, "\n")
		opcion = strings.TrimSuffix(opcion, "\r")
		poslibro, _ = strconv.Atoi(opcion)
		if poslibro > len(libros) {
			fmt.Printf("Ingrese opcion válida.")
		} else {
			exito = 1
		}
	}

	libro := libros[poslibro]
	nombreLibro := libro.NombreLibro
	datanode1 := libro.Datanode1
	datanode2 := libro.Datanode2
	datanode3 := libro.Datanode3
	cantidadPartes := libro.CantidadPartes
	var fragments [][]byte
	for _, parte := range datanode1 {
		fragmento := retrieveChunkList(int64(parte), nombreLibro, "dist141:8080")
		// fragments[parte] = fragmento
		fragments = append(fragments, fragmento)
	}

	for _, parte := range datanode2 {
		fragmento := retrieveChunkList(int64(parte), nombreLibro, "dist142:8080")
		// fragments[parte] = fragmento
		fragments = append(fragments, fragmento)
	}

	for _, parte := range datanode3 {
		fragmento := retrieveChunkList(int64(parte), nombreLibro, "dist143:8080")
		// fragments[parte] = fragmento
		fragments = append(fragments, fragmento)
	}

	splitter.Combiner(nombreLibro, uint64(cantidadPartes), fragments)
	log.Printf("El Libro $s ha sido descargado", nombreLibro)

}

func main() {
	// fmt.Println("uwu")
	// splitter.Splitter("cosa.pdf")
	// splitter.Combiner("cosa.pdf", 3)

	// name, err := os.Hostname()
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("hostname:", name)
	exito1 := 0
	var reader *bufio.Reader
	for exito1 == 0 {
		fmt.Println("1) Subir Libro")
		fmt.Println("2) Descargar Libro")
		reader = bufio.NewReader(os.Stdin)
		fmt.Printf("Ingrese opcion: ")
		opcion, _ := reader.ReadString('\n')
		opcion = strings.TrimSuffix(opcion, "\n")
		opcion = strings.TrimSuffix(opcion, "\r")
		if opcion == "1" {
			uploadLibro()
			exito1 = 1
		} else if opcion == "2" {
			// Descargar
			descagarLibro()
			exito1 = 1
		} else {
			fmt.Println("Ingrese opcion una opción válida: ")
		}
	}

}
