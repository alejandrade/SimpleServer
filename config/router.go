package config

import (
	"SimpleServer/endpoints"
	"SimpleServer/util"
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

// injecting all my depndencies like this is not scalable. I would love to figure out a better solution.
func HandleRequest(dbClient *dynamodb.DynamoDB, uploader *s3manager.Uploader, downloader *s3manager.Downloader, properties *util.Properties) {
	var ctx = &endpoints.AppContext{DB: dbClient, S3Uploader: uploader, S3Downloader: downloader, Properties: properties}
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/file", (ctx).UploadFile).Methods("POST")
	myRouter.HandleFunc("/file", (ctx).GetAllFiles).Methods("GET")
	myRouter.HandleFunc("/file/{fileId}", (ctx).DownloadFile).Methods("GET")
	log.Fatal(http.ListenAndServe(":"+properties.Port, BasicAuth(myRouter)))
}
