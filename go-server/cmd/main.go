package main

import (
	"net/http"

	"github.com/x022513319/asynchronous-processing-practice/api"
	"github.com/x022513319/asynchronous-processing-practice/api/queue"
)

// API server entry point
func main() {

	redisClient := queue.NewRedisClient("localhost")
	publisher := queue.NewPublisher(redisClient)

	r := api.NewRouter(publisher)

	http.ListenAndServe(":8080", r)
}
