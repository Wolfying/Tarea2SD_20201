#················································Name Node·························································
runNamenode:
	rm -f log.txt
	protoc -I NameNode/namenode NameNode/namenode/namenode.proto --go_out=plugins=grpc:./
	go run NameNode/namenodeserver.go

compileNamenode:
	protoc -I NameNode/namenode NameNode/namenode/namenode.proto --go_out=plugins=grpc:./
#··················································································································
#················································Data Node·························································
runDatanode:
	protoc -I DataNode/datanode DataNode/datanode/datanode.proto --go_out=plugins=grpc:./
	go run DataNode/datanodeserver.go

compileDatanode:
 	protoc -I DataNode/datanode DataNode/datanode/datanode.proto --go_out=plugins=grpc:./
 
#··················································································································
#················································Cliente···························································
runCliente:
	go run Clients/main.go
#··················································································································

