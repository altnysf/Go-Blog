package config

import (
	admin "blogmod/admin/controllers"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Routes() *httprouter.Router {
	r := httprouter.New()

	//Admin
	//Blog Posts
	r.GET("/admin", admin.Dashboard{}.Index)
	r.GET("/admin/yeni-ekle", admin.Dashboard{}.NewItem)

	r.ServeFiles("/admin/assets/*filepath", http.Dir("admin/assets/"))

	return r
}
