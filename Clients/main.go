package main

import (
	splitter "Tarea2/Clients/Splitter"
	"Tarea2/DataNode/datanode"
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

func main() {
	// fmt.Println("uwu")
	// splitter.Splitter("cosa.pdf")
	// splitter.Combiner("cosa.pdf", 3)

	// name, err := os.Hostname()
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("hostname:", name)

	var conn *grpc.ClientConn
	conn, err := grpc.Dial("dist141:8080", grpc.WithInsecure())

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

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Ingrese nombre del archivo: ")
	file, _ := reader.ReadString('\n')
	file = strings.TrimSuffix(file, "\n")
	file = strings.TrimSuffix(file, "\r")

	cantidadPartes := splitter.Splitter(file)

	var mensaje datanode.Chunk
	var i uint64
	for i = 0; i < cantidadPartes; i++ {
		nombreParte := file + "_" + strconv.FormatUint(i, 10)
		chunkBytes, errBytes := ioutil.ReadFile(nombreParte)

		if errBytes != nil {
			fmt.Print(errBytes)

		}

		mensaje = datanode.Chunk{
			Content:   chunkBytes,
			ChunkName: nombreParte,
			FileName:  file,
			ChunkPos:  i,
		}

		// fmt.Printf(partes[i])

		if err := response.Send(&mensaje); err != nil {
			log.Fatalf("Failed to send a note: %v", err)
		}

		time.Sleep(1 * time.Second)
		// }
	}

	response.CloseSend()
	<-waitc

}
