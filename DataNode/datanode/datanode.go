package datanode

import (
	"io"
	"log"
)

// ServerDataNode ...
type ServerDataNode struct {
	integer int32
}

// UploadFile ...
func (dns *ServerDataNode) UploadFile(incomestream DataNodeHandler_UploadFileServer) error {

	for {
		in, err := incomestream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		// mensaje := in.Content
		// nameFile := in.ChunkName

		// if _, err12 := os.Stat("/libro"); os.IsNotExist(err12) {
		// 	errFolder := os.Mkdir("libro", 0755)
		// 	fmt.Printf("El andres se la come")
		// 	if errFolder != nil {
		// 		log.Fatal(err)
		// 	}
		// }

		// Andres := ioutil.WriteFile("libro/"+nameFile+"andresql", mensaje, 0644)
		// if Andres != nil {
		// 	log.Fatal(Andres)
		// }
		/*
			Hago algo con lo que recibo con in
		*/
		log.Printf("%s", in.ChunkName)

		//Se envia la respuesta
		if err := incomestream.Send(&Response{
			Message: in.ChunkName,
			Status:  StatusCode_Success}); err != nil {
			return err
		}

	}

	// return nil
}

// DownloadFile ...
func (dns *ServerDataNode) DownloadFile(incomestream DataNodeHandler_DownloadFileServer) error {

	return nil
}
