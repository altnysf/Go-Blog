package main

import (
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// go get github.com/julienschmidt/httprouter

func main() {
	r := httprouter.New()
	r.GET("/yazilar/:slug", Anasayfa)
	http.ListenAndServe(":8080", r)
}

func Anasayfa(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, _ := template.ParseFiles("index.html")
	data := params.ByName("slug")
	view.Execute(w, data)
}
