package model_image

import (
	"math/rand"
	"time"
	"strings"
	"fmt"
)

var (
	Random_image  = []string{"imd-dings08890", "imd-dings08891", "imd-dings08892", "imd-dings08893", "imd-dings08894", "imd-dings08895"}
)

func ImagesMap () map[string]map[string]string {
	images := map[string]map[string]string{}
	
	images["imd-dings08890"] = map[string]string{"first_image":"2345455nd466-dings.png", "second_image":"23456k9jh0opk-dings.png", "third_image":"23456789hsgsg788-dings.png", "fourth_image":"23456789hs902op-dings.png"}
	images["imd-dings08891"] = map[string]string{"first_image":"23456h3190hgddd-dings.png", "second_image":"2345uuu75fewbdi-dings.png", "third_image":"234567hu8954di-dings.png", "fourth_image":"jdjd783ht48303ejd-dings.png"}
	images["imd-dings08892"] = map[string]string{"first_image":"234567890478-dings.png", "second_image":"23456789988dhdjifhk38393-dings.png", "third_image":"jk3e3y893yrh3fd334ufm-dings.png", "fourth_image":"8979yuho98e09udeguydh-dings.png"}
	images["imd-dings08893"] = map[string]string{"first_image":"99juy76463gdy3-dings.png", "second_image":"99juy7ff930n45n6-dings.png", "third_image":"99juy76463sseei3o3-dings.png", "fourth_image":"99juy764uuwr56dswb3-dings.png"}
	images["imd-dings08894"] = map[string]string{"first_image":"et4657egb05bgir-dings.png", "second_image":"et46jk8493jjendkfr-dings.png", "third_image":"et4657widkfpeir-dings.png", "fourth_image":"et46we677hgb05bgir-dings.png"}
	images["imd-dings08895"] = map[string]string{"first_image":"hd74jhfy56e9fh59gj4y-dings.png", "second_image":"et46jk8493jjw5yg4h-dings.png", "third_image":"et46jk849390w2tdb3yd-dings.png", "fourth_image":"et46jk84939edorut46wh1-dings.png"}

	return images
}

func Model_All_Images () string {
	images := ImagesMap()

	// Initialize the random number generator with the current time
	rand.Seed(time.Now().UnixNano())


	
	// Generate a random index within the range of the slice
	index := rand.Intn(len(Random_image))
 
	// Retrieve the element at the random index
	randomElement := Random_image[index]

	if len(randomElement) > 0 {

		fmt.Println(randomElement)
		scrape := []string{}

		for k, _ := range images[randomElement] {
			scrape = append(scrape, k)

		}
		fmt.Println(scrape)

		fill := map[string]string{}
		mainfill := []string{}
		i := 0

		// Seed the random number generator outside the loop
		// rand.Seed(time.Now().UnixNano())

		var OriginalImage string
		
		for len(fill) < len(scrape) {
			indexscrape := rand.Intn(len(scrape))
			if len(fill[scrape[indexscrape]]) <= 0 {
				fill[scrape[indexscrape]] = "new"
				mainfill = append(mainfill, "<div style='padding:1px;' class='img_jss40090 img_jss40090_class' %%% data='"+randomElement+"/"+images[randomElement][scrape[indexscrape]]+"'><img style='pointer-events:none;' src='/dingtra-web-assets/verify_images/"+randomElement+"/"+images[randomElement][scrape[indexscrape]]+"'></div>")
				OriginalImage = "<img style='pointer-events:none;' src='/dingtra-web-assets/verify_images/"+randomElement+"/main_img_tool.png'>"
			}
			i++
		}

		if len(mainfill) == 4 {
			data := `
			   <div class="err_puzzle890">
			   <div class="verify_img_parent">
			        <div class="verify_img_puzzle">
					    <div class="verify_img_puzzle_jsss090">
					        <span>Complete the puzzle to continue</span>
						</div>
						<div class="verify_img_puzzle_jsss091">
						    <div class="puzzle_jss800">
								`+OriginalImage+`
								<div class="verify_img_puzzle_jsss091_check">
									<i data-feather="check"></i>
								</div>
							</div>
						</div>
					</div>

					<div class="img_main_jss3490">
				        <div class="img_main_jss3491">
					        `+strings.Join(mainfill, "")+`
					    </div>
				    </div>

					<div class="verify_img_jss094">
						<div class="verify_img_jss0941">
							<div class="verify_imgjss080" data="first_image">
								<div style="pointer-events:none;" class="verify_imgjss081">
								    <div style="width:fit-content;margin:auto;position:relative;top:35px;font-size:14px;">
								        Click here
									</div>
								</div>
							</div>
							<div class="verify_imgjss080" data="second_image">
								<div style="pointer-events:none;" class="verify_imgjss081">
								    <div style="width:fit-content;margin:auto;position:relative;top:35px;font-size:14px;">
								        Click here
									</div>
								</div>
							</div>
						</div>
						<div class="verify_img_jss0941">
							<div class="verify_imgjss080" data="fourth_image">
								<div style="pointer-events:none;" class="verify_imgjss081">
								    <div style="width:fit-content;margin:auto;position:relative;top:35px;font-size:14px;">
								        Click here
									</div>
								</div>
							</div>
							<div class="verify_imgjss080" data="third_image">
								<div style="pointer-events:none;" class="verify_imgjss081">
								    <div style="width:fit-content;margin:auto;position:relative;top:35px;font-size:14px;">
								        Click here
									</div>
								</div>
							</div>
						</div>
					</div>
					
					<div class="verify_done_jss700">
					</div>
				</div>
				</div>
			`

			return strings.ReplaceAll(data, "%%%", "class_data='img_jss40090_class'")

		}else if len(mainfill) == 3 {
			data := `
			    <div class="verify_img_jss093">
					<div class="verify_img_jss0931">
					    <div class="verify_imgjss080" data="second_div_content">
						    <div style="pointer-events:none;" class="verify_imgjss081">
							</div>
						</div>
						<div class="verify_imgjss080" data="third_div_content">
						    <div style="pointer-events:none;" class="verify_imgjss081">
							</div>
						</div>
					</div>

					<div class="verify_img_jss0932">
					    <div class="verify_imgjss080 verify_imgjss080_third" data="first_div_content">
						    <div style="pointer-events:none;" class="verify_imgjss081">
							</div>
						</div>
					</div>
				</div>

				<div class="img_main_jss3490">
				    <div class="img_main_jss3491">
					    `+strings.Join(mainfill, "")+`
					</div>
				</div>
			`

			return strings.ReplaceAll(data, "%%%", "class_data='img_jss40090_third'")
		}
		
	}


	return ""
}


