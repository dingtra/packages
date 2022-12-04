$(document).ready(function(){
    $("#filer092").click(function(e){
        console.log("kkkk")
        document.getElementById(e.target.id).onchange = function(){

            var f = e.target.files;var o = [];var img = [];var vid = [];var frmt = [];
            if (f.length > 0 ){
    
                for(var item of f){
                    o.push(VPLYR(item.type))
                }

                for(var i=0; i < o.length; i++){
                    console.log(o[i])
                    if (/(jpeg|jpg|png|gif)$/i.test(o[i])){
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
                        $("#fl091282").hide();
                        $("#fl091281ch").hide();
                        $("#fl091280").show();
                        if (img.length > 0 || vid.length > 0){
                            $("#rmfile001").show();
                        }
                        tgglr()

                    }else if(img.length > 0 ){
                        for(var item in f){
                            if (isNaN(item) === false){
                                Cmpl(f[item], item)
                            }
                        }
                    }
                }
            }
        }
    })
})
