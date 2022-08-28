package main

import (
	"html/template"
	"net/http"
)

func main() {
	// http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/", Anasayfa)
	// http.HandleFunc("/detay", Detay)
	http.ListenAndServe(":8080", nil)
}

func Anasayfa(w http.ResponseWriter, r *http.Request) {
	view, _ := template.ParseFiles("index.html", "detay.html", "navbar.html")
	data := "Go'dan gelen veri..."
	view.ExecuteTemplate(w, "anasayfa", data)
}

// func Detay(w http.ResponseWriter, r *http.Request) {
// 	view, _ := template.ParseFiles("detay.html","navbar.html")
// 	view.ExecuteTemplate(w, "detay", nil)
// }
