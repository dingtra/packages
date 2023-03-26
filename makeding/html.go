package makeding


import (
	"rundb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
	"strconv"
	// "html/template"
	// "encoding/json"
	// "fmt"
	
)

// css 096

func (let *MakeDingStruct ) NewDing () string{
	details := []string{}

	// details = append(details, `
	
	// <html>

	// <head>
	// 	<title>Home Page</title>
	// 	<link rel="icon" href="/jskl899nxhsjas/" type="image/x-icon">
	// 	<meta charset="UTF-8">
	// 	<meta name="viewport" content="width=device-width, initial-scale=1.0" />
	// 	<link rel="stylesheet" href="/assets/css/mkdng0xr0.css" />
	// 	<link rel="stylesheet" href="/assets/css/mkdngbssnss090.css" />
	// 	<link rel="stylesheet" href="/assets/css/ldx.css" />
	// 	<script src="/assets/js/feather/feather.js"></script>
	// 	<script src="/assets/js/jquery.js"></script>
	// 	<script src="/assets/js/mkdngs090/mkding090.js"></script>
	// 	<script src="/assets/js/mkdngs090/mkding0901.js"></script>
	// 	<script src="/assets/js/mkdngs090/mkding092.js"></script>
	// 	<script src="/assets/js/mkdngs090/mkding093.js"></script>
	// 	<script src="/assets/js/mkdngs090/mkding094.js"></script>
	// 	<script src="/assets/js/mkdngs090/mkding095.js"></script>
	// 	<script src="/assets/js/mkdngs090/mkding096.js"></script>
	// 	<script src="https://cdn.jsdelivr.net/npm/emoji-mart@latest/dist/browser.js"></script>
	// </head>
	
	// <body >
	// `)

	// parent still open
	details = append(details, 
	   `<div style="position:fixed;top:0;bottom:0;right:0;left:0;z-index:300;background:rgba(0, 0, 0, 0.5);" class="newdngsxr090"><div style="background:white;padding-bottom:50px;" class="newdngsxr090-chld">
	        <div class='erx090r'></div>
	`)


	details = append(details, `
	    <div class="newdngsxr090ax">
		    <div class="newdngsxr090ax0">
			    <span>Create a Ding</span> 
			</div>

			<div class="newdngsxr090ax1">
				<div id='ntfyxr0k87cncl'>
					<div style='pointer-events:none;'>
						<i data-feather='x' color='black'></i>
					</div>
				</div>	
			</div>

		</div>
	`)


	// start form
	details = append(details, 
		`<form enctype="multipart/form-data" name="dngfrm">
	`)

	details = append(details, `
		<div class="nwdngtext090">
		    <div style="position:relative;" class="nwdngtext091x">
				<div style="position:relative;" class="nwdngtext092x">
				    <div style="position:absolute;border:1px solid #ccc;border-bottom:white;border-radius:4px;border-bottom-left-radius: unset;border-bottom-right-radius: unset;top:0px;padding:5px 10px 5px 10px;border-bottom:1px solid #e8e4e4;left:0;right:0;" class="icons-xvg">
					    <div style="display:flex;">
							<div id="iconsx00" data="emojixl090">
								<span style=" pointer-events: none;">
									<i data-feather='smile'></i>
								</span>
							</div>
							<div style="display:flex;margin-left:10px;padding:3px 3px 3px 3px;background:#f2f2f2;border-radius:5px;font-size:13px;" id="iconsx01" data="emojixl090">
								<span style=" pointer-events: none;">
									AI Assistant
								</span>
							</div>

							<div style="margin:5px 0px 0px 10px;font-size:15px;color:green;" class="phsxr092-desc">
					            <div id="phsxr094-desc"></div>
					        </div>
						</div>

					</div>
					<textarea id='frmpst27890editable' name='about' data='frmpst27890submtbtn' placeholder='What&#39;s the ding all about?'></textarea>
				</div>
				<div style="display:none;" id="emojixl090">
					<div style="position:fixed;bottom:0;z-index:400;" class="nwdngsxkl00">
						<div id="emoji-picker-container"></div>
					</div>
				</div>
					
				<div style="background:#f2f2f2;z-index:500;box-shadow: 0 3px 10px rgb(0 0 0 / 0.1);border:0px solid #ccc;" class="nwdngxr092"></div>
			</div>

			<div class="nwdngtext091k">
			    <div class="nwdngtext091k-chld">
				    <div class="nwdngtext092k">
					    <div style="padding:5px 0px 5px 0px;"><span>Search by username to connect</span></div>
						<div style="position:relative;">
						    <span> <input type="text" placeholder="Connect a branch" id="lkbrnchxr09x" /></span>
							<i style="position:absolute;right:10px;top:10px;height:16px;width:16px;color:#ccc;" data-feather='search' color='black'></i>
						</div>
					</div>
					<div style='display:none;'>
		                <span><input type='text' name='inbranch' id='brx0x9x1' /></span>
		            </div>
					<div class="nwdngtext093k">
						<div id="lkupxr09x"></div>
					</div>
					<div class="nwdngtext094k">
					    `+ConnectBranch(let.Usersid)+`
					</div>
				</div>
			</div>
		</div>
	`)

	// open options

	details = append(details, `
	    <div class='create-new-ding-class-files flxkr09clo0'>
	`)

	// 88888888888888888888888888888888888888888888888888888888888888888888888888888888888
	// this is the first files => for serving raw videos and photos without more context
	// the aim is to switch the files and business filess which allow for bussiness to express more in a great way
	// this id will be => flchngxr090-ti
	details = append(details, "<div class='flchngxr090' style='display:none;' ><span><input type='file' id='filer092' name='files' accept='video/*,image/*' multiple/></span></div>")
	// 8888888888888888888888888888888888888888888888888888888888888888888888888888888888888
	
	details = append(details, `
		    `+BranchOptions("files")+`
		    <div class="flxjss090">
				<div class="flchngxr091"> 
				    <label for='filer092'>
						<div id='mkdnglabl090' class='mkdnglabelx090'>
							<i data-feather='cloud'></i> <span> Add files here</span>
						</div>
						<div id='mkdnglabl091' class='mkdnglabelx091'>
							<span>Atleast 15 MB for videos & 5mb for photos, we are currently on test stage</span>
						</div>
				    </label>
				</div>

				<div style='display:none;' id='fl091281'><div style='display:flex;overflow: auto;white-space: nowrap;' id='fl091281ch'></div></div>
				<div style='display:none;' id='fl091280'><video id='vdfl091280' controls >not supported</video></div>
			</div>
			<div class="setxkjss090"></div>
		</div>
	`)
	// close options

	// submit
	details = append(details, "<div class='xk090xk09'><div style='display:none;padding:10px 0px 10px 0px;margin-top:10px;' id='frmpst27890submtbtn'><div style='pointer-events:none;' id='frmpst27890submt'><span>New ding</span></div></div></div>")
	// submit

	details = append(details, `</form>`)
	// end form

	details = append(details, `</div></div>`)
	// parent

	details = append(details, `
	<script>

      var picker = new EmojiMart.Picker({
        container: document.querySelector('#emoji-picker-container'),
        emojiSize: 20,
        showPreview: false,
        showSkinTones: false,
        perLine: 8,
		onClickOutside:function(e){
			if(e.target.id != "iconsx00"){
				document.getElementById("emojixl090").style.display="none";
			}
		},
        onEmojiSelect:function(event){
			var html = document.getElementById("frmpst27890editable");html.value = html.value+event.native;
		},
        defaultCategory: 'smileys',
      });

       document.getElementById("emoji-picker-container").appendChild(picker)

    </script>
	<script>
	feather.replace()
    </script>
	`)

	// details = append(details, `
	// </body>
    // </html>
	// `)

	return strings.Join(details, "")
}



