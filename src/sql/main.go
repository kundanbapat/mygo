package main

import (
	"fmt"
	cl "grpc_server_lib"
)

func main() {
	fmt.Println("===== Entering grpc_server_lib/main.go =====")
	cl.RunGRPCServer()
}