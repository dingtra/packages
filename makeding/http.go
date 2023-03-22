package makeding


import (
	"net/http"
	"rundb"
	"github.com/dingtra/packages/makeding/loadbranch"
	// "loadbranch"
	"github.com/dingtra/packages/makeding/dingtag"
	// "dingtag"
	// "bussiness"
	"github.com/dingtra/packages/makeding/bussiness"
	"encoding/json"
	"html/template"
	"strings"
	"fmt"
)


func Http (w http.ResponseWriter, r *http.Request ){
	let := MakeDingStruct{}
	let.Route(r)

	session, _ := rundb.Store.Get(r, "session")
	w.Header().Set("Content-Type", "text/html")

	fmt.Println(let.Urls)
	

	if session.Values["username"] != nil && session.Values["usersid"] != nil {
		let.Usersid = session.Values["usersid"].(string)
		if len(let.Urls) == 0 {
			fmt.Fprintf(w, "%s",let.NewDing())

		}else{
			if let.Urls[0] == "s" {
				fmt.Fprintf(w, `
				<html>
				<head>
				    <meta charset="UTF-8">
				    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
				    <script src="/assets/js/jquery.js"></script>
				    <script src="/assets/js/mkdngs090/mkdngs0905xr.js"></script>
				</head>
				<body>
				  <div class="xrkl090"></div>
				</body>
				</html>
				`)
			}else if let.Urls[0] == "shr09x" {
				fmt.Fprintf(w, "%s", template.HTML(let.SharedDing(r)))
			}else if let.Urls[0] == "lkbrx0" {
				let.BranchLookUpForm(r)
				fmt.Fprintf(w, "%s", let.JsonVal)
			
			}else if let.Urls[0] == "lkbrx1" {
				if strings.ToLower(r.Method) == "post" {
					data := strings.TrimSpace(strings.ToLower(r.FormValue("data")))

					if len(data) > 0 {
						let.ConnectExternalBranch(r)
					    fmt.Fprintf(w, "%s", let.JsonVal)	
					}else{
						jsn, _ := json.Marshal(map[string]template.HTML{"xr008xr09":template.HTML(ConnectBranch(let.Usersid))+
						`<script>
						feather.replace()
						</script>`})
						fmt.Fprintf(w, "%s", string(jsn))
					}
				}

			}else if let.Urls[0] == "optnsxr090" {
				let.FindOptions(w,r)

				if string.ToLower(r.Method) == "post" {
					data := strings.TrimSpace(strings.ToLower(r.FormValue("data")))

					if data != "forget"{
						fmt.Fprintf(w, "%s", let.JsonVal)	
					}else{
						fmt.Fprintf(w, "")
					}
				}
			}else{
				loadbranch.Http(w,r,let.Urls)
				// tg1x
				dingtag.Http(w,r,let.Urls)
				bussiness.Http(w,r,let.Urls)
			}
		}

	}else{
		fmt.Fprintf(w, "Nothung here bro")
	}


}