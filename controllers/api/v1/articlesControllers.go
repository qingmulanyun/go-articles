package controllersV1

import (
	"net/http"
	"go-articles/models"
	"encoding/json"
	u "go-articles/utils"
	"github.com/gorilla/mux"
	"go-articles/serializers"
)

var CreateArticle = func(w http.ResponseWriter, r *http.Request) {

	article := &models.Article{}

	err := json.NewDecoder(r.Body).Decode(article)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	resp := article.Create()
	u.Respond(w, resp)
}

var GetArticles = func(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	tagName := vars["tagName"]
	date := vars["date"]
	articles :=  models.GetArticles(tagName, date)
	ids := []uint{}
	tags := []string{}
	// get tags and ids
	for i, a := range articles {
		if i <= 10 { ids = append(ids, a.ID) } //only get first 10 elements
		tags = append(tags, a.Tags...)
	}

	data := serializers.Articles{tagName, len(tags), ids, u.Unique(tags)}
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}


var GetArticle = func(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]
	data := models.GetArticle(id)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}