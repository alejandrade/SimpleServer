package endpoints

import "net/http"

func GetUserFromRequest(r *http.Request) string {
	user, _, _ := r.BasicAuth()
	return user
}
