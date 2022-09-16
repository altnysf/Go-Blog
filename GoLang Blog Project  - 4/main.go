package main

import (
	admin_models "blogmod/admin/models"
	"blogmod/config"
	"net/http"
)

func main() {
	admin_models.Post{}.Migrate()

	http.ListenAndServe(":8080", config.Routes())

}
