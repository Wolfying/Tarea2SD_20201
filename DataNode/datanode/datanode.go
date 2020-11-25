package datanode

// ServerDataNode ...
type ServerDataNode struct {
}

// UploadFile ...
func (dns *ServerDataNode) UploadFile(incomestream DataNodeHandler_UploadFileServer) error {

	return nil
}

// DownloadFile ...
func (dns *ServerDataNode) DownloadFile(incomestream DataNodeHandler_DownloadFileServer) error {

	return nil
}
