package notify

import (
	"net/http"
	"rundb"
	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
	"encoding/json"
	"html/template"
)


func (let *NotifyStruct) VerifyNotify (r *http.Request) {
	Error := map[string]template.HTML{}
	if strings.ToLower(r.Method) == "post" {
		getcoll := rundb.Connect("app").Collection("notification")
		rundb.UpdateMany(getcoll, bson.M{"owner":let.Usersid}, bson.M{"opened":"open"})

		Error["ntfy0981"] = template.HTML(GetNotification(let.Usersid))
	}

	jsn, _ := json.Marshal(Error)
	
	let.JsonVal = string(jsn)
}


func (let *NotifyStruct) MarkAllRead ( r *http.Request ) {
	if strings.ToLower(r.Method) == "post" {
		getcoll := rundb.Connect("app").Collection("notification")
		rundb.UpdateMany(getcoll, bson.M{"owner":let.Usersid}, bson.M{"opened":"open"})

		let.Details = template.HTML(GetUnreadNotification(let.Usersid))
	}
}