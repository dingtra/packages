package makeding


import (
	"net/http"
	"rundb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
	"fmt"
)


func  (let *MakeDingStruct ) SharedDing (r *http.Request )  string {
	details := []string{}

	if strings.ToLower(r.Method) == "post" {

		conns := rundb.Connect("app").Collection("dings")
		postid := strings.TrimSpace(r.FormValue("id"))
		theid, _ := primitive.ObjectIDFromHex(postid)
		find := rundb.FindOne(conns, bson.M{"_id":theid})
		fmt.Println(theid, "CJEKCD ",find)	
	
		if len(find) > 0 {
			// shared post available
	
			// parent 
			details = append(details, "<div style='position:fixed;top:100px;z-index:400;left:0;right:0;' class='nwdngxr090'><div style='background:white;border:1px solid #ccc;border-radius:4px;box-shadow: 0 3px 10px rgb(0 0 0 / 0.1);padding:20px 20px 20px 20px;border-radius:4px;width:fit-content;margin:auto;' class='nwdngxr090-chld'>")
			details = append(details, "<div class='erx090r'></div>")
			details = append(details, "<div id='ntfyxr0k87cncl'><div style='pointer-events:none;display:flex;justify-content:end;padding:5px 10px 10px 0px;'><button><i data-feather='x' color='black'></i></button> </div></div>")
			details = append(details, "<form enctype='multipart/form-data' name='dngfrm'>")
			// phase
			details = append(details, "<div class='phsxr09x1'>")
		
			// phase one
			details = append(details, "<div style='width:100%;' class='nwdngxr090-phase0x90'>")
		
			details = append(details, "<div style='width:fit-content;margin:auto;padding:10px 0px 10px 0px;' class='nwdngxr090-dscrptn'><span>Write something</span></div>")
		
			// description
			details = append(details, "<div style='padding:0px 0px 10px 0px;' class='nwdngxr091-textr09xyz'>")
			details = append(details, "<div class='nwdngxr091'><span> <textarea style='width:100%;padding:10px 5px 10px 5px;border:1px solid #ccc;border-radius:4px;resize:none;' id='frmpst27890editable' name='about' placeholder='What&#39;s the ding all about?'></textarea> </span></div>")
		
			// 
			details = append(details, "<div style='display:none;' class='nwdngxr091-textr09x'><div class='nwdngxr091-textr09x-chld'>  <div id='nwdngxr091-text092x'>0/1000</div> </div></div>")
		
			// 
		
			details = append(details, "</div>")
			// description
		
		
			// srch details
			details = append(details, "<div class='nwdngxr092'>")
			details = append(details, "</div>")
			// srch details
	
			// submit
			details = append(details, "<div style='display:block;' id='frmpst27890submtbtn'><div id='frmpst27890submt' class='frmpst27890submtbtn'><div><div style='pointer-events:none;' class='frmpst27890submt'><div id='frmpst27890submt'><span>Reding</span></div></div> </div></div></div>")
		
			details = append(details, "</div>")
			// phase one
		
			details = append(details, "</div>")
			// phase

			// 00000000000000000000000000000000000000000000000000000000000000000000000000000000
			details = append(details, "<div style='display:none;'><span><input value='"+postid+"' type='text' name='share'/></span></div>")
			// 00000000000000000000000000000000000000000000000000000000000000000000000000000
			
			// 00000000000000000000000000000000000000000000000000000000000000000000000000000000
			details = append(details, "<div style='display:none;'><span><input value='"+find["usersid"].(primitive.ObjectID).Hex()+"' type='text' name='ownr'/></span></div>")
			// 00000000000000000000000000000000000000000000000000000000000000000000000000000
			
		
			details = append(details, "</form>")
		
			details = append(details, "</div></div>")
			// parnt
		
		
			details = append(details, `
			<script>
			  feather.replace()
			</script>
			`)
			
		}else{
			// shared post not available
			// parent 
			details = append(details, "<div class='nwdngxr090'><div class='nwdngxr090-chld'>")
			details = append(details, "<div id='ntfyxr0k87cncl'><div style='pointer-events:none;display:flex;justify-content:end;padding:5px 10px 10px 0px;'><button><i data-feather='x' color='black'></i></button> </div></div>")
			details = append(details, "<h2>We cant find ding, deleted or from your side</h2>")
			details = append(details, "</div></div>")
			// parnt
		}
	}

	return strings.Join(details, "")
}