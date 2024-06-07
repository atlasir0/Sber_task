package main

import (
	"net/http"
	"strconv"
	"todolist/internal/api"
	"todolist/internal/db"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	router := mux.NewRouter()

	dbConn, cfg, err := db.InitDB()
	if err != nil {
		log.Fatalf("could not connect to the database: %v", err)
	}
	defer db.CloseDB(dbConn, cfg)

	apiInstance, err := api.NewAPI(router, dbConn)
	if err != nil {
		log.Fatalf("could not create API instance: %v", err)
	}
	apiInstance.RegisterHandlers()

	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	port := cfg.PortBK.Port
	log.Printf("Starting server on port %d", port)
	if err := http.ListenAndServe(":"+strconv.Itoa(port), router); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
