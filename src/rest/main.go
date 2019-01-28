package main

import (
	"fmt"
	"net/http"
	"log"
	"github.com/gorilla/mux"
	// "encoding/json"
	"bytes"
	cl "grpc_client_lib"
)

const (
	passthrough = 0	
	sql_query = 1
	orm = 2
)

type post_content struct {
	content_type int
	content string
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HomeHandler"))
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for key, val := range params {
		fmt.Println("GetHandler: ", key, val)
	}
	w.Write([]byte("GetHandler"))
}

type test_struct struct {
	Test string
}
func PostHandler(w http.ResponseWriter, r *http.Request) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	reqStr := buf.String()
	w.Write([]byte(reqStr))
	w.Write([]byte("XXPostHandler\n"))

	cl.RunGRPCClient(reqStr)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", HomeHandler)
	router.HandleFunc("/gettest", GetHandler).Methods("GET")
	router.HandleFunc("/posttest", PostHandler).Methods("POST")
	http.Handle("/", router)

	fmt.Println("ReST Server started...")

	log.Fatal(http.ListenAndServe(":8000", router))
}
