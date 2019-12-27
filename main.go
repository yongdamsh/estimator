package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/yongdamsh/estimator/task"
)

var (
	port = flag.String("port", ":5050", "server port")
)

func main() {
	flag.Parse()

	http.Handle("/tasks/", task.NewHandler())
	log.Fatal(http.ListenAndServe(*port, nil))
}
