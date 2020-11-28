package datanode

import (
	"io"
	"io/ioutil"
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

		nombreParte := in.ChunkName
		fileName := in.FileName
		_ = fileName
		file := in.Content

		// if _, err12 := os.Stat("/libros"); os.IsNotExist(err12) {
		// 	errFolder := os.Mkdir("libros", 0755)
		// 	fmt.Printf("Carpeta Creada")
		// 	if errFolder != nil {
		// 		log.Fatal(err)
		// 	}
		// }
		filePath := "libros/" + nombreParte
		filePart := ioutil.WriteFile(filePath, file, 0644)

		if filePart != nil {
			log.Fatal(filePart)
		}

		/*
			Preparación respuesta a cliente.
		*/
		log.Printf("Archivo creado en %s", filePath)

		//Se envia la respuesta
		if err := incomestream.Send(&Response{
			Message: in.ChunkName,
			Status:  StatusCode_Success}); err != nil {
			return incomestream.Send(&Response{Message: err.Error(),
				Status: StatusCode_InternalError})
		}

	}

	// return nil
}

// DownloadFile ...
func (dns *ServerDataNode) DownloadFile(incomestream DataNodeHandler_DownloadFileServer) error {

	return nil
}
