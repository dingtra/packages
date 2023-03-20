package linkpreview


import (
	"fmt"
	"net/http"
	"golang.org/x/net/html"
)

func Previewss() {
	// Fetch the HTML of the webpage.
	resp, err := http.Get("https://www.percona.com/blog/cloud-of-serfdom-or-cloud-of-freedom-what-would-you-choose/?utm_campaign=pzcampaign&utm_source=twitter&utm_medium=paidsocial&utm_content=blog&twclid=2-1plpzm37y7tnradpo5ycg2hww")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	// Parse the HTML.
	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Find all the <img> tags in the HTML.
	var imageLinks []string
	var findImageLinks func(*html.Node)
	findImageLinks = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "img" {
			for _, a := range n.Attr {
				if a.Key == "src" {
					imageLinks = append(imageLinks, a.Val)
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			findImageLinks(c)
		}
	}
	findImageLinks(doc)

	// Print the image links.
	for _, link := range imageLinks {
		fmt.Println(link)
	}
}
