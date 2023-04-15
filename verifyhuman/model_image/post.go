package model_image



import (
	"net/http"
	"strings"
	"rundb"
	// "strconv"
	"encoding/json"
	"fmt"
)

// data := map[string]string{"first_image":"imd-dings08894/et4657egb05bgir-dings.png"}
// data := map[string]string{"second_image":"imd-dings08894/et46jk8493jjendkfr-dings.png"}
// data := map[string]string{"third_image":"imd-dings08894/et46we677hgb05bgir-dings.png"}
// data := map[string]string{"fourth_image":"imd-dings08894/et4657widkfpeir-dings.png"}


func (let *ModelStruct) CheckPuzzle (w http.ResponseWriter, r *http.Request) {
	
	Error := map[string]string{}

	if strings.ToLower(r.Method) == "post" {
		files :=  strings.TrimSpace(r.FormValue("files"))
		var GetImage_directory string

		data := make(map[string]string)
		if(len(files) > 0) {
			// files, _ = strconv.Unquote(files)
	
			err := json.Unmarshal([]byte(files), &data)
			if err != nil {
				fmt.Println(err)
			}
		}

		if len(data) > 0 {
			if len(data["first_image"]) > 0 {
				getfirst_data := strings.Split(data["first_image"],"/")

				if len(getfirst_data) > 0 {
					GetImage_directory = getfirst_data[0]
				}
			}
		}

		GetImages := ImagesMap()[GetImage_directory]

		counter := 0

		if len(GetImages) > 0 {
			for k, v := range data {
				getsecond_data := strings.Split(v,"/")

				if len(getsecond_data) > 1 {
					if GetImages[k] == getsecond_data[1] {
						counter += 1
					}
				}
			}
		}


		if counter < 4 {
			Error[".err_puzzle890"] = Model_All_Images()
		} else{
			session, _ := rundb.Store.Get(r, "session")
			session.Values["verify_human"] = "true"
			session.Save(r, w)
			
			if session.Values["fullname"] != nil && len(session.Values["fullname"].(string)) > 0 {
				let.Username = session.Values["fullname"].(string)
			}else{
				let.Username = session.Values["username"].(string)
			}
			// success
			Error[".err_puzzle890"] = `
				<div class="welcome_activity_btn">
					<div class="welcome_jss09kl090">
						<span class="welcome_span_090">Welcome back <b style="text-transform:capitalize;">`+let.Username+`</b></span>
					</div>
					<div class="welcome_jss09kl091">
						<div class="welcome_svg_jss090">
							<span> <i data-feather="check"></i> </span>
						</div>
			
						<div class="welcome_svg_jss091">
							<span> Authenticated already</span>
						</div>
					</div>
			
					<div class="welcome_jss09kl092">
						<div class="welcome_verify_jss090">
							<span> <input type="checkbox" checked> </span>
						</div>
			
						<div class="welcome_verify_jss091">
							<span>Verified </span> 
						</div>

					</div>
					<div class="welcome_continue_class">
					    <a style="text-decoration:none;color:unset;" href="/">
					    <div class="welcome_continue">
							<div class="welcome_continue_jss090">
								<span>Continue</span>
							</div>
						</div>
						</a>
					</div>
				</div>
			`

			// Model_All_Images()
		}

	}

	json, _ := json.Marshal(Error)
	let.Details = string(json)
}