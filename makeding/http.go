package makeding


import (
	"net/http"
	"rundb"
	"github.com/dingtra/packages/makeding/loadbranch"
	"github.com/dingtra/packages/makeding/dingtag"
	"fmt"
)


func Http (w http.ResponseWriter, r *http.Request ){
	let := MakeDingStruct{}
	let.Route(r)

	session, _ := rundb.Store.Get(r, "session")
	w.Header().Set("Content-Type", "text/html")

	fmt.Println(let.Urls)
	

	if session.Values["username"] != nil && session.Values["usersid"] != nil {
		
		if len(let.Urls) == 0 {
			fmt.Println(" NAAAAA here")
			fmt.Fprintf(w, "%s",let.NewDing())

		}else{

			if let.Urls[0] == "flx0r" {
				let.GetFiles()
				fmt.Fprintf(w, "%s", let.JsonVal)

			}else if let.Urls[0] == "sk0x" {
				let.GetFiles()
				fmt.Fprintf(w, "%s", let.JsonVal)
			}else if let.Urls[0] == "sk1x" {
				let.Polls()
				fmt.Fprintf(w, "%s", let.JsonVal)
			}else{
				loadbranch.Http(w,r,let.Urls)
				dingtag.Http(w,r,let.Urls)
			}
		}

	}else{
		fmt.Fprintf(w, "Nothung here bro")
	}


}