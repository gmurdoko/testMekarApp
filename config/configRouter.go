package config

import (
	"fmt"
	"log"
	"net/http"
	"testMekarApp/utils"

	"github.com/gorilla/mux"
)

func CreateRouter() *mux.Router {
	router := mux.NewRouter()
	return router
}

func RunServer(router *mux.Router) {
	appHost := utils.ViperGetEnvironment("APP_HOST", "localhost")
	appPort := utils.ViperGetEnvironment("APP_PORT", "8080")
	hostListening := fmt.Sprintf("%v:%v", appHost, appPort)
	log.Printf("App ready to listening on %v", hostListening)
	err := http.ListenAndServe(hostListening, router)
	if err != nil {
		log.Fatal(err)
	}
}
