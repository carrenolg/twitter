package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/carrenolg/twitter/middlew"
	"github.com/carrenolg/twitter/routes"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Handler() {
	router := mux.NewRouter()

	// middle
	router.HandleFunc("/registry", middlew.Checkdb(routes.Registry)).Methods("POST")
	router.HandleFunc("/login", middlew.Checkdb(routes.Login)).Methods("POST")
	router.HandleFunc("/showprofile", middlew.Checkdb(middlew.ValidateJwt(routes.ShowProfile))).Methods("GET")
	router.HandleFunc("/updateprofile", middlew.Checkdb(middlew.ValidateJwt(routes.ModifyProfile))).Methods("PUT")
	router.HandleFunc("/tweet", middlew.Checkdb(middlew.ValidateJwt(routes.InsertTweet))).Methods("POST")
	router.HandleFunc("/readtweets", middlew.Checkdb(middlew.ValidateJwt(routes.ReadTweets))).Methods("GET")
	router.HandleFunc("/deletetweet", middlew.Checkdb(middlew.ValidateJwt(routes.DeleteTweet))).Methods("DELETE")
	router.HandleFunc("/uploadavatar", middlew.Checkdb(middlew.ValidateJwt(routes.UploadAvatar))).Methods("POST")
	router.HandleFunc("/getavatar", middlew.Checkdb(routes.GetAvatar)).Methods("GET")
	router.HandleFunc("/uploadbanner", middlew.Checkdb(middlew.ValidateJwt(routes.UploadBanner))).Methods("POST")
	router.HandleFunc("/getbanner", middlew.Checkdb(routes.GetBanner)).Methods("GET")

	router.HandleFunc("/setrelation", middlew.Checkdb(middlew.ValidateJwt(routes.SetRelation))).Methods("POST")
	router.HandleFunc("/unsetrelation", middlew.Checkdb(middlew.ValidateJwt(routes.UnsetRelation))).Methods("DELETE")
	router.HandleFunc("/checkrelation", middlew.Checkdb(middlew.ValidateJwt(routes.CheckRelation))).Methods("GET")
	router.HandleFunc("/listusers", middlew.Checkdb(middlew.ValidateJwt(routes.ListUsers))).Methods("GET")
	router.HandleFunc("/gettweetsfollowers", middlew.Checkdb(middlew.ValidateJwt(routes.GetTweetsFollowers))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
