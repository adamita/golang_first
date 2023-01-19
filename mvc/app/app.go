package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/adamita/golang_first/mvc/controllers"
)

func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
