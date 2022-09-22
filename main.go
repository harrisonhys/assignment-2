package main

import (
	"assignment-two/middleware"
	"assignment-two/api"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	r := api.SetupRouter()
	r.Use(middleware.CORS)

	var PORT = ":9090"

	log.Info("Starting server on port ", PORT)

	log.Fatal(http.ListenAndServe(PORT, r))
}