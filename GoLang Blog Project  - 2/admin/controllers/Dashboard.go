package controllers

import (
	"blogmod/admin/helpers"
	"fmt"
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Dashboard struct{}

func (dashboard Dashboard) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, err := template.ParseFiles(helpers.Include("dashboard/list")...)

	if err != nil {
		fmt.Println(err)
		return
	}

	view.ExecuteTemplate(w, "index", nil)
}
