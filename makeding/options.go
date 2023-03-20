package makeding


import (
	"net/http"
	"strings"
	"html/template"
	"encoding/json"
)


func (let *MakeDingStruct ) FindOptions (w http.ResponseWriter, r *http.Request ){
	Error := map[string]template.HTML{}

	if strings.ToLower(r.Method) == "post" {

		data := strings.TrimSpace(strings.ToLower(r.FormValue("data")))

		if len(data) > 0 {

			if data == "bussiness" {
				Error["flxkr09clo0"] = template.HTML(BranchOptions("bussiness")+let.BusinessFiles())

			}else if data == "polls" {
				Error["flxkr09clo0"] = template.HTML(BranchOptions("polls")+let.Polls())

			}else if data == "files" {
				Error["flxkr09clo0"] = template.HTML(BranchOptions("files")+let.GetFiles())
			}
		}
	}

	json, _ := json.Marshal(Error)

	let.JsonVal = string(json)
}

func BranchOptions (name string) string {
	details := []string{}


	// options
	details = append(details, "<div class='nwdngoptns090'>")

	if name == "files" {
		// 
		details = append(details, 
			`<div id="xl09br5409" x="files" class="nwdngoptns091hr" style="border:unset;">
				<div style="pointer-events:none;" class="nwdngoptns091">
					<div class="nwdngoptns091-svg">
						<span><i style="color:green;" data-feather="file-plus" color="black"></i></span>
					</div>
					
					<div class="nwdngoptns091-title">
						<span style="color:green;">Add files</span>
					</div>
				</div>
			</div>
		`)
		//
	}else{
		// 
		details = append(details, 
			`<div id="xl09br5409" x="files" class="nwdngoptns091hr" style="border:unset;">
				<div style="pointer-events:none;" class="nwdngoptns091">
					<div class="nwdngoptns091-svg">
						<span><i data-feather="file-plus" color="black"></i></span>
					</div>
					
					<div class="nwdngoptns091-title">
						<span>Add files</span>
					</div>
				</div>
			</div>
		`)
		//
	}

	if name == "polls" {
		// 
		details = append(details, 
			`<div id="xl09br5409" x="polls" class="nwdngoptns091hr">
				<div style="pointer-events:none;" class="nwdngoptns091">
					<div class="nwdngoptns091-svg">
						<span><i style="color:green;" data-feather='bar-chart' color='black'></i></span>
					</div>
					
					<div class="nwdngoptns091-title">
						<span style="color:green;">New poll</span>
					</div>
				</div>
			</div>
		`)
		//
	}else{
		// 
		details = append(details, 
			`<div id="xl09br5409" x="polls" class="nwdngoptns091hr">
				<div style="pointer-events:none;" class="nwdngoptns091">
					<div class="nwdngoptns091-svg">
					    <span><i data-feather='bar-chart' color='black'></i></span>
					</div>
					
					<div class="nwdngoptns091-title">
						<span>New poll</span>
					</div>
				</div>
			</div>
		`)
		//
	}

	if name == "bussiness" {
		// 
		details = append(details, 
			`<div id="xl09br5409" x="bussiness" class="nwdngoptns091hr">
				<div style="pointer-events:none;" class="nwdngoptns091">
					<div class="nwdngoptns091-svg">
						<span><i style="color:green;" data-feather='trending-up' color='black'></i></span>
					</div>
					
					<div class="nwdngoptns091-title">
						<span style="color:green;">Bussiness</span>
					</div>
				</div>
			</div>
		`)
		//
	}else{
		// 
		details = append(details, 
			`<div id="xl09br5409" x="bussiness" class="nwdngoptns091hr">
				<div style="pointer-events:none;" class="nwdngoptns091">
					<div class="nwdngoptns091-svg">
					    <span><i data-feather='trending-up' color='black'></i></span>
					</div>
					
					<div class="nwdngoptns091-title">
						<span>Bussiness</span>
					</div>
				</div>
			</div>
		`)
		//
	}
	
	// show more || less
	details = append(details, 
		`<div id="xl09bjj090" x="survey" class="nwdngoptns091hr">
			<div style="pointer-events:none;" class="nwdngoptns091">
				<div class="nwdngoptns091-svg">
					<span><i data-feather='chevrons-down' color='black'></i></span>
				</div>
				
				<div class="nwdngoptns091-title">
					<span>Less</span>
				</div>
			</div>
		</div>
	`)
	//


	details = append(details,
	`
	<script>
	  feather.replace()
    </script>
	`)
	details = append(details, "</div>")
	// options


	return strings.Join(details, "")
}