func ConnectBranch (usersid string ) string {
	getcoll := rundb.Connect("app").Collection("branch")

	finder := rundb.FindAll(getcoll, bson.M{})

	details := []string{}

	mapper := make(map[string]string)

	for _, item := range finder {

		if usersid == item["usersid"].(primitive.ObjectID).Hex() {

			details = append(details, item["name"].(string))
			mapper[item["name"].(string)] = item["_id"].(primitive.ObjectID).Hex()
		}
	}


	final := []string{}

	for _, k := range details {
		final = append(final, "<div class='brnch82791-slct'>")
		final = append(final, "<div id='brnch82791-slct01'><span>"+RemoveSlash(k)+"</span></div>")
		final = append(final, "<div id='brnch82791-slct02'><input type='checkbox' id='brnc12' name='"+strconv.Itoa(len(details))+"' value='"+mapper[k]+"' x='"+RemoveSlash(k)+"'></div>")
		final = append(final, "</div>")
	}

	if len(final) > 0 {
		return strings.Join([]string{
			"<div class='brnch-form-class'><div class='brnch-form-class-ch'>",
			"<div id='brnch82791-title'><span>Connect to an existing branch</span>  </div>",
			"<div style='margin-top:10px;' id='brtogglex90r'>",
			"<div id='brnchjax090res'>"+strings.Join(final, "")+"</div>",
			"<div id='brnch82791-about'><span>Connect to a single branch at a time <a href='#'>Learn more about branch</a></span></div>",
			"</div>",
			"</div></div>",
		}, "")

	}else{
		return strings.Join([]string{
			"<div class='brnch-form-class'><div class='brnch-form-class-ch'>",
			"<div id='brnch82791-title'><span>Connect to an existing branch</span> </div>",
			// 
			"<div style='margin-top:10px;' id='brtogglex90r'>",
			"<div id='brnch82791-about'><span>Cant connect to an empty branch.</span></div>",
			// create new branch
			`<div id='brnch82791-about' style='margin-top:5px;'>
			    <div style="background:green;color:white;padding:5px 10px 5px 10px;border-radius:4px;" class='asp-new-brnch'>
				    <div style="width:fit-content;margin:auto;"><i style="height:16px;width:16px;position:relative;top:2px;color:white" data-feather='folder-plus' color='black'></i> <span>Create new branch</span></div>
			    </div>
			</div>`,
			// create new branch
			"</div>",
			// 
			"</div></div>",
		}, "")
	}

	return ""
}


func RemoveSlash (val string) string {

	return strings.Replace(val, "-", " ", 2)
}
