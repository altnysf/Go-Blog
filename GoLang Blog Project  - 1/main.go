package main

import (
	"blogmod/config"
	"net/http"
)

func main() {
	http.ListenAndServe(":8080", config.Routes())
}
