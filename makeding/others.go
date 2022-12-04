package makeding


import (
	"net/http"
	"strings"
)


type MakeDingStruct struct {
	Urls []string
	JsonVal string
}


func (let *MakeDingStruct ) Route (r *http.Request) {

	url := strings.Split(r.URL.Path[len("/xoz/qr1/"):], "/")

	for _, k := range url {

		if len(k) > 0 {
			let.Urls = append(let.Urls, k)
		}
	}
}