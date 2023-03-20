package makeding



import (
	"net/http"
	"html/template"
	"encoding/json"
	"strings"
)


func (let *MakeDingStruct ) BranchLookUpForm (r *http.Request) {
	Error := map[string]template.HTML{}

	if strings.ToLower(r.Method) == "post" {
		// the aim of branch lookup is to show user a form
		// where they can look up for any open branch to connect with
		// this happens when user clicks the toggle icon at the connect branch div
	
		// generating a form for branch lookups
		Error["nwdngmore092"] =  template.HTML(
			`
			    <div class="brnch-form-class"><div class="brnch-form-class-ch">
				    <div style='display:flex;' id='brnch82791-title'><span>Look for an existing branch</span> <div id='refxr09' url='lkbrx1' style='padding-left:10px;'><i style='pointer-events:none;' data-feather='refresh-cw' color='black'></i></div> </div>
					<div style="padding:10px 0px 10px 0px;" class="brnchlkupxr090">
						<span> <input style="width:100%;height:30px;padding-left:10px;border-radius:4px;border:1px solid #ccc;" type="text" placeholder="look up branch" id="lkbrnchxr09x" />
					</div>
	
					<div class="brnchlkupxr091">
						<div id="lkupxr09x"></div>
					</div>
	
					<div style="font-size:14px;" class="brnchlkupxr092">
						<div><span>Connect to external branch only for open branches <a href="#">learn more</a></span></div>
					</div>
				</div></div>
				<script>
				feather.replace()
			    </script>
			`)
	}
	


	jsn, _ := json.Marshal(Error)
	let.JsonVal = string(jsn)
}

