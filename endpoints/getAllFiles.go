package endpoints

import (
	"SimpleServer/service"
	"encoding/json"
	"fmt"
	"net/http"
)

func (app *AppContext) GetAllFiles(w http.ResponseWriter, r *http.Request) {
	fileRecords := service.GetAllFilesDb(app.User, app.DB)
	fmt.Println("allFiles Endpoint Hit")
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(fileRecords)
}
