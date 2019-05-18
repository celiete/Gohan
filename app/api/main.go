package main

import (
	"log"
	"net/http"
	"os"

	"github.com/celiete/gohan"
	"github.com/celiete/gohan/handler"
	"github.com/julienschmidt/httprouter"
	"github.com/Chensienyong/godotenv"
)

func main() {
	godotenv.Load()
	gohan := gohan.NewGohan()
	handler := handler.NewHandler(gohan)

	router := httprouter.New()

	router.GET("/healthz", handler.Healthz)

	// Start server
	log.Println("Listening at port", os.Getenv("API_PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("API_PORT"), router))
}
