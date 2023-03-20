package makeding


import (
	"net/http"
	"rundb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
	"strconv"
	"html/template"
	"encoding/json"
	"fmt"
	
)


func (let *MakeDingStruct) ConnectExternalBranch (r *http.Request) {
	Error := map[string]template.HTML{}

	if strings.ToLower(r.Method) == "post" {

		data := strings.TrimSpace(strings.ToLower(r.FormValue("data")))

		getcoll := rundb.Connect("app").Collection("branch")
		getuser := rundb.Connect("app").Collection("register")
	
		finder := rundb.FindAll(getcoll, bson.M{})
		theid, _ := primitive.ObjectIDFromHex(data)
		finduser := rundb.FindOne(getuser, bson.M{"_id":theid })

		fmt.Println(finduser, theid)
	
		details := []string{}
	
		mapper := make(map[string]string)
	
		for _, item := range finder {
	
			if data == item["usersid"].(primitive.ObjectID).Hex() {
	
				details = append(details, item["name"].(string))
				mapper[item["name"].(string)] = item["_id"].(primitive.ObjectID).Hex()
			}
		}
	
		final := []string{}

		if len(details) > 0 {
			for _, k := range details {
				final = append(final, "<div class='brnch82791-slct'>")
				final = append(final, "<div id='brnch82791-slct01'><span>"+RemoveSlash(k)+"</span></div>")
				final = append(final, "<div id='brnch82791-slct02'><input type='checkbox' id='brnc12' name='"+strconv.Itoa(len(details))+"' value='"+mapper[k]+"' x='"+RemoveSlash(k)+"'></div>")
				final = append(final, "</div>")
			}
		}
	
	
		if len(final) > 0 {
			Error["nwdngtext094k"] = template.HTML(strings.Join([]string{
				"<div class='brnch-form-class'><div class='brnch-form-class-ch'>",
				"<div style='display:flex;' id='brnch82791-title'><span>Connect to an existing branch</span>  </div>",
				MakeUser(finduser),
				"<div style='margin-top:10px;border:1px solid #ccc;border-radius:4px;padding:5px 5px 5px 5px;'  id='brtogglex90r'>",
				"<div id='brnchjax090res'>"+strings.Join(final, "")+"</div>",
				"<div id='brnch82791-about'><span>Connect to a single branch at a time <a href='#'>Learn more about branch</a></span></div>",
				"</div>",
				"</div></div>",
				"<script>feather.replace()</script>",
			}, ""))
	
		}else{

			if finduser["_id"].(primitive.ObjectID).Hex() != let.Usersid {

				Error["nwdngtext094k"] = template.HTML( strings.Join([]string{
					"<div class='brnch-form-class'><div class='brnch-form-class-ch'>",
					"<div style='display:flex;' id='brnch82791-title'><span>Connect to an existing branch</span>  </div>",
					MakeUser(finduser),
					"<div style='margin-top:10px;border:1px solid #ccc;border-radius:4px;padding:5px 5px 5px 5px;' id='brtogglex90r'>",
					"<div id='brnch82791-about'><span> The owner has no existing branch yet.</span></div>",
					"<div id='brnch82791-about' style='margin-top:5px;'><span style='font-size:13px;'>cant connect to an empty branch</span></div>",
					"</div>",
					"</div></div>",
					"<script>feather.replace()</script>",
				}, ""))
			}else{
				Error["nwdngtext094k"] = template.HTML( strings.Join([]string{
					"<div class='brnch-form-class'><div class='brnch-form-class-ch'>",
					"<div style='display:flex;' id='brnch82791-title'><span>Connect to an existing branch</span>  </div>",
					MakeUser(finduser),
					// 
					"<div style='margin-top:10px;' id='brtogglex90r'>",
					"<div id='brnch82791-about'><span>Cant connect to an empty branch.</span></div>",
					// create new branch
					`<div id='brnch82791-about' style='margin-top:5px;'>
						<div style="background:green;color:white;padding:5px 10px 5px 10px;border-radius:4px;" class='asp-new-brnch'>
							<div style="width:fit-content;margin:auto;"><i style="height:16px;width:16px;position:relative;top:2px;color:white" data-feather='folder-plus' color='black'></i> <span>Create new branch</span></div>
						</div>
					</div>`,
					// create new branch
					"</div>",
					// 
					
					"</div></div>",
					"<script>feather.replace()</script>",
				}, ""))
			}
		}

	}

	jsn, _ := json.Marshal(Error)
	let.JsonVal = string(jsn)

}


func MakeUser (item primitive.M) string {
	details := []string{}

	if len(item) > 0 {
		
		details  = append(details, `
		 <div style="display:flex;margin-top:10px;" class="userbrnxr0x90">

		`)

		details  = append(details, `
		  <div class="userbrnxr0x91"><span>`+VerifyPic(item)+`</span></div>

		`)


		if item["fullname"] != nil && len(item["fullname"].(string)) > 0 {
			details  = append(details, `
		      <div class="userbrnxr0x92" style="text-transform:capitalize;margin-left:10px;margin-top:3px;"><span>`+item["fullname"].(string)+`</span></div>

		    `)
		}else{
			details  = append(details, `
		      <div class="userbrnxr0x92" style="text-transform:capitalize;margin-left:10px;margin-top:3px;"><span>`+item["username"].(string)+`</span></div>

		    `)
		}

		details  = append(details, `
		    </div>

		`)
	}


	return strings.Join(details, "")
}

func VerifyPic (text primitive.M) string {

	var pic string

	if text["pic"] != nil &&  len(text["pic"].(string)) > 0 {
		if text["pic"].(string)[0:1] == "#" {
			pic = "<span id='userxr090img'><img style='height:25px;width:25px;border-radius:100%;' src='"+text["pic"].(string)[1:]+"' /></span>"
		}else{
			pic = "<span id='userxr090img'><img style='height:25px;width:25px;border-radius:100%;' src='/dingtra-web-assets/images/profile/"+text["pic"].(string)+"' /></span>"
		}
	}else{
		pic = "<span id='userxr090img'><img style='height:25px;width:25px;border-radius:100%;' src='/dingtra-web-assets/svg/pr.png' /></span>"
	}

	return pic
}