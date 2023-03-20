document.addEventListener("click", function(e){
    if(e.target.id == "frmpst27890submtbtn"){
        var j = document.forms.namedItem("dngfrm");
        var frm = new FormData(j);

        // check for rearranged files if any
        // compare with current files 
        // if rearranged files length != current files length discard the arrangements and proceed
        var checkfiles;

        if(document.querySelector("#filer092")){
            checkfiles = document.getElementById("filer092");
        }else if(document.querySelector("#filer092-bssnss")){
            checkfiles = document.getElementById("filer092-bssnss");
        }
        if(checkfiles){
            if(localStorage.getItem('files')){
                x = checkfiles.files;var checkimg;
    
                for(var item of x){
                    if (/(mp4|webm|ogg)$/i.test(item.type)){
                        checkimg = true
                    }
                }
                console.log(checkimg, " CHCHCHC")
                if(checkimg != true){
                    var itms = JSON.parse(localStorage.getItem('files'));
                    if(itms != null && Object.keys(itms).length === x.length ){
                        // files and rearranged files are thesame
                        console.log(JSON.stringify(localStorage.getItem('files')))
                        frm.append("arrangement", JSON.stringify(localStorage.getItem('files')));
                    }
                }
            }
        }
        
        var req = new XMLHttpRequest();
        $("#frmpst27890submt").html("<div style='width:fit-content;margin:auto;height:25px;'> <div style='margin-top:2px;height:20px;width:20px;' class='loader'></div> </div>")
         req.onreadystatechange = function() {

            if (this.readyState == 4 && this.status == 200) {
                var jsn = JSON.parse(this.responseText);
                console.log(jsn)
                for( const key in jsn){
                   if(key === "xk090xk09"){
                        // btn
                        $("."+key).html(jsn[key])
                    }else{
                        if(key != "erx090r"){ $("#"+key).html(jsn[key])}else{$("."+key).html(jsn[key])}
                    }
                }       
            }
        };
        req.open("POST", "/dngx/x09kl/");
        // req.open("POST", "/ddd/fff")
         
        req.send(frm);
         
    }
})
