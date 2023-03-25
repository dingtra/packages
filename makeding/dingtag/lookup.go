package dingtag


import (
	"net/http"
	"rundb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"regexp"
	"html/template"
	"strings"
	"encoding/json"
	"fmt"
)



func (let *DingtagStruct ) LookUpUser ( r *http.Request) {
	getcoll := rundb.Connect("app").Collection("register")
	usernames := rundb.FindAll(getcoll, bson.M{})

	Error  := map[string]template.HTML{}

	if strings.ToLower(r.Method) == "post" {
		name := strings.ToLower(strings.TrimSpace(r.FormValue("name")))
		getnames := []string{}

		for _, item := range usernames {
	
			matched, _ := regexp.MatchString(fmt.Sprintf(`^%s`, strings.ToLower(name)), strings.ToLower(item["username"].(string)))
			var getfullname bool;

			if item["fullname"] != nil {
				if len(item["fullname"].(string)) > 0 {
					fullnamebool, _ := regexp.MatchString(fmt.Sprintf(`^%s`, strings.ReplaceAll(strings.ToLower(name), " ", "")), strings.ReplaceAll(strings.ToLower(item["fullname"].(string)), " ", ""))
					getfullname = fullnamebool
				}
			}

			if matched == true || getfullname == true {

				//open
				getnames = append(getnames, "<div id='tglkupxr0x' class='usrtg6790' mx='"+item["_id"].(primitive.ObjectID).Hex()+"'><div style='pointer-events:none;' class='usrtg6790-ch'>")
	
				//image
				getnames = append(getnames, "<div id='usrtg6790pic'><span>")
				if item["pic"] != nil {
					getnames = append(getnames, ImageShow(item["pic"].(string)))
				}else{
					getnames = append(getnames, ImageShow(""))
				}
				getnames = append(getnames, "</span></div>")
				//image

				if item["fullname"] != nil {
					if len(item["fullname"].(string)) > 0 {
						//name
						getnames = append(getnames, "<div id='usrtg6790user'><span>")
						getnames = append(getnames,  item["fullname"].(string))
						getnames = append(getnames, "</span></div>")
						//name
					}else{
						//name
						getnames = append(getnames, "<div id='usrtg6790user'><span>")
						getnames = append(getnames,  item["username"].(string))
						getnames = append(getnames, "</span></div>")
						//name
					}
				}else{
					//name
					getnames = append(getnames, "<div id='usrtg6790user'><span>")
					getnames = append(getnames,  item["username"].(string))
					getnames = append(getnames, "</span></div>")
					//name
				}

				getnames = append(getnames, "</div></div>")
				//close
	
			}
		}

		
		if len(getnames) > 0 {

			Error["lkupxr09x"] = template.HTML("<div class='sug009'>"+strings.Join(getnames, "")+"</div>")
		}else{
			Error["lkupxr09x"]  = template.HTML("<div class='sug009'><div style='padding:10px 5px 10px 5px;'><span style='font-size:14px;color:red;'>No record for <b>"+name+"</b></span></div></div>")
		}

	}

	jsn, _ := json.Marshal(Error)
		
	let.JsonVal = string(jsn)

}