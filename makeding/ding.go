package makeding


import (
	"rundb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
	"strconv"
	
)



func (let *MakeDingStruct ) NewDing () string {
	details := []string{}

	// details = append(details, `
	
	// <html>

	// <head>
	// 	<title>Home Page</title>
	// 	<link rel="icon" href="/jskl899nxhsjas/" type="image/x-icon">
	// 	<meta charset="UTF-8">
	// 	<meta name="viewport" content="width=device-width, initial-scale=1.0" />
	// 	<link rel="stylesheet" href="/assets/css/mkdng0xr0.css" />
	// 	<link rel="stylesheet" href="/assets/css/ldx.css" />
	// 	<script src="/assets/js/feather/feather.js"></script>
	// 	<script src="/assets/js/jquery.js"></script>
	// 	<script src="/assets/js/mkdngs090/mkdngs090xr.js"></script>
	// 	<script src="/assets/js/mkdngs090/mkdngs0901xr.js"></script>
	// 	<script src="/assets/js/mkdngs090/mkdngs0902xr.js"></script>
	// 	<script src="/assets/js/mkdngs090/mkdngs0903xr.js"></script>
	// </head>
	
	// <body >
	// `)
	
	// parent 
	details = append(details, "<div class='nwdngxr090'><div class='nwdngxr090-chld'>")
	details = append(details, "<div class='erx090r'></div>")
	details = append(details, "<div id='ntfyxr0k87cncl'><div style='pointer-events:none;display:flex;justify-content:end;padding:5px 10px 10px 0px;'><button><i data-feather='x' color='black'></i></button> </div></div>")
	details = append(details, "<form enctype='multipart/form-data' name='dngfrm'>")
	// phase
	details = append(details, "<div class='phsxr09x1'>")

	// phase one
	details = append(details, "<div class='nwdngxr090-phase0x90'>")

	details = append(details, "<div class='nwdngxr090-dscrptn'><span>Start a new ding</span></div>")

	// description
	details = append(details, "<div class='nwdngxr091-textr09xyz'>")
	details = append(details, "<div class='nwdngxr091'><span> <textarea id='frmpst27890editable' name='about' placeholder='What&#39;s the ding all about?'></textarea> </span></div>")

	// 
	details = append(details, "<div style='display:none;' class='nwdngxr091-textr09x'><div class='nwdngxr091-textr09x-chld'>  <div id='nwdngxr091-text092x'>0/1000</div> </div></div>")

	// 

	details = append(details, "</div>")
	// description


	// srch details
	details = append(details, "<div class='nwdngxr092'>")
	details = append(details, "</div>")
	// srch details

	// 00000000000000000000000000000000000000000000000000000000000000000000000000000000
	details = append(details, "<div style='display:none;'><span><input type='text' name='branchchild' id='brx0x9x1' /></span></div>")
	// 00000000000000000000000000000000000000000000000000000000000000000000000000000


	// options
	details = append(details, "<div class='nwdngxr093-options'><div class='nwdngxr093-options-ch'>")

	// 
	details = append(details, "<div id='xl09br5409' x='/xoz/qr1/sk0x'><div style='pointer-events:none;' class='nwdngxr093-options-phase  nwdngxr093-xr'>")
	details = append(details, "<div id='nwdngxr093-options-phase-svg'><span><i data-feather='file-plus' color='black'></i></span></div>")
	details = append(details, "<div id='nwdngxr093-options-phase-title'><span>Add files</span></div>")
	details = append(details, "</div></div>")
	//

	// 
	details = append(details, "<div id='xl09br5409' x='/xoz/qr1/sk1x'><div style='pointer-events:none;border-left:1px solid #ccc;' class='nwdngxr093-options-phase  nwdngxr093-pushox'>")
	details = append(details, "<div id='nwdngxr093-options-phase-svg'><span><i data-feather='bar-chart' color='black'></i></span></div>")
	details = append(details, "<div id='nwdngxr093-options-phase-title'><span>New poll</span></div>")
	details = append(details, "</div></div>")
	//

	// 
	details = append(details, "<div id='xl09br5409' x='/xoz/qr1/'><div style='pointer-events:none;border-left:1px solid #ccc;' class='nwdngxr093-options-phase  nwdngxr093-pushox'>")
	details = append(details, "<div id='nwdngxr093-options-phase-svg'><span><i data-feather='trending-up' color='black'></i></span></div>")
	details = append(details, "<div id='nwdngxr093-options-phase-title'><span>Business</span></div>")
	details = append(details, "</div></div>")
	//

	// 
	details = append(details, "<div id='xl09br5409' x='/xoz/qr1/'><div style='pointer-events:none;border-left:1px solid #ccc;' class='nwdngxr093-options-phase  nwdngxr093-pushox'>")
	details = append(details, "<div id='nwdngxr093-options-phase-svg'><span><i data-feather='list' color='black'></i></span></div>")
	details = append(details, "<div id='nwdngxr093-options-phase-title'><span>Waitlist</span></div>")
	details = append(details, "</div></div>")
	//

	details = append(details, "</div></div>")
	// options

	details = append(details, ConnectBranch("632230919cc0a078594616fc"))

	details = append(details, "</div>")
	// phase one

	// phase two
	details = append(details, "<div class='nwdngxr094-phase0x94'>")
	
	// 88888888888888888888888888888888888888888888888888888888888888888888888888888888888
	details = append(details, "<div style='display:none;'><span><input type='file' id='filer092' name='files' accept='video/*,image/*' multiple/></span></div>")
	// 8888888888888888888888888888888888888888888888888888888888888888888888888888888888888
	

	// files
	details = append(details, "<div class='create-new-ding-class-files'><div class='create-new-ding-class-files-ch xro9polx89'>")

	details = append(details, "<div id='fl091282'><label for='filer092'><div id='create-new-ding-class-files01'><i data-feather='image'></i> <span> Add files here</span></div></label></div>")
	details = append(details, "<div style='display:none;' id='fl091281'><div style='display:flex;overflow: auto;white-space: nowrap;' id='fl091281ch'></div></div>")
	// video
	details = append(details, "<div style='display:none;' id='fl091280'><video id='vdfl091280' controls >not supported</video></div>")

	// remove
	details = append(details, "<div style='display:none;' id='rmfile001'><div style='pointer-events:none;' class='rmfile'><div id='rmfile'><span>Remove file</span></div></div></div>")

	details = append(details, "</div></div>")
	// files

	details = append(details, "</div></div>")
	
	// submit
	details = append(details, "<div style='display:none;' id='frmpst27890submtbtn'><div class='frmpst27890submtbtn-block'><div><div style='pointer-events:none;' class='frmpst27890submt'><div id='frmpst27890submt'><span>New ding</span></div></div> </div></div></div>")


	details = append(details, "</div>")
	// phase two

	details = append(details, "</div>")
	// phase

	details = append(details, "</form>")

	details = append(details, "</div></div>")
	// parnt


	details = append(details, `
	<script>
	  feather.replace()
    </script>
    </body>
    </html>
	`)

	return strings.Join(details, "")
}


