package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gocraft/web"
)

// GoSlow waits until you are ready
func GoSlow(response web.ResponseWriter, request *web.Request) {
	rawDelayStr := request.PathParams["delay"]
	delayStr := fmt.Sprintf("%vms", rawDelayStr)
	delay, err := time.ParseDuration(delayStr)
	if err != nil {
		log.Printf("error: %v", err.Error())
		response.WriteHeader(http.StatusNotAcceptable)
		fmt.Fprint(response, `{"error": "'GET /:delay', delay must be an integer"}`)
		return
	}

	time.Sleep(delay)
	response.WriteHeader(http.StatusOK)
	fmt.Fprintf(response, `{"delay": %v}`, rawDelayStr)
}

// Healthcheck responds with online true
func Healthcheck(response web.ResponseWriter, request *web.Request) {
	response.WriteHeader(http.StatusOK)
	fmt.Fprint(response, `{"online": true}`)
}

func main() {
	router := web.New(struct{}{}).
		Middleware(web.LoggerMiddleware).
		Middleware(web.ShowErrorsMiddleware).
		Get("/:delay", GoSlow).
		Get("/healthcheck", Healthcheck)

	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}
	host := fmt.Sprintf("0.0.0.0:%v", port)
	log.Printf("Listening on: %s", host)
	http.ListenAndServe(host, router) // Start the server!
}
