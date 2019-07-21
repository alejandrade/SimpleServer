package config

import (
	"SimpleServer/endpoints"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func BasicAuth(h http.Handler, context *endpoints.AppContext) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("middleware", r.URL)
		user, _, ok := r.BasicAuth()
		if !ok {
			w.WriteHeader(401)
			w.Write([]byte("Unauthorized.\n"))
		} else {
			context.SetName(user)
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
	log.Fatal(http.ListenAndServe(":8080", BasicAuth(myRouter, &context)))
}
