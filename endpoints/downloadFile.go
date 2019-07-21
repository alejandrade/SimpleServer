package endpoints

import (
	"SimpleServer/service"
	"bytes"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"net/url"
)

func (app *AppContext) DownloadFile(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fileId := params["fileId"]
	fileRecord := service.GetFileDb(fileId, app.User, app.DB)

	fmt.Println("downloadFile Endpoint Hit")
	w.Header().Set("content-type", "application/octet-stream")
	w.Header().Set("content-disposition", "attachment; filename="+url.QueryEscape(fileRecord.FileName))

	fileBytes := service.GetFileS3(fileRecord, app.S3Downloader)
	io.Copy(w, bytes.NewReader(fileBytes))
}
