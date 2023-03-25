package notify

import (
	"rundb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
	"strings"
	"strconv"
)

var (
	X = `<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="black" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="feather feather-x" color="black"><line x1="18" y1="6" x2="6" y2="18"></line><line x1="6" y1="6" x2="18" y2="18"></line></svg>`
	BellSvg = `<svg style='position:relative;top:2px;' xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="feather feather-bell"><path d="M18 8A6 6 0 0 0 6 8c0 7-3 9-3 9h18s-3-2-3-9"></path><path d="M13.73 21a2 2 0 0 1-3.46 0"></path></svg>`
)

func GetNotification (id string) string {
	getcoll := rundb.Connect("app").Collection("notification")
	stuff := rundb.QueryJoin(getcoll)

	rundb.UpdateMany(getcoll, bson.M{"owner":id}, bson.M{"viewed":"count"})

	countall := []string{}
	countunread := []string{}

	details := []string{}
	final := []string{}

	
	for _, item := range stuff{

		if len(item["theid"].(primitive.A)) > 0 && item["owner"].(string)  == id {

			//notifications
			if item["type"].(string) == "comment" || item["type"].(string) == "reply"  {
				if item["type"].(string) == "comment" {

					details = append(details, "<a href='/?comment="+item["postid"].(string)+""+CheckIfNotificationIsUnopened(item)+"&notify="+item["_id"].(primitive.ObjectID).Hex()+"#cmmshw"+item["postid"].(string)+"' style='text-decoration:none;color:black;'>")
				}else{
					details = append(details, "<a href='/?reply="+item["postid"].(string)+""+CheckIfNotificationIsUnopened(item)+"&notify="+item["_id"].(primitive.ObjectID).Hex()+"#rpshw"+item["postid"].(string)+"' style='text-decoration:none;color:black;'>")
				}
			}

			if item["type"].(string) == "comment-tag" {
				details = append(details, "<a href='/?comment="+item["postid"].(string)+""+CheckIfNotificationIsUnopened(item)+"&notify="+item["_id"].(primitive.ObjectID).Hex()+"#cmmshw"+item["postid"].(string)+"' style='text-decoration:none;color:black;'>")
			}

			if item["type"].(string) == "reply-tag" {
				details = append(details, "<a href='/?reply="+item["postid"].(string)+""+CheckIfNotificationIsUnopened(item)+"&notify="+item["_id"].(primitive.ObjectID).Hex()+"#cmmshw"+item["postid"].(string)+"' style='text-decoration:none;color:black;'>")
			}

			// others branch
			if item["type"].(string) == "others branch" || item["type"].(string) == "poll"  {
				details = append(details, "<a href='/?box=ding&id="+item["postid"].(string)+""+CheckIfNotificationIsUnopened(item)+"&notify="+item["_id"].(primitive.ObjectID).Hex()+"#cmmshw"+item["postid"].(string)+"' style='text-decoration:none;color:black;'>")
			}

			if item["type"].(string) == "follow" {
				check := ""
				if item["opened"] == nil {
					check = "?target=true"
				}
				details = append(details, "<a href='/"+item["username"].(string)+check+"&notify="+item["_id"].(primitive.ObjectID).Hex()+"' style='text-decoration:none;color:black;'>")
			}

			if item["type"].(string) == "like" || item["type"].(string) == "tag" ||  item["type"].(string) == "share" {
				details = append(details, "<a href='/?box=ding&id="+item["postid"].(string)+""+CheckIfNotificationIsUnopened(item)+"&notify="+item["_id"].(primitive.ObjectID).Hex()+"' style='text-decoration:none;color:black;'>")
			}

			countall = append(countall, item["postid"].(string))

			if item["opened"] != nil {
				details = append(details, "<div class='notification-box' style='margin-top:10px;'>")
			}else{
				countunread = append(countunread, item["postid"].(string))
				details = append(details, "<div class='notification-box' style='background-color: #f2f2f2;margin-top:10px;border-radius:4px;'>")
			}

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

			if item["opened"] != nil {
				details = append(details, "<div id='name'><span>"+item["theid"].(primitive.A)[0].(primitive.M)["username"].(string)+"</span> <div id='notify-date'>"+item["date"].(string)+"</div>  </div>")
			}else{
				details = append(details, "<div id='name'><span>"+item["theid"].(primitive.A)[0].(primitive.M)["username"].(string)+"</span> <div id='notify-date'>"+item["date"].(string)+"</div> </div>")
			}

			if item["type"].(string) == "like" {
				details = append(details, "<div id='type-like' style='font-size:13px;'><i style='height:14px;width:14px;position:relative;top:2px;' data-feather='heart' color='red' fill='red'></i> "+item["body"].(string)+"</div>")
			}

			if item["type"].(string) == "commentlike"{
				details = append(details, `
				<div id='type-like' style='font-size:13px;'>
				    <div style="display:flex;padding-bottom:5px;">
				        <i style='height:14px;width:14px;position:relative;top:2px;' data-feather='heart' color='red' fill='red'></i>
					    <div style='padding-left:10px;'>Liked your comment</div>
				    </div>
				    
					<div style="background:#ccc;padding:5px 5px 5px 5px;border-radius:4px;">
				       `+item["body"].(string)+`
				    </div>
				</div>
				`)
			}

			if item["type"].(string) == "replylike"{
				details = append(details, `
				<div id='type-like' style='font-size:13px;'>
				    <div style="display:flex;padding-bottom:5px;">
				        <i style='height:14px;width:14px;position:relative;top:2px;' data-feather='heart' color='red' fill='red'></i>
					    <div style='padding-left:10px;'>Liked your reply</div>
				    </div>
				    
					<div style="background:#ccc;padding:5px 5px 5px 5px;border-radius:4px;">
				       `+item["body"].(string)+`
				    </div>
				</div>
				`)
			}

			if item["type"].(string) == "comment" {
				details = append(details, "<div id='type-comment'><div id='f-title' style='font-size:14px;'> <i style='height:14px;width:14px;position:relative;top:2px;' data-feather='message-circle' color='black' fill='black'></i> "+item["body"].(string)+"</div> <div id='notify-details' style='font-size:13px;'>"+item["comment"].(string)+"</div></div>")
			}

			if item["type"].(string) == "reply" {
				details = append(details, "<div id='type-comment'><div id='f-title'  style='font-size:14px;'> <i style='height:14px;width:14px;position:relative;top:2px;' data-feather='message-circle' color='black' fill='black'></i> "+item["body"].(string)+"</div> <div id='notify-details' style='font-size:13px;'>"+item["reply"].(string)+"</div></div>")
			}

			if item["type"].(string) == "tag" || item["type"].(string) == "comment-tag" || item["type"].(string) == "reply-tag"{
				details = append(details, "<div id='type-tag'><i data-feather='tag' color='black' fill='black'></i>  "+item["body"].(string)+"</div>")
			}


			if item["type"].(string) == "follow" {
				details = append(details, "<div id='type-tag' style='font-size:13px;'><i style='height:14px;width:14px;position:relative;top:2px;' data-feather='user' color='black' fill='black'></i>  "+item["body"].(string)+"</div>")
			}

			if item["type"].(string) == "share" {
				details = append(details, "<div id='type-tag' style='font-size:13px;'><i style='height:14px;width:14px;position:relative;top:2px;' data-feather='repeat' color='black' fill='black'></i>  "+item["body"].(string)+"</div>")
			}

			if item["type"].(string) == "poll" {
				details = append(details, 
				    `<div id='type-poll'>
				        <div id='f-title' style='font-size:14px;background:white;padding:5px 15px 5px 15px;border-radius:4px;margin-top:5px;'> 
						    <i style='height:14px;width:14px;position:relative;top:2px;' data-feather='bar-chart' color='black'></i> <b style="text-transform:capitalize;">`+item["option"].(string)+`</b> 
						</div>
						<div id='notify-details' style='font-size:13px;'>
						`+item["body"].(string)+`
						</div>
					</div>`)
			}

			if item["type"].(string) == "contributor" {
				details = append(details, 
					`
						<div class="contntfy090">
							<div style="margin-top:5px;" class="contntfy091">
							   <span>Invited you to join</span>
							</div>

							<div style="padding:5px 0px 5px 0px;" class="contntfy092">
								<span style="text-transform:capitalize;font-weight:600;">`+item["branchname"].(string)+`</span>
							</div>

							<div style="padding:0px 10px 0px 10px;display:flex;justify-content:space-between;" class="contntfy091">
							   <div class="contntfy091-btn0">
								   <button style="border-radius:4px;padding:5px 40px 4px 40px;background:green;border:green;color:white;">Accept</button>
							   </div>

							   <div class="contntfy091-btn0">
								   <button style="border-radius:4px;padding:5px 40px 4px 40px;background:red;border:red;color:white;">Reject</button>
							   </div>
							</div>

						</div>
					`)
			}

			if item["type"].(string) == "others branch" {
				details = append(details, 
				    `<div id='type-comment'>
				        <div id='f-title' style='font-size:14px;'> 
						    <i style='height:14px;width:14px;position:relative;top:2px;' data-feather='folder' color='black' fill='black'></i> Connected to <b style="text-transform:capitalize;">`+item["branchname"].(string)+`</b>
						</div>
						<div id='notify-details' style='font-size:13px;'>
						`+item["body"].(string)+`
						</div>
					</div>`)
			}

			details = append(details, "</div>")
			//name

			//top
			details = append(details, "</div>")

			details = append(details, "</div>")
			//notifications

			details = append(details, "</a>###")
			
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
			final = append(final, "<div class='notify-span' style='border-bottom:2px solid blue;' id='clckntf901'><div id='notify-svg' style='color:blue;pointer-events:none;'>All Notifications "+strconv.Itoa(len(countall))+"</div></div>")
		}else{
			final = append(final, "<div class='notify-span' style='border-bottom:2px solid blue;' id='clckntf901'><div id='notify-svg' style='color:blue;pointer-events:none;'>All Notifications </div></div>")
		}

		if len(countunread) > 0 {
			final = append(final, "<div class='notify-span' id='clckntf902'><div id='notify-svg' style='pointer-events:none;'>Unread "+strconv.Itoa(len(countunread))+"</div></div>")
		}else{
			final = append(final, "<div class='notify-span' id='clckntf902'><div id='notify-svg' style='pointer-events:none;'>Unread </div></div>")
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
			final = append(final, "<div class='notify-span' style='border-bottom:2px solid blue;' id='clckntf901'><div id='notify-svg' style='color:blue;pointer-events:none;'>All Notifications "+strconv.Itoa(len(countall))+"</div></div>")
		}else{
			final = append(final, "<div class='notify-span' style='border-bottom:2px solid blue;' id='clckntf901'><div id='notify-svg' style='color:blue;pointer-events:none;'>All Notifications</div></div>")
		}
		
		if len(countunread) > 0 {
			final = append(final, "<div class='notify-span' id='clckntf902'><div id='notify-svg' style='pointer-events:none;'>Unread "+strconv.Itoa(len(countunread))+"</div></div>")
		}else{
			final = append(final, "<div class='notify-span' id='clckntf902'><div id='notify-svg' style='pointer-events:none;'>Unread </div></div>")
		}
		final = append(final, "</div>")
		//btn
		
		final = append(final, "<div style='width:fit-content;margin:auto;' class='g239'><h2>Nothing to show here</h2></div>")

		final = append(final, "</div></div>")

	}

	return "<div style='position:fixed;top:0;bottom:0;left:0;right:0;background: rgba(0, 0, 0, 0.5);z-index:200;'>"+strings.Join(final, "")+"</div> <script>feather.replace()</script>"
}



func ImageShow(text string) string {
	var result string
	if len(text) > 0 {
		if text[0:1] == "#" {
			result  = "<img src='"+text[1:]+"' />"
		}else if text[0:4] == "auth" {
			result  = "<img src='"+text[4:]+"' />"
		}else{
			result  = "<img src='/dingtra-web-assets/images/profile/"+text+"' />"
		}
	}else{
		result = "<img src='/dingtra-web-assets/svg/pr.png' />"
	}

	return result
}

func CheckIfNotificationIsUnopened (item primitive.M) string {
	if item["opened"] == nil {

		return "&target=true"
	}

	return ""
}