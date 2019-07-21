package endpoints

import (
	"SimpleServer/service"
	"encoding/json"
	"log"
	"net/http"
)

func (app *AppContext) GetAllFiles(w http.ResponseWriter, r *http.Request) {
	fileRecords, err := service.GetAllFilesDb(app.User, app.DB)
	log.Println("allFiles Endpoint Hit")

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if len(fileRecords) == 0 {
		http.Error(w, "not found", 404)
		return
	}

	w.Header().Set("content-type", "application/json")
	err = json.NewEncoder(w).Encode(fileRecords)
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Println("", err.Error())
		return
	}

}
