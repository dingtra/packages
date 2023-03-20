package loadbranch


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

var (
	CheckSvg = `<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="feather feather-check" style="color:green;height:13px;width:13px;position:relative;top:2px;" color="black"><polyline points="20 6 9 17 4 12"></polyline></svg>`
)


func (let *AjaxStruct ) VerifyQuery ( r *http.Request ) {
	Error := map[string]template.HTML{}
	if strings.ToLower(r.Method) == "post" {

		data := strings.TrimSpace(strings.ToLower(r.FormValue("data")))
		name := strings.TrimSpace(strings.ToLower(r.FormValue("x")))
		
		branchchild := rundb.Connect("app").Collection("inbranch")
		findall := rundb.FindAll(branchchild, bson.M{})

		details := []string{}
		counter := 0

		for _, item := range findall {

			if item["branchid"].(string) == data {
				counter +=1
				details = append(details, `
					<div class="brnch82791-slct">
						<div id="brnch82791-slct01"><span style="font-size:14px;">`+RemoveSlash(item["name"].(string))+`</span></div>
						<div id="brnch82791-slct02"><input type="checkbox" id="inbranch-res`+strconv.Itoa(counter)+`" value="`+item["_id"].(primitive.ObjectID).Hex()+`" ></div>
					</div>
				`)
			}
		}

		// // grabmodes := map[string]bool{}
		// allitems := map[string]map[string]string{}
		// record := map[string]string{}
		// counter := 0
		// sum := 0

		// for i, item := range findall {
	
		// 	if item["usersid"].(primitive.ObjectID).Hex() == let.Usersid  && item["branchid"].(string) == data {

		// 		allitems["mode"+strconv.Itoa(i)] = map[string]string{item["mode"].(string):item["mode"].(string)}
		// 		record[allitems["mode"+strconv.Itoa(i)][item["mode"].(string)]] = `<div class="inbranch-name" style="padding:5px 0px 5px 0px;border-bottom:1px solid #ccc;text-transform:uppercase;font-size:15px;"><span>`+item["mode"].(string)+`</span></div>`
		// 		sum +=1
		// 	}
		// }

		// details := []string{}

		// for _, item := range findall {
	
		// 	if item["usersid"].(primitive.ObjectID).Hex() == let.Usersid  && item["branchid"].(string) == data {

		// 		if len(record[item["mode"].(string)]) > 0 {
		// 			counter +=1
		// 			record[item["mode"].(string)] = record[item["mode"].(string)] + `
		// 			<div class="brnch82791-slct">`+
		// 				`<div id="brnch82791-slct01"><span style="font-size:14px;">`+RemoveSlash(item["name"].(string))+`</span></div>`+
		// 				`<div id="brnch82791-slct02"><input type="checkbox" id="inbranch-res`+strconv.Itoa(counter)+`"  kp="`+strconv.Itoa(sum)+`" value="`+item["_id"].(primitive.ObjectID).Hex()+`" ></div>`+
		// 			`</div>`
		// 		}
		// 	}
		// }

		// for _, k := range record {
		// 	details =append(details, k)
		// }
		

		if len(details) > 0 {
			Error["brnchjax090res"] = template.HTML( "<div style='padding:5px 0px 5px 0px;'> <div style='background:white;padding:5px 5px 5px 5px;border-radius:4px;text-transform:capitalize;'>"+name+" <input style='display:none;' type='checkbox' name='branch' value='"+data+"' checked>"+CheckSvg+"</div> </div>"+ strings.Join(details, ""))
		}else{
			Error["brnchjax090res"] = template.HTML( "<div style='padding:5px 0px 5px 0px;'> <div style='background:white;padding:5px 5px 5px 5px;border-radius:4px;text-transform:capitalize;'>"+name+" <input style='display:none;' type='checkbox' name='branch' value='"+data+"' checked>"+CheckSvg+"</div> </div> No in branch available")
		}

	}

	jsn, _ := json.Marshal(Error)
		
	let.JsonVal = string(jsn)

}

