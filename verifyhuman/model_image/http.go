package model_image


import (
	"net/http"
	"rendertemplates"
	"html/template"
	"fmt"
)


func Http(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	let := ModelStruct{}
	let.Route(r)

	if len(let.Urls) > 0 {

		if let.Urls[0] == "puzzle" {
			let.CheckPuzzle(w, r)
			fmt.Fprintf(w, "%s", let.Details)
		}
	}else{
		let.TemplateHTML = template.HTML(Model_All_Images())
		fmt.Fprintf(w, rendertemplates.HTML(let, "http://localhost:7070/dingtra-web-assets/templates/home/image.html"))
	}
}