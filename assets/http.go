package assets


import (
	"net/http"
	"path"
	"strings"
	"fmt"
)


func Http(w http.ResponseWriter, r *http.Request) {
	sort := []string{}

	brk :=  strings.Split(r.URL.Path,"/")

	for _, item := range brk {
		if len(item) > 0 {
			sort  = append(sort, item)
		}
	}

	if len(sort[0]) > 0 {
		root := "./"+sort[0]

		urls := strings.Split(r.URL.Path[len("/"+sort[0]+"/"):], "/")
		
		links := []string{}
	
		for _, k := range urls {
			if k != ""{
				links = append(links, k)
			}
		}
		
		restrictarea := map[string]bool{"svg":true}
	
		if len(links) >= 2 {

			if restrictarea[links[0]] == true {
				http.ServeFile(w, r, path.Join(root, "/"+strings.Join(links, "/")))
			}else{

				if len(links) >= 3 {
					http.ServeFile(w, r, path.Join(root, "/"+strings.Join(links, "/")))
				}else{
					fmt.Fprintf(w, "Cant process this page!!")
				}
			}
		}else{
			fmt.Fprintf(w, "Cant process this page!!")
		}	

	}else {
		fmt.Fprintf(w, "Cant process this page!!")
	}	
}