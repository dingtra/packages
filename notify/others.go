package notify

import (
	"net/http"
	"strings"
	"html/template"
)

type NotifyStruct struct {
	Urls []string
	Details template.HTML
	JsonVal string
	Usersid string
}



func (d *NotifyStruct ) Route(r *http.Request) {

	url := strings.Split(r.URL.Path[len("/notify/xr0x7/"):], "/")
	for _, k := range url {

		if k != ""{
			d.Urls = append(d.Urls, strings.ToLower(k))
		}
	}

}