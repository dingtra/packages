package notify

import (
	"net/http"
	"rundb"
	"strings"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"html/template"
	"log"
	"fmt"
)


func Http(w http.ResponseWriter, r *http.Request) {
// 	all mens are ...
	session, _ := rundb.Store.Get(r, "session")
	let := NotifyStruct{}

	if session.Values["username"] != nil {
		let.Route(r)
		let.Usersid = session.Values["usersid"].(string)

		if len(let.Urls) > 0 {

			if let.Urls[0] == "xr1"{
				let.Details = template.HTML(GetNotification(session.Values["usersid"].(string)))
				let.MarkAllRead(r)
				fmt.Fprintf(w, "%s", `
					<html>

					<head>
						<title>Home Page</title>
						<link rel="icon" href="/jskl899nxhsjas/" type="image/x-icon">
						<meta charset="UTF-8">
						<meta name="viewport" content="width=device-width, initial-scale=1.0" />
						<link rel="stylesheet" href="/dingtra-web-assets/css/cssntfy001/ntfy9013res80.css"/>
						<script src="/dingtra-web-assets/js/feather/feather.js"></script>
						<script src="/dingtra-web-assets/js/jquery.js"></script>
					
					</head>
					
					<body >
						`+
						let.Details+
						`
						<script>
							feather.replace()
						</script>
					</body>
					
					</html>
					`)

			}else if let.Urls[0] == "cuntxr"{
				if strings.ToLower(r.Method) == "post" {
					keys := strings.TrimSpace(r.FormValue("count"))
					if len(keys) > 0 {
						fmt.Println("counter AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA ", CountNotify(keys), keys)
						// /notify/xr0x7/?count=id
						fmt.Fprintf(w, "%s", CountNotify(keys))
					}else{
						http.Redirect(w, r, "/app/xrio8363/login/", http.StatusSeeOther)
					}
				}
			}
		}else{
			let.Details = template.HTML(GetNotification(session.Values["usersid"].(string)))
			
			keys, ok := r.URL.Query()["default"]
				
			if !ok || len(keys[0]) < 1 {
				log.Println("Url Param 'key' is missing")
			}
	
			if len(keys) > 0 {
				key := keys[0]
				
	
				if key == "unread" { 
					let.Details = template.HTML(GetUnreadNotification(session.Values["usersid"].(string)))
					fmt.Fprintf(w, "%s", `
					<html>

					<head>
						<title>Home Page</title>
						<link rel="icon" href="/jskl899nxhsjas/" type="image/x-icon">
						<meta charset="UTF-8">
						<meta name="viewport" content="width=device-width, initial-scale=1.0" />
						<link rel="stylesheet" href="/dingtra-web-assets/css/cssntfy001/ntfy9013res80.css"/>
						<script src="/dingtra-web-assets/js/feather/feather.js"></script>
						<script src="/dingtra-web-assets/js/jquery.js"></script>
					
					</head>
					
					<body >
						`+
						let.Details+
						`
						<script>
							feather.replace()
						</script>
					</body>
					
					</html>
					`)
				}else{
					fmt.Fprintf(w, "%s", `
					<html>

					<head>
						<title>Home Page</title>
						<link rel="icon" href="/jskl899nxhsjas/" type="image/x-icon">
						<meta charset="UTF-8">
						<meta name="viewport" content="width=device-width, initial-scale=1.0" />
						<link rel="stylesheet" href="/dingtra-web-assets/css/cssntfy001/ntfy9013res80.css"/>
						<script src="/dingtra-web-assets/js/feather/feather.js"></script>
						<script src="/dingtra-web-assets/js/jquery.js"></script>
					
					</head>
					
					<body >
						`+
						let.Details+
						`
						<script>
							feather.replace()
						</script>
					</body>
					
					</html>
					`)
				}
			}else{
				fmt.Fprintf(w, "%s", `
					<html>

					<head>
						<title>Home Page</title>
						<link rel="icon" href="/jskl899nxhsjas/" type="image/x-icon">
						<meta charset="UTF-8">
						<meta name="viewport" content="width=device-width, initial-scale=1.0" />
						<link rel="stylesheet" href="/dingtra-web-assets/css/cssntfy001/ntfy9013res80.css"/>
						<script src="/dingtra-web-assets/js/feather/feather.js"></script>
						<script src="/dingtra-web-assets/js/jsquery/jquery.js"></script>
						<script src="/dingtra-web-assets/js/ntfy0xr/ntfxr090.js"></script>
					
					</head>
					
					<body >
						`+
						let.Details+
						`
						<script>
							feather.replace()
						</script>
					</body>
					
					</html>
				`)
			}
		}
		
	
	}
}
