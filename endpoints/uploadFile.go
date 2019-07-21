package endpoints

import (
	"SimpleServer/service"
	"bytes"
	"encoding/json"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"io"
	"log"
	"net/http"
)

func (app *AppContext) UploadFile(w http.ResponseWriter, r *http.Request) {
	fileRecord := service.FileRecord{User: app.User}
	err := handleFile(app.S3Uploader, r, &fileRecord)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = service.SaveFileDb(fileRecord, app.DB)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	log.Println("uploadFile Endpoint Hit")
	w.Header().Set("content-type", "application/json")
	err = json.NewEncoder(w).Encode(fileRecord)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

}

func handleFile(s3 *s3manager.Uploader, r *http.Request, fileRecord *service.FileRecord) error {
	var Buf bytes.Buffer
	// in your case file would be fileupload
	file, header, err := r.FormFile("file")
	if err != nil {
		return err
	}
	defer file.Close()

	// Copy the file data to my buffer
	_, err = io.Copy(&Buf, file)

	if err != nil {
		return err
	}

	err = service.SaveFileS3(&Buf, *header, s3, fileRecord)

	if err != nil {
		return err
	}

	return nil
}
