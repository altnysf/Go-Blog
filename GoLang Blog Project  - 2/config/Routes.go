package config

import (
	admin "blogmod/admin/controllers"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Routes() *httprouter.Router {
	r := httprouter.New()
	
	//Admin
	r.GET("/admin", admin.Dashboard{}.Index)

	r.ServeFiles("/admin/assets/*filepath", http.Dir("admin/assets/"))

	return r
}
