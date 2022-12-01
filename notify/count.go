package notify


import (
	"rundb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"
)


func CountNotify(id string) string {
	counter := []string{}

	getcoll := rundb.Connect("app").Collection("notification")

	for _, item := range rundb.FindAll(getcoll){
		if item["owner"] != nil {
			if item["owner"].(string) == id && item["viewed"] == nil{
				counter = append(counter, item["usersid"].(primitive.ObjectID).Hex())
			}
		}
	}

	if len(counter) > 0 {
		return "<div id='countnotify-div'><span><button>"+strconv.Itoa(len(counter))+"</button></span></div>"
	}else{
		return ""
	}

	return ""
}
