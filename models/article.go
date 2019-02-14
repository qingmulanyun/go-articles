package models

import (
	u "go-articles/utils"
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
	"fmt"
)

type Article struct {
	gorm.Model
	Title string `gorm:"type:varchar(64);not null" json:"title"`
	Date string  `gorm:"type:date;not null;index" json:"date"`
	Body string `gorm:"type:text" json:"body"`
	Tags pq.StringArray `gorm:"type:varchar(64)[]" json:"tags"`
}

/*
 This struct function validate the required parameters sent through the http request body

returns message and true if the requirement is met
*/
func (article *Article) Validate() (map[string] interface{}, bool) {

	if article.Title == "" {
		return u.Message(false, "article title should be on the payload"), false
	}

	if article.Body == "" {
		return u.Message(false, "article body should be on the payload"), false
	}

	if article.Date == "" {
		return u.Message(false, "article date should be on the payload"), false
	}

	return u.Message(true, "success"), true
}

func (article *Article) Create() (map[string] interface{}) {

	if resp, ok := article.Validate(); !ok {
		return resp
	}

	GetDB().Create(article)

	resp := u.Message(true, "success")
	resp["article"] = article
	return resp
}

func GetArticle(id string) (*Article) {

	article := &Article{}
	err := GetDB().Table("articles").Where("id = ?", id).First(article).Error
	if err != nil {
		return nil
	}
	return article
}

func GetArticles(tagName string, date string) ([]*Article) {

	articles := make([]*Article, 0)
	err := GetDB().Table("articles").Where("date = ?", date).Where("? = ANY (tags)", tagName).Order("created_at desc").Find(&articles).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return articles
}

