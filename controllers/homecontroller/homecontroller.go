package homecontroller

import (
	"html/template"
	"net/http"
)

func Welcome(web http.ResponseWriter, request *http.Request) {
	temp, err := template.ParseFiles("views/home/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(web, nil)
}
