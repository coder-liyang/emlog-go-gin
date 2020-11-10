package main

import (
	"emlog-go-gin/models"
	"emlog-go-gin/routes"
	"net/http"
)

func main() {
	models.ClientDb()
	r := routes.RegisterRoutes()
	r.LoadHTMLGlob("views/**/*")
	r.StaticFS("/public", http.Dir("./public"))
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
