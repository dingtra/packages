package notify



import (
	"rundb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/bson"
	"strings"
	"strconv"
)


func GetUnreadNotification(id string) string {
	getcoll := rundb.Connect("app").Collection("notification")

	stuff := rundb.QueryJoin(getcoll)
	// fmt.Println("from home")

	countall := []string{}
	countunread := []string{}

	details := []string{}
	final := []string{}

	
	for _, item := range stuff{

		if  item["owner"].(string)  == id {

			//notifications
			countall = append(countall, item["postid"].(string))

			if item["opened"] == nil {
				countunread = append(countunread, item["postid"].(string))

				//notifications
				if item["type"].(string) == "comment" || item["type"].(string) == "reply"  {
					if item["type"].(string) == "comment" {
	
						details = append(details, "<a href='/?box=ding&id="+item["mainid"].(string)+"&notify="+item["_id"].(primitive.ObjectID).Hex()+"#cmmshw"+item["postid"].(string)+"' style='text-decoration:none;color:black;'>")
					}else{
						details = append(details, "<a href='/?box=ding&id="+item["mainid"].(string)+"&notify="+item["_id"].(primitive.ObjectID).Hex()+"#rpshw"+item["postid"].(string)+"' style='text-decoration:none;color:black;'>")
					}
				}
	
				if item["type"].(string) == "like" || item["type"].(string) == "tagger" || item["type"].(string) == "tag" ||  item["type"].(string) == "share" {
					details = append(details, "<a href='/?box=ding&id="+item["postid"].(string)+"&notify="+item["_id"].(primitive.ObjectID).Hex()+"' style='text-decoration:none;color:black;'>")
				}
	
				if item["type"].(string) == "subscribe" {
					details = append(details, "<a href='/"+item["username"].(string)+"&notify="+item["_id"].(primitive.ObjectID).Hex()+"' style='text-decoration:none;color:black;'>")
				} 
				
			
				details = append(details, "<div class='notification-box' style='background-color: #f2f2f2;margin-top:10px;border-radius:4px;'>")
				
				//top
				details = append(details, "<div class='notify-top-details'>")
				
				//image
				details = append(details, "<div id='image'>")
				if item["theid"].(primitive.A)[0].(primitive.M)["pic"] != nil{
					details = append(details, ImageShow(item["theid"].(primitive.A)[0].(primitive.M)["pic"].(string)))
				}else{
					details = append(details, ImageShow(""))
				}
				
				details = append(details, "</div>")
				//image
				
				//name 
				details = append(details, "<div class='notify-name'>")
				details = append(details, "<div id='name'><span>"+item["theid"].(primitive.A)[0].(primitive.M)["username"].(string)+"</span> <div id='notify-date'>"+item["date"].(string)+"</div> </div>")
				
				if item["type"].(string) == "like" {
					details = append(details, "<div id='type-like' style='font-size:13px;'><i style='height:14px;width:14px;position:relative;top:2px;' data-feather='heart' color='red' fill='red'></i> "+item["body"].(string)+"</div>")
				}
	
				if item["type"].(string) == "comment" {
					details = append(details, "<div id='type-comment'><div id='f-title' style='font-size:14px;'> <i style='height:14px;width:14px;position:relative;top:2px;' data-feather='message-circle' color='black' fill='black'></i> "+item["body"].(string)+"</div> <div id='notify-details' style='font-size:13px;'>"+item["comment"].(string)+"</div></div>")
				}
	
				if item["type"].(string) == "reply" {
					details = append(details, "<div id='type-comment'><div id='f-title'  style='font-size:14px;'> <i style='height:14px;width:14px;position:relative;top:2px;' data-feather='message-circle' color='black' fill='black'></i> "+item["body"].(string)+"</div> <div id='notify-details' style='font-size:13px;'>"+item["reply"].(string)+"</div></div>")
				}
	
				if item["type"].(string) == "tag" {
					details = append(details, "<div id='type-tag'><i data-feather='tag' color='black' fill='black'></i>  "+item["body"].(string)+"</div>")
				}
	
				if item["type"].(string) == "follow" {
					details = append(details, "<div id='type-tag' style='font-size:13px;'><i style='height:14px;width:14px;position:relative;top:2px;' data-feather='user' color='black' fill='black'></i>  "+item["body"].(string)+"</div>")
				}
	
				if item["type"].(string) == "share" {
					details = append(details, "<div id='type-tag' style='font-size:13px;'><i style='height:14px;width:14px;position:relative;top:2px;' data-feather='repeat' color='black' fill='black'></i>  "+item["body"].(string)+"</div>")
				}


				details = append(details, "</div>")//name
				
				//top
				details = append(details, "</div>")
				
				details = append(details, "</div>")
				//notifications
				details = append(details, "</a>###")
			}
			
		}
	}
	

	if len(details) > 0  {
		vals := strings.Join(details, "")

		volks := strings.Split(vals, "###")

		
		emmiter := []string{}

		for i := len(volks) -1; i >= 0; i-- {
			if volks[i] != ""{
				emmiter = append(emmiter, volks[i])
			}
		}
		final = append(final, "<div id='ntfy0981'><div class='notification-page'><div class='child-notification-page'>")
		final = append(final, "<div id='ntfyxr0k87cncl'><div style='pointer-events:none;display:flex;justify-content:end;padding:10px 10px 10px 0px;'><button>"+X+"</button> </div></div>")
		final = append(final, "<h3 style='display:flex;justify-content:space-between;'><span>Notifications "+BellSvg+"</span> <span id='no98090'><i style='pointer-events:none;' data-feather='check' color='black' ></i> Mark all as read</span></h3>")

		//btn
		final = append(final, "<div class='btn-details'>")
		if len(countall) > 0 {
			final = append(final, "<div class='notify-span' id='clckntf901'><div id='notify-svg' style='pointer-events:none;'>All Notifications "+strconv.Itoa(len(countall))+"</div></div>")
		}else{
			final = append(final, "<div class='notify-span' id='clckntf901'><div id='notify-svg' style='pointer-events:none;'>All Notifications </div></div>")
		}
		
		if len(countunread) > 0 {
			final = append(final, "<div class='notify-span' style='border-bottom:2px solid blue;' id='clckntf902'><div id='notify-svg' style='color:blue;pointer-events:none;'> Unread "+strconv.Itoa(len(countunread))+"</div></div>")
		}else{
			final = append(final, "<div class='notify-span' style='border-bottom:2px solid blue;' id='clckntf902'><div id='notify-svg' style='color:blue;pointer-events:none;'>Unread </div></div>")
		}
		final = append(final, "</div>")
		//btn

		final = append(final, "<div style='margin-top:15px;'><div class='view-all-notify'>"+strings.Join(emmiter, "")+"</div></div>")

		final = append(final, "</div></div></div>")

	}else{
		final = append(final, "<div class='notification-page' style='border: 1px solid rgba(var(--ce3,239,239,239),1);'><div class='child-notification-page'>")
		final = append(final, "<div id='ntfyxr0k87cncl'><div style='pointer-events:none;display:flex;justify-content:end;padding:10px 10px 10px 0px;'><button>"+X+"</button> </div></div>")
		final = append(final, "<h3 style='display:flex;justify-content:space-between;'><span>Notifications "+BellSvg+"</span> </h3>")

		//btn
		final = append(final, "<div class='btn-details'>")
		if len(countall) > 0 {
			final = append(final, "<div class='notify-span' id='clckntf901'><div id='notify-svg' style='pointer-events:none;'>All Notifications "+strconv.Itoa(len(countall))+"</div></div>")
		}else{
			final = append(final, "<div class='notify-span' id='clckntf901'><div id='notify-svg' style='pointer-events:none;'>All Notifications </div></div>")
		}
		
		if len(countunread) > 0 {
			final = append(final, "<div class='notify-span' style='border-bottom:2px solid blue;' id='clckntf902'><div id='notify-svg' style='color:blue;pointer-events:none;'> Unread "+strconv.Itoa(len(countunread))+"</div></div>")
		}else{
			final = append(final, "<div class='notify-span' style='border-bottom:2px solid blue;' id='clckntf902'><div id='notify-svg' style='color:blue;pointer-events:none;'> Unread </div></div>")
		}
		final = append(final, "</div>")
		//btn

		
		final = append(final, "<div style='width:fit-content;margin:auto;' class='g239'><h2>Nothing to show here</h2></div>")

		final = append(final, "</div></div>")
	}

	return "<div style='position:fixed;top:0;bottom:0;left:0;right:0;background: rgba(0, 0, 0, 0.5);z-index:200;'>"+strings.Join(final, "")+"</div> <script>feather.replace()</script>"
}