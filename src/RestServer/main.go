package main

import (
	"fmt"
	"net/http"
	"strings"
)

func work(ch chan string) {
	ch <- "Howdy "
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	ch := make(chan string)
	go work(ch)
	var x = <-ch

	fmt.Println("Request handled: ", x)

	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = x + message
	w.Write([]byte(message))
}

func main() {
	http.HandleFunc("/", handleRequest)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("listen failed ", err.Error())
	}
}
