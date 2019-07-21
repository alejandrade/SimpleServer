package config

import (
	"SimpleServer/endpoints"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func BasicAuth(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Security Checkpoint", r.URL)
		_, _, ok := r.BasicAuth()
		//ignore password becuase this isn't real security
		if !ok {
			http.Error(w, "Unauthorized", 401)
		} else {
			h.ServeHTTP(w, r)
		}
	})
}

func HandleRequest(dbclient *dynamodb.DynamoDB, s3ClientUploader *s3manager.Uploader, s3ClientDownload *s3manager.Downloader) {
	context := endpoints.AppContext{DB: dbclient, S3Uploader: s3ClientUploader, S3Downloader: s3ClientDownload}
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/file", (&context).UploadFile).Methods("POST")
	myRouter.HandleFunc("/file", (&context).GetAllFiles).Methods("GET")
	myRouter.HandleFunc("/file/{fileId}", (&context).DownloadFile).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", BasicAuth(myRouter)))
}
