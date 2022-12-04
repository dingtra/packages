package dingtag

import (
	"net/http"
	"rundb"
	"fmt"
)


func Http (w http.ResponseWriter, r *http.Request, url []string) {
	let := DingtagStruct{}
	session, _ := rundb.Store.Get(r, "session")
	let.Usersid = session.Values["usersid"].(string)

	if url[0] == "tg0x" {
		let.TagName(r)
		fmt.Fprintf(w, "%s",let.JsonVal)
	}
}