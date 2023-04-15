package model_image

import (
	"net/http"
	"strings"
	"html/template"
)


type ModelStruct struct {
	Urls []string
	Details string
	TemplateHTML template.HTML
	Username string
}


func (let *ModelStruct) Route (r *http.Request) {

	url := strings.Split(r.URL.Path[len("/model_img/"):], "/")

	for _, item := range url {

		if len(item) > 0 {
			let.Urls = append(let.Urls, item)
		}
	}
}