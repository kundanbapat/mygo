package main

import (
	"fmt"
	"net/http"
	"strings"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Howdyyyy " + message
	w.Write([]byte(message))
}

func getData(w http.ResponseWriter, r *http.Request) {
	// serve Files
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/getData/")
	message = GetDataResponse()

	message = "<html>" + message + Footer() + "</html>"

	w.Write([]byte(message))
}

func postData(w http.ResponseWriter, r *http.Request) {
	// serve Files
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/getData/")
	message = PostDataResponse()

	message = "<html>" + message + Footer() + "</html>"

	w.Write([]byte(message))
}

func main() {
	http.HandleFunc("/getData/", getData)
	http.HandleFunc("/postData/", postData)

	fs := http.FileServer(http.Dir("./"))
	http.Handle("/", http.StripPrefix("/", fs))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("listen failed ", err.Error())
	}
}
