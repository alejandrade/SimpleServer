package endpoints

import (
	"SimpleServer/service"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"io"
	"net/http"
)

func (app *AppContext) UploadFile(w http.ResponseWriter, r *http.Request) {
	fileRecord := service.FileRecord{User: app.User}
	handleFile(app.S3Uploader, r, &fileRecord)
	service.SaveFileDb(fileRecord, app.DB)
	fmt.Println("uploadFile Endpoint Hit")
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(fileRecord)
}

func handleFile(s3 *s3manager.Uploader, r *http.Request, fileRecord *service.FileRecord) {
	var Buf bytes.Buffer
	// in your case file would be fileupload
	file, header, err := r.FormFile("file")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Copy the file data to my buffer
	io.Copy(&Buf, file)
	service.SaveFileS3(&Buf, *header, s3, fileRecord)
}
