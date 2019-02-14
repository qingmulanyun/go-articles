package main

import (
	"github.com/gorilla/mux"
	"os"
	"fmt"
	"net/http"
	"go-articles/controllers/api/v1"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/v1/articles", controllersV1.CreateArticle).Methods("POST")
	router.HandleFunc("/v1/articles/{id}", controllersV1.GetArticle).Methods("GET") //  articles/1
	router.HandleFunc("/v1/tags/{tagName}/{date}", controllersV1.GetArticles).Methods("GET") //  tags/tag2/date/20190214

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" //localhost
	}

	fmt.Println(port)

	err := http.ListenAndServe(":" + port, router) //Launch the app
	if err != nil {
		fmt.Print(err)
	}
}