func (let *AjaxStruct ) VerifyQuer ( r *http.Request ) {
	Error := map[string]template.HTML{}
	if strings.ToLower(r.Method) == "post" {

		data := strings.TrimSpace(strings.ToLower(r.FormValue("data")))
		name := strings.TrimSpace(strings.ToLower(r.FormValue("x")))
		
		branchchild := rundb.Connect("app").Collection("inbranch")
		findall := rundb.FindAll(branchchild, bson.M{})
		
		findshop := []string{}
		findwaitlis := []string{}
		findbusiness:= []string{}
		findothers := []string{}
		var finalview string

		grabber := make(map[int]string)

		for i, item := range findall {
	
			if item["usersid"].(primitive.ObjectID).Hex() == let.Usersid  && item["branchid"].(string) == data && item["mode"].(string) == "shop" {
				grabber[i] =  "shop##<div class='brnch82791-slct'>"+
				"<div id='brnch82791-slct01'><span style='font-size:14px;'>"+RemoveSlash(item["name"].(string))+"</span></div>"+
				"<div id='brnch82791-slct02'><input type='checkbox' %s value='"+item["_id"].(primitive.ObjectID).Hex()+"' ></div>"+
				"</div>"

			}else if item["usersid"].(primitive.ObjectID).Hex() == let.Usersid  && item["branchid"].(string) == data && item["mode"].(string) == "business" {
				grabber[i] =  "busi##<div class='brnch82791-slct'>"+
				"<div id='brnch82791-slct01'><span style='font-size:14px;'>"+RemoveSlash(item["name"].(string))+"</span></div>"+
				"<div id='brnch82791-slct02'><input type='checkbox' %s value='"+item["_id"].(primitive.ObjectID).Hex()+"' ></div>"+
				"</div>"
			}else if item["usersid"].(primitive.ObjectID).Hex() == let.Usersid  && item["branchid"].(string) == data && item["mode"].(string) == "waitlist" {
				grabber[i] =  "wait##<div class='brnch82791-slct'>"+
				"<div id='brnch82791-slct01'><span style='font-size:14px;'>"+RemoveSlash(item["name"].(string))+"</span></div>"+
				"<div id='brnch82791-slct02'><input type='checkbox' %s value='"+item["_id"].(primitive.ObjectID).Hex()+"' ></div>"+
				"</div>"
			}else if item["usersid"].(primitive.ObjectID).Hex() == let.Usersid  && item["branchid"].(string) == data && item["mode"].(string) == "other" {
				grabber[i] =  "othe##<div class='brnch82791-slct'>"+
				"<div id='brnch82791-slct01'><span style='font-size:14px;'>"+RemoveSlash(item["name"].(string))+"</span></div>"+
				"<div id='brnch82791-slct02'><input type='checkbox' %s value='"+item["_id"].(primitive.ObjectID).Hex()+"' ></div>"+
				"</div>"
			}
		}
		
		for k, item := range grabber {
			grabber[k] = fmt.Sprintf(item, "id='inbranch-res"+strconv.Itoa(k)+"'  kp='"+strconv.Itoa(len(grabber))+"'")
		}

		for _, item := range grabber {

			if item[0:4] == "shop" {
				findshop = append(findshop, item[6:])

			}else if item[0:4] == "busi"  {
				findbusiness = append(findbusiness,  item[6:])
			}else if item[0:4] == "wait"  {
				findwaitlis = append(findwaitlis,  item[6:])

			}else if item[0:4] == "othe"  {
				findothers = append(findothers,  item[6:])
			}
		}

		if len(findshop) > 0 {
			finalview += "<div class='inbranch-name' style='padding:5px 0px 5px 0px;border-bottom:1px solid #ccc;text-transform:uppercase;font-size:12px;font-weight:600;'><span>Shop</span></div>"+strings.Join(findshop, "")
		}

		if len(findbusiness) > 0 {
			finalview += "<div class='inbranch-name' style='padding:5px 0px 5px 0px;border-bottom:1px solid #ccc;text-transform:uppercase;font-size:12px;font-weight:600;'><span>Business</span></div>"+strings.Join(findbusiness, "")
		}

		if len(findwaitlis) > 0 {
			finalview += "<div class='inbranch-name' style='padding:5px 0px 5px 0px;border-bottom:1px solid #ccc;text-transform:uppercase;font-size:12px;font-weight:600;'><span>Waitlist</span></div>"+strings.Join(findwaitlis, "")
		}

		if len(findothers) > 0 {
			finalview += "<div class='inbranch-name' style='padding:5px 0px 5px 0px;border-bottom:1px solid #ccc;text-transform:uppercase;font-size:12px;font-weight:600;'><span>Others</span></div>"+strings.Join(findothers, "")
		}

		if len(finalview) > 0 {
			Error["brnchjax090res"] = template.HTML( "<div style='padding:5px 0px 5px 0px;'> <div style='background:white;padding:5px 5px 5px 5px;border-radius:4px;text-transform:capitalize;font-size:14px;'>"+name+" <input style='display:none;' type='checkbox' name='branch' value='"+data+"' checked>"+CheckSvg+"</div> </div>"+finalview)
		}else{
			Error["brnchjax090res"] = template.HTML( "<div style='padding:5px 0px 5px 0px;'> <div style='background:white;padding:5px 5px 5px 5px;border-radius:4px;text-transform:capitalize;font-size:14px;'>"+name+" <input style='display:none;' type='checkbox' name='branch' value='"+data+"' checked>"+CheckSvg+"</div> </div> No in branch available")
		}
	}

	jsn, _ := json.Marshal(Error)
		
	let.JsonVal = string(jsn)
}


func RemoveSlash (val string) string {

	return strings.Replace(val, "-", " ", 2)
}

