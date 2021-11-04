package middlew

import (
	"fmt"
	"net/http"

	"github.com/carrenolg/twitter/routes"
)

func ValidateJwt(handle http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routes.CheckToken(r.Header.Get("Authorization"))
		if err != nil {
			msg := fmt.Errorf("The token is not valid - Error: %s", err)
			http.Error(w, msg.Error(), http.StatusBadRequest)
			return
		}
		handle.ServeHTTP(w, r)
	}
}
