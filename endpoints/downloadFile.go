package endpoints

import (
	"SimpleServer/service"
	"bytes"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"net/url"
)

func (app *AppContext) DownloadFile(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fileId := params["fileId"]
	fileRecord, err := service.GetFileDb(fileId, GetUserFromRequest(r), app.DB)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	log.Println("downloadFile Endpoint Hit")

	w.Header().Set("content-type", "application/octet-stream")
	w.Header().Set("content-disposition", "attachment; filename="+url.QueryEscape(fileRecord.FileName))

	fileBytes, err := service.GetFileS3(fileRecord, app.S3Downloader)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	_, err = io.Copy(w, bytes.NewReader(fileBytes))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

}
