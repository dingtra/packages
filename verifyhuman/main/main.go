package main


import (
	"net/http"
	"model_image"
	"github.com/dingtra/packages/assets"
	"log"
	"fmt"
)



func main (){
	AssetsRoutes := []string{"/verify-assets/", "/branch/"}

	for _, item := range AssetsRoutes {
		http.HandleFunc(item, assets.Http)
	}
	http.HandleFunc("/model_img/", model_image.Http)
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, `
		<html>
		<head>
		<meta name="viewport" content="width=device-width, initial-scale=1.0" />
		<link rel="stylesheet" href="/verify-assets/css/hm090/csshm090.css" />
		<script src="/verify-assets/js/jsquery/jquery.js"></script>
		<script src="/verify-assets/js/feather/feather.js"></script>
		</head>
		<body>
		%s
		<script>
		localStorage.clear();
		document.addEventListener("click", function(e){
			if(e.target.classList.contains("verify_imgjss080")){
				var data = e.target.getAttribute("data");

				// get json-files for puzzle verification
				var json = JSON.parse(localStorage.getItem("files"))
				if(json != undefined){
					// checks if it's not empty
					// when this blocks happens it means a user has cliked 
					// a specific block to insert puzzle
					// then remove any key-value associated to that block of div cliked
					// the block of divs can be first_image out of {second_image, third_image}
					// if it's first image then remove it from old array and create a new one
					NewArr = {};

					for (var k in json){
						if(k != data) {
							NewArr[k] = json[k] 
						}
					}

					if(Object.keys(NewArr).length < 4){
						$(".verify_done_jss700").html("");
					}
					localStorage.setItem('files', JSON.stringify(NewArr))
				}

			    var query = document.querySelectorAll(".verify_imgjss080"); 
				query.forEach(function(element) {
			        element.setAttribute("active", data);
					element.style.border = "1px solid #e7e6e6";
		        });
				e.target.style.border="1px solid #b77676";
				const divPlank = document.querySelector("div.verify_imgjss080[data='"+data+"']");

				divPlank.innerHTML = "<div style='width:fit-content;margin:auto;position:relative;top:30px;flex-direction:column;'><div style='padding:5px 0px 5px 0px;font-size:14px;'>Start puzzle</div> <div style='padding:5px 0px 5px 0px;font-size:12px;'>click image</div> </div>";
			}else if(e.target.classList.contains("img_jss40090")){
				var data = e.target.getAttribute("data");
				var query = document.querySelectorAll(".verify_imgjss080"); 
				var check = "";
				query.forEach(function(element) {
					console.log(element.getAttribute("active"))
					if(element.getAttribute("active") != null){
						check = element.getAttribute("active");
					}
				});

				var JsonObj = {}

				if(check.length > 0){

					// get item files
					var json = JSON.parse(localStorage.getItem("files"))
					if(json != undefined){
						// check if it's available 
						// it means it's ready for receiving new items
						// now add the new image clicked json[check] = data to the object json
						json[check] = data

						if(Object.keys(json).length === 4){
							var lenjson = document.querySelector(".verify_done_jss700");
							$(".verify_done_jss700").html("<div style='margin-top:10px;' class='btn_xr090_lk verify_img_btn090'><div style='pointer-events:none;' class='verify_img_btn091'><span>Verify puzzle</span></div></div>");
						}
						
						localStorage.setItem('files', JSON.stringify(json));
					}else{
						// if it's unavailable 
						// it means we have to create a new object and store it
						JsonObj[check] = data
						localStorage.setItem('files', JSON.stringify(JsonObj));
					}

					const divPlank = document.querySelector("div.verify_imgjss080[data='" +check+ "']");
					divPlank.innerHTML = "<div class='"+e.target.getAttribute("class_data")+"' style='pointer-events:none;'><img src='/verify-assets/verify_images/"+data+"'></div>"
				}
				

				// query.forEach(function(element) {
				// 	console.log(element.textContent, " HEHE");
				// });
			}else{
				var query = document.querySelectorAll(".verify_imgjss080"); 
				query.forEach(function(element) {
			        element.setAttribute("active", data);
					element.style.border = "1px solid #e7e6e6";
		        });
			}
		})


		document.addEventListener("click", function(e){
			if(e.target.classList.contains("btn_xr090_lk")){
			  var json = JSON.parse(localStorage.getItem("files"));
			  var puzzle_frm = new FormData(); puzzle_frm.append("files", JSON.stringify(json))
			  fetch("/model_img/puzzle/", {
				method:"POST",
				body:puzzle_frm,
			  }).then(function (resp){
				return resp.json();
			  }).then(function (data){
				for(var k in data){
					$(k).html(data[k]);
				}
			  })
			}
		})

		</script>
		<script>
		feather.replace()  
	   </script>
		</body>
		</html>
		`, model_image.Model_All_Images())
	})
	log.Fatal(http.ListenAndServe(Port(), nil))
}


func Port () string {
	port := ":8080"
	fmt.Println("PORT STARTED AT ", port)

	return port
}