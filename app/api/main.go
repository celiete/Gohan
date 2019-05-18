package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/celiete/gohan"
	"github.com/celiete/gohan/handler"
	"github.com/julienschmidt/httprouter"
)

func main() {
	gohan := gohan.NewGohan()
	handler := handler.NewHandler(gohan)

	router := httprouter.New()

	router.GET("/healthz", handler.Healthz)

	// Start server
	log.Println("Listening at port", os.Getenv("API_PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("API_PORT"), router))
}
