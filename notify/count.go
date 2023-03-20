package notify


import (
	"rundb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
	// "strconv"
	// "fmt"
)


func CountNotify(id string) string {
	counter := []string{}

	getcoll := rundb.Connect("app").Collection("notification")

	for _, item := range rundb.FindAll(getcoll, bson.M{}){
		if item["owner"] != nil {
			if item["owner"].(string) == id && item["viewed"] == nil{
				counter = append(counter, item["usersid"].(primitive.ObjectID).Hex())
			}
		}
	}

	if len(counter) > 0 {
		
		return `<button style="background:red;border:red;color:white;font-weight:bold;height:10px;width:10px;border-radius:100%;"></button>`
	}else{
		return ""
	}

	return ""
}
