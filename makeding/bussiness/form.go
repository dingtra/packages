package bussiness


import (
	"net/http"
	"strings"
)


func GenerateForm (r *http.Request) string {
	name := strings.TrimSpace(r.FormValue("name"))
	filenameid := strings.TrimSpace(strings.ToLower(r.FormValue("id")))

	details := []string{}


	// parent 
	details = append(details, `
	    <div style="position:fixed;top:0px;left:0px;right:0px;top:0;bottom:0;background:rgba(0, 0, 0, 0.5);z-index:500;">
	    <div class="bssnssxr090"> <div class="bssnssxr090-chld">
		    
	        <div style="display:flex;justify-content:end;" class="imgxr090cncl" x='`+filenameid+`'>
			    <button style="pointer-events:none;"><i data-feather='x' color='black'></i></button>
			</div>
	`)

	// parent 
	details = append(details, `
	   `+strings.TrimSpace(r.FormValue("name"))+`:main
	`)

	// title
	details = append(details, `
	    <div class="bssnssxr091">
	        <div class="bssnssxr092"><span>Title</span></div>
			<div class="bssnssxr093">
			    <span><input type="text" name='`+name+`title' placeholder="Enter title"/> </span>
			</div>
	    </div>
	`)

	// price
	details = append(details, `
	    <div class="bssnssxr091">
	        <div class="bssnssxr092"><span>Price</span></div>
			<div style="display: flex;justify-content:space-between;" class="bssnssxr093">
			    <div class="bssnssxr093-select">
					<select name='`+name+`currency'>
						<option name="naira">₦</option>
						<option name="dollar">$</option>
					</select>
				</div>
				<div class="bssnssxr093-box">
					<span><input type="number" name='`+name+`price' placeholder="Enter price"/> </span>
				</div>
			</div>
	    </div>
	`)


	// Discount
	details = append(details, `
	    <div class="bssnssxr091">
	        <div class="bssnssxr092"><span>Add Discount % off</span></div>
			<div class="bssnssxr093-discount">
			    <span><input type="number" name='`+name+`discount' placeholder="Add discount"/> </span>
			</div>
	    </div>
	`)


	// parent 
	details = append(details, `
	   </div></div>
	   </div>
	`)

	details = append(details,
		`
		<script>
		  feather.replace()
		</script>
	`)

	return strings.Join(details, "")
}