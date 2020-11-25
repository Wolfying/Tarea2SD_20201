package main

import (
	"log"
	"net"
)

func getOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

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

}
