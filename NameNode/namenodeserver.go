package main

import (
	"Tarea2/DataNode/datanode"
	"Tarea2/NameNode/namenode"
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"math"
	"net"
	"os"
	"sync"
	"time"

	"google.golang.org/grpc"
)

// ServerNameNode ...
type ServerNameNode struct {
	integer int32
	nodos   []string
	loglog  sync.Mutex
}

// ManejarPropuesta ...
func (snn *ServerNameNode) ManejarPropuesta(incomestream namenode.NameNodeHandler_ManejarPropuestaServer) error {

	for {
		in, err := incomestream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("Se recibe una propuesta del datanode")
		nodo1 := PingDataNode(snn.nodos[0])
		nodo2 := PingDataNode(snn.nodos[1])
		nodo3 := PingDataNode(snn.nodos[2])
		cantidad := in.CantidadPartes
		nombreLibro := in.NombreLibro

		var propNodo1 []int32
		var propNodo2 []int32
		var propNodo3 []int32
		var mensaje *namenode.Propuesta
		if nodo1+nodo2+nodo3 == 3 {
			mensaje = in
			log.Printf("Propuesta aprobada")
			if err := incomestream.Send(mensaje); err != nil {
				log.Printf("Error en propuesta  %s", err)
				return err
			}
		} else if nodo1+nodo2+nodo3 == 2 {

			if nodo1 == 0 {
				propNodo2 = makeRange(0, int32(math.Floor(float64(cantidad)/float64(2))))
				propNodo3 = makeRange(int32(math.Floor(float64(cantidad)/float64(2))), int32(cantidad))
				propNodo1 = []int32{}

			} else if nodo2 == 0 {
				propNodo1 = makeRange(0, int32(math.Floor(float64(cantidad)/float64(2))))
				propNodo3 = makeRange(int32(math.Floor(float64(cantidad)/float64(2))), int32(cantidad))
				propNodo2 = []int32{}

			} else if nodo3 == 0 {
				propNodo1 = makeRange(0, int32(math.Floor(float64(cantidad)/float64(2))))
				propNodo2 = makeRange(int32(math.Floor(float64(cantidad)/float64(2))), int32(cantidad))
				propNodo3 = []int32{}

			}
			mensaje = &namenode.Propuesta{
				Datanode1:      propNodo1,
				Datanode2:      propNodo2,
				Datanode3:      propNodo3,
				CantidadPartes: cantidad,
				Status:         namenode.PropuestaStatus_Rechazado,
				NombreLibro:    nombreLibro}
			log.Printf("Propuesta Rechazada")
			if err := incomestream.Send(mensaje); err != nil {
				log.Printf("Error en propuesta  %s", err)
				return err
			}
		} else if nodo1+nodo2+nodo3 == 1 {
			if nodo1 == 1 {
				propNodo2 = []int32{}
				propNodo3 = []int32{}
				propNodo1 = makeRange(0, int32(cantidad))

			} else if nodo2 == 1 {
				propNodo1 = []int32{}
				propNodo3 = []int32{}
				propNodo2 = makeRange(0, int32(cantidad))

			} else if nodo3 == 1 {
				propNodo2 = []int32{}
				propNodo1 = []int32{}
				propNodo3 = makeRange(0, int32(cantidad))

			}
			mensaje = &namenode.Propuesta{
				Datanode1:      propNodo1,
				Datanode2:      propNodo2,
				Datanode3:      propNodo3,
				CantidadPartes: cantidad,
				Status:         namenode.PropuestaStatus_Rechazado,
				NombreLibro:    nombreLibro}
			log.Printf("Propuesta Rechazada")
			if err := incomestream.Send(mensaje); err != nil {
				log.Printf("Error en propuesta  %s", err)
				return err
			}
		} else {
			propNodo2 = []int32{}
			propNodo1 = []int32{}
			propNodo3 = []int32{}
			mensaje = &namenode.Propuesta{
				Datanode1:      propNodo1,
				Datanode2:      propNodo2,
				Datanode3:      propNodo3,
				CantidadPartes: cantidad,
				Status:         namenode.PropuestaStatus_Rechazado,
				NombreLibro:    nombreLibro}
			log.Printf("Propuesta Rechazada")
			if err := incomestream.Send(mensaje); err != nil {
				log.Printf("Error en propuesta  %s", err)
				return err
			}
		}
		snn.loglog.Lock()
		savePropuesta(*mensaje)
		log.Printf("Informacion del libro subido guardada")
		snn.loglog.Unlock()
	}
}

func savePropuesta(propuesta namenode.Propuesta) bool {
	infoLibro := []string{}
	infoLibro = append(infoLibro, propuesta.NombreLibro+" Cantidad_Partes "+fmt.Sprint(propuesta.CantidadPartes)+"\n")
	for _, filenumber := range propuesta.Datanode1 {
		linea := propuesta.NombreLibro + "_parte_" + fmt.Sprintf("%d", filenumber) + " " + "dist141" + "\n"
		infoLibro = append(infoLibro, linea)
	}
	for _, filenumber := range propuesta.Datanode2 {
		linea := propuesta.NombreLibro + "_parte_" + fmt.Sprintf("%d", filenumber) + " " + "dist142" + "\n"
		infoLibro = append(infoLibro, linea)
	}
	for _, filenumber := range propuesta.Datanode3 {
		linea := propuesta.NombreLibro + "_parte_" + fmt.Sprintf("%d", filenumber) + " " + "dist143" + "\n"
		infoLibro = append(infoLibro, linea)
	}
	archivo := "log.txt"
	file, err := os.OpenFile(archivo, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Error creando el archivo: %s", err)
		return false
	}

	datawriter := bufio.NewWriter(file)

	for _, data := range infoLibro {
		_, _ = datawriter.WriteString(data + "\n")
	}

	datawriter.Flush()
	file.Close()

	return true
}

func makeRange(min, max int32) []int32 {
	a := make([]int32, max-min)
	for i := range a {
		a[i] = min + int32(i)
	}
	return a
}

// PingDataNode ...
func PingDataNode(maquina string) int {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(maquina, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Conexion fallida: %s", err)
	}

	defer conn.Close()

	link := datanode.NewDataNodeHandlerClient(conn)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	response, err := link.Ping(ctx)

	if err != nil {
		log.Printf("Conexion fallida: %s", err)
	}

	waitc := make(chan bool)

	go func() {
		for {
			in, err := response.Recv()
			if err == io.EOF {
				close(waitc)
				return
			}

			if err != nil {
				log.Fatalf("Error al realizar un ping al datanode %s : %v", maquina, err)
			}

			if in.CanReceive == false {
				log.Printf("La maquina no puede recibir archivos, y retornó el siguiente mensaje: %s ", in.Message)
			}
			waitc <- in.CanReceive
		}
	}()

	mensaje := datanode.Ping{
		CanReceive: true,
		Message:    "",
	}

	if err := response.Send(&mensaje); err != nil {
		log.Fatalf("Failed to send a note: %v", err)
	}

	response.CloseSend()
	var respuesta int
	respuesta = 0
	receive := <-waitc
	if receive {
		respuesta = 1
	}
	return respuesta
}

func main() {
	server := ServerNameNode{}
	puerto := ":8080"
	server.nodos = []string{"dist141:8080", "dist142:8080", "dist143:8080"}
	lis, err := net.Listen("tcp", puerto)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", puerto, err)
	}
	grpcServer := grpc.NewServer()

	namenode.RegisterNameNodeHandlerServer(grpcServer, &server)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port %s: %v", puerto, err)
	}
}
