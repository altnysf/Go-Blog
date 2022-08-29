package main

import (
	"html/template"
	"io"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

// go get github.com/julienschmidt/httprouter

func main() {
	r := httprouter.New()
	r.GET("/", Anasayfa)
	r.POST("/deneme", Deneme)
	http.ListenAndServe(":8080", r)
}

func Anasayfa(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, _ := template.ParseFiles("index.html")
	view.Execute(w, nil)
}

func Deneme(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	r.ParseMultipartForm(10 << 20) // Max 10 Mb and 20 Parts
	file, header, _ := r.FormFile("fileName")
	f, _ := os.OpenFile(header.Filename, os.O_WRONLY|os.O_CREATE, 8666)
	io.Copy(f, file)
}
