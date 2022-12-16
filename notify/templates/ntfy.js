
document.addEventListener("click", function(e){
    if(e.target.id === "ntfyxr0k87") {
        $("#cnclnotifyx780").html("");$("#cnclnotifyx781").html("")
        $("#notify-dsply09x").html("<div style='position:fixed;top:0;bottom:0;left:0;right:0;background: rgba(0, 0, 0, 0.5);z-index:100;'><div style='width:300px;margin:auto;padding:20px 10px 20px 10px;background:white;border-radius:4px;margin-top:50px;'><img style='position:relative;top:2px;height:15px;width:15px;' src='/dingtra-web-assets/svg/l.jpg'/> Processing request</div></div>")
        $.post("/notify/in/",
        {},
        function(dt,s){
            if(s === "success"){
                $("#notify-dsply09x").html(dt)
            }else{
                alert("something went wrong")
            }
       });

    }else if(e.target.id === "ntfyxr0k87cncl"){
        $("#notify-dsply09x").html("")
    }
})


document.addEventListener("click", function(e){
    if(e.target.id === "clckntf902" || e.target.id === "clckntf901") {
        $("#notify-dsply09x").html("<div style='position:fixed;top:0;bottom:0;left:0;right:0;background: rgba(0, 0, 0, 0.5);z-index:100;'><div style='width:300px;margin:auto;padding:20px 10px 20px 10px;background:white;border-radius:4px;margin-top:50px;'><img style='position:relative;top:2px;height:15px;width:15px;' src='/dingtra-web-assets/svg/l.jpg'/> Processing request</div></div>")
        $.post(e.target.id === "clckntf902" ? "/notify/in/?default=unread" : "/notify/in/",
        {},
        function(dt,s){
            if(s === "success"){
                $("#notify-dsply09x").html(dt)
            }else{
                alert("something went wrong")
            }
       });
    }else if(e.target.id === "ntfyxr0k87cncl"){
        $("#notify-dsply09x").html("")
    }
})



document.addEventListener("click", function(e){
    if(e.target.id === "no98090") {
        $("#notify-dsply09x").html("<div style='position:fixed;top:0;bottom:0;left:0;right:0;background: rgba(0, 0, 0, 0.5);z-index:100;'><div style='width:300px;margin:auto;padding:20px 10px 20px 10px;background:white;border-radius:4px;margin-top:50px;'><img style='position:relative;top:2px;height:15px;width:15px;' src='/dingtra-web-assets/svg/l.jpg'/> Processing request</div></div>")
        $.post("/notify/in/xr1/",
        {},
        function(dt,s){
            if(s === "success"){
                $("#notify-dsply09x").html(dt)
            }else{
                alert("something went wrong")
            }
       });
    }else if(e.target.id === "ntfyxr0k87cncl"){
        $("#notify-dsply09x").html("")
    }
})

