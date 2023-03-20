package linkpreview


import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
)


func CheckForLink (item primitive.M) string {

	var checkfiles bool
	var getsinglelink string

	if (item["image"] != nil && len(item["image"].(string)) > 0 ) ||   (item["video"] != nil && len(item["video"].(string)) > 0 ){
		checkfiles = true
	}

	if checkfiles == false {

		if item["body"] != nil && len(item["body"].(string)) > 0 {

			for _, k := range strings.Split(item["body"].(string), " ") {

				if len(k) > 0 {
					// check for https

					if len(k) > 4 {

						if k[0:5] == "https" {
							getsinglelink = k
						}
					}

					// check for www

					if len(getsinglelink) == 0 {

						if len(k) > 3 {
							if k[0:3] == "www" {
								getsinglelink = "https://"+k
							}
						}
					}
					
				}
			} 
		}
	}

	if len(getsinglelink) > 0 {

		// it means a link has been found 
		// now render link

		// res := Preview(getsinglelink)

		return ""

	}

	return ""
}