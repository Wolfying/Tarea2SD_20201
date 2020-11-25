package main

import (
	"Tarea2/DataNode/datanode"
	"context"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	// fmt.Println("uwu")
	// splitter.Splitter("cosa.pdf")
	// splitter.Combiner("cosa.pdf", 3)

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9444", grpc.WithInsecure())

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

	// for i := 0; i < len(partes); i++ {
	// chunkBytes, errBytes := ioutil.ReadFile(partes[i])

	// if errBytes != nil {
	// 	fmt.Print(errBytes)
	// }

	mensaje = datanode.Chunk{
		Content:   nil,
		ChunkName: "Hola putito",
		FileName:  "hola puto",
		ChunkPos:  1,
	}

	// fmt.Printf(partes[i])

	if err := response.Send(&mensaje); err != nil {
		log.Fatalf("Failed to send a note: %v", err)
	}

	time.Sleep(1 * time.Second)
	// }

	response.CloseSend()
	<-waitc

}
