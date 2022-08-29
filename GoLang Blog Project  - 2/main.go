package main

import (
	admin_models "blogmod/admin/models"
	"blogmod/config"
	"net/http"
)

func main() {
	admin_models.Post{}.Migrate()

	post := admin_models.Post{}.Get(1)
	post.Delete()

	http.ListenAndServe(":8080", config.Routes())

}
