
document.addEventListener("click", function(e){
    if(e.target.id === "filer092" || e.target.id === "filer092-bssnss"){
        document.getElementById(e.target.id).onchange = function(){

            var f = e.target.files;var o = [];var img = [];var vid = [];var frmt = [];
            if (f.length > 0 ){
    
                for(var item of f){
                    o.push(VPLYR(item.type))
                }

                for(var i=0; i < o.length; i++){
                    console.log(o[i])
                    if (/(jpeg|jpg|png|gif|webp)$/i.test(o[i])){
                        img.push(o[i])
                    }else if (/(mp4|webm|ogg)$/i.test(o[i])){
                        vid.push(o[i])

                    }else{
                        frmt.push(o[i])
                    }
                }

                if(frmt.length > 0){
                    alert("format not supported")
                }else{
                    if(img.length > 0 && vid.length > 0){
                        e.target.value=""
                        $("#err9280").html("<div style='padding:5px 5px 5px 5px;font-size:13px;'><span style='color:red;'>select image or video per ding</span></div>")
                        setTimeout(function() { 
                            $("#err9280").html("")
                        }, 4000);
    
                    }else if(vid.length > 0 && vid.length > 1){
                        $("#err9280").html("<div style='padding:5px 5px 5px 5px;font-size:13px;'><span style='color:red;'>Select one video at a time</span></div>")
                        setTimeout(function() { 
                            $("#err9280").html("")
                        }, 4000);
                    }else if(vid.length > 0){
                        let fle =e.target.files[0]
                        let blb = URL.createObjectURL(fle);
                        document.querySelector("video").src = blb;
                        $("#mkdnglabl090").hide();$("#mkdnglabl091").hide();
                        $("#fl091281").hide();
                        $("#fl091280").show();
                        // if (img.length > 0 || vid.length > 0){
                        //     $("#rmfile001").show();
                        // }
                        tgglr()

                       

                    }else if(img.length > 0 ){
                        // clear previous files if any
                        // arrangement files
                        localStorage.removeItem("files");
                        for(var item in f){
                            if (isNaN(item) === false){
                                if(e.target.id === "filer092-bssnss"){
                                    Cmpl(f[item], item)
                                }else{
                                    if(e.target.id === "filer092"){
                                        NormCmpl(f[item], item, f)
                                        // alert(f.length)
                                    }
                                }
                            }
                        }
                    }
                }
            }
        }
    }
})
