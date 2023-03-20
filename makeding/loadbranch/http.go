package loadbranch


import (
	"net/http"
	"rundb"
	"fmt"
)


func Http( w http.ResponseWriter, r *http.Request, url []string ){
	let := AjaxStruct{}
	session, _ := rundb.Store.Get(r, "session")
	let.Usersid = session.Values["usersid"].(string)

	if url[0] == "xk1" {
		let.VerifyQuery(r)
		fmt.Fprintf(w, "%s",let.JsonVal)
	}
}