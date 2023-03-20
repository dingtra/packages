package bussiness

import (
	"net/http"
	"html/template"
	"strings"
	"encoding/json"
)

type AjaxStruct struct {
	Usersid string
	JsonVal string
}


func (let *AjaxStruct) VerifyBussiness (r *http.Request) {
	Error := map[string]template.HTML{}
	if strings.ToLower(r.Method) == "post" {
		filenameid := strings.TrimSpace(strings.ToLower(r.FormValue("id")))
		Error["attachxr09"+filenameid] = template.HTML(GenerateForm(r))
	}

	jsn, _ := json.Marshal(Error)
		
	let.JsonVal = string(jsn)
}