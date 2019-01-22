package main

import (
	"io/ioutil"
)

// GetDataResponse returns response received from backend
func GetDataResponse() string {
	return "<h2>Response to GET from grpc & mysql Microservices</h2>"
}

// PostDataResponse returns repsonse to a POST request
func PostDataResponse() string {
	return "<h2>Response to POST from grpc & mysql Microservices</h2>"
}

// Footer returns standard links at end of every reponse
func Footer() string {
	data, err := ioutil.ReadFile("./footer.html")
	if err != nil {
		return ""
	}
	return string(data)

}
