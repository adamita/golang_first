package app

import (
	"net/http"
	"https://github.com/adamita/golang_first/mvc/app/controllers"
)

func StartApp(){
	http.HandleFunc("/users", controllers.GetUser)
}