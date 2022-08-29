package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// go get github.com/julienschmidt/httprouter

func main() {
	r := httprouter.New()
	r.GET("/", Anasayfa)
	r.GET("/deneme", Deneme)
	http.ListenAndServe(":8080", r)
}

func Anasayfa(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, _ := template.ParseFiles("index.html")
	view.Execute(w, nil)
}

func Deneme(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	username := r.FormValue("username")
	check := r.FormValue("check")
	selection := r.FormValue("selection")
	fmt.Println(username, check, selection)
}
