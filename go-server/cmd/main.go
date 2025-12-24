package main

import (
	"fmt"
	"net/http"

	"github.com/x022513319/asynchronous-processing-practice/api"
	"github.com/x022513319/asynchronous-processing-practice/api/queue"
)

var (
	server_addr = "localhost"
	server_port = ":8080"
	redis_port  = ":6379"
)

// API server entry point
func main() {

	redisClient := queue.NewRedisClient(server_addr + redis_port)
	publisher := queue.NewPublisher(redisClient)

	r := api.NewRouter(publisher)

	fmt.Println("Go-Server Listening on " + server_addr + server_port)
	http.ListenAndServe(server_port, r)

}
