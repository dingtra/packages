package linkpreview


import (
	"fmt"
	"net/http"
	"io/ioutil"
	"regexp"
	"strings"
)

func Preview(url string ) string {

	fmt.Println(url, " POLPOL")

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	displaylink := []string{}

	if len(string(html)) > 0 {
		// Use a regular expression to extract the page title
		titleRegex := regexp.MustCompile(`<title>(.*?)</title>`)
		title := titleRegex.FindStringSubmatch(string(html))
	
		// Use a regular expression to extract the page description
		descriptionRegex := regexp.MustCompile(`<meta name="description" content="(.*?)"`)
		
		description := descriptionRegex.FindStringSubmatch(string(html))
		
	
		// Use a regular expression to extract the page image
		imageRegex := regexp.MustCompile(`<meta property="og:image" content="(.*?)"`)
		image := imageRegex.FindStringSubmatch(string(html))
	
		if len(image) > 0 || len(title) > 0 || len(description) > 0 {
	
			displaylink = append(displaylink, "<div class='linkxr8790'><div class='linkxr8790-chld'>")
	
			if len(image) > 0 {
				stream :=  CheckIfLinkIsVideoStreaming(url)

				if stream == true {
					displaylink = append(displaylink, `
					<div style="position:relative;" class='linkxr8791'>
					    <span> <img src='`+image[1]+`' /> </span>
						<div style="position:absolute;top:35%;left:40%;" class="linkxr8791-strm">
						    <svg style="height:60px;width:60px;color:white;fill:white;" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true"><path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zM9.555 7.168A1 1 0 008 8v4a1 1 0 001.555.832l3-2a1 1 0 000-1.664l-3-2z" clip-rule="evenodd"></path></svg>
						</div>
					</div>
					`)			
				}else{
					displaylink = append(displaylink, "<div class='linkxr8791'> <span> <img src='"+image[1]+"' /> </span> </div>")	
				}
			}
	
			// 
			displaylink = append(displaylink, "<div class='linkxr8792'><div class='linkxr8792-chld'>")
			// website link
			displaylink = append(displaylink, "<div class='linkxr8793'> <span> "+url+" </span> </div>")
			
			if len(title) > 0 {
				displaylink = append(displaylink, "<div class='linkxr8794'> <span> "+title[1]+" </span> </div>")
			}
	
			if len(description) > 0 {
				displaylink = append(displaylink, "<div class='linkxr8795'> <span> "+description[1]+" </span> </div>")
			}
			
			displaylink = append(displaylink, "</div></div>")
			// 
	
			displaylink = append(displaylink, "</div></div>")
		}

	}


	return strings.Join(displaylink, "")
}


func CheckIfLinkIsVideoStreaming (url string ) bool {
	ext := map[string]bool{}
	streams := []string{"youtube"}
	answer :=  false

	if len(url) > 0 {
		for _, item := range strings.Split(url, ".") {
			ext[item] = true
		}


		for _, item := range streams {

			if ext[item] == true {
				answer = true
			}
		}
	}

	return answer
}