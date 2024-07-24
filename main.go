package main

import (
	"go-web-native/config"
	"go-web-native/controllers/categorycontroller"
	"go-web-native/controllers/homecontroller"
	"log"
	"net/http"
)

func main() {
	config.ConnectDb()
	http.HandleFunc("/", homecontroller.Welcome)
	http.HandleFunc("/categories", categorycontroller.Index)
	http.HandleFunc("/categories/add", categorycontroller.Add)
	http.HandleFunc("/categories/edit", categorycontroller.Edit)
	http.HandleFunc("/categories/delete", categorycontroller.Delete)

	log.Println("Server running on port http://localhost:8080")
	http.ListenAndServe("127.0.0.1:8080", nil)
}
