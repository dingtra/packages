package dingtag

import (
	// 
)

func ImageShow(result string) string {
	var details string;
	if len(result) > 0 {
		if result[0:1] == "@" {
			details  = "<img src='"+result[1:]+"' />"
		}else{
			details = "<img src='/dingtra-web-assets/images/profile/"+result+"' />"
		}
	}else{
		details = "<img src='/dingtra-web-assets/svg/pr.png' />"
	}

	return details
	
}