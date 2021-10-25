package middlew

import (
	"net/http"

	"github.com/carrenolg/twitter/db"
)

func Checkdb(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db.CheckConnection() == 0 {
			err := "lost connection to mongodb"
			http.Error(w, err, 500)
			return
		}
		handler.ServeHTTP(w, r)
	}
}
