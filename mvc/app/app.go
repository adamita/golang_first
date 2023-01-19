package app

import (
	"net/http"
	"github.com/adamita/golang_first/mvc/controllers"
)

func StartApp(){
	http.HandleFunc("/users", controllers.GetUser)
}