func ConnectBranch (usersid string ) string {
	getcoll := rundb.Connect("app").Collection("branch")

	finder := rundb.FindAll(getcoll)

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
			"<div id='brnch82791-title'><span>Connect to an existing branch</span> <i data-feather='chevron-down' color='black'></i></div>",
			"<div style='display:block;' id='brtogglex90r'>",
			"<div id='brnchjax090res'>"+strings.Join(final, "")+"</div>",
			"<div id='brnch82791-about'><span>you can connect to a single branch at a time <a href='#'>Learn more about branch</a></span></div>",
			"</div>",
			"</div></div>",
		}, "")

	}else{
		return strings.Join([]string{
			"<div class='brnch-form-class'><div class='brnch-form-class-ch'>",
			"<div id='brnch82791-title'><span>Connect to an existing branch</span> <i data-feather='chevron-down' color='black'></i></div>",
			"<div style='display:block;' id='brtogglex90r'>",
			"<div id='brnch82791-about'><span>Please create a branch and try again.</span></div>",
			"<div id='brnch82791-about' style='margin-top:5px;'><span style='font-size:13px;'>cant connect to an empty branch</span></div>",
			"</div>",
			"</div></div>",
		}, "")
	}

	return ""
}


func RemoveSlash (val string) string {

	return strings.Replace(val, "-", " ", 2)
}
