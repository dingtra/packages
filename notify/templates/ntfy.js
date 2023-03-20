
document.addEventListener("click", function(e){
    if(e.target.id === "ntfyxr0k87") {
        $("#cnclnotifyx780").html("");$("#cnclnotifyx781").html("")
        $("#notify-dsply09x").html("<div style='position:fixed;top:0;bottom:0;left:0;right:0;background: rgba(0, 0, 0, 0.5);z-index:100;'><div style='width:300px;margin:auto;padding:20px 10px 20px 10px;background:white;border-radius:4px;margin-top:50px;'>  <div style='width:fit-content;margin:auto;padding:10px 10px 10px 10px;'><div class='loader'></div></div> <div style='width:fit-content;margin:auto;'>Processing request</div> </div>")
        $.post("/notify/xr0x7/",
        {},
        function(dt,s){
            if(s === "success"){
                $("#notify-dsply09x").html(dt);var xk = $(".view-all-notify").html();var height = window.innerHeight;
                if(xk.length > 0 ){ if(height > 500){$(".view-all-notify").css("height", height-200+"px")}}
                $(".xklo090").html("");$(".xklo091").html("");
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
        $("#notify-dsply09x").html("<div style='position:fixed;top:0;bottom:0;left:0;right:0;background: rgba(0, 0, 0, 0.5);z-index:100;'><div style='width:300px;margin:auto;padding:20px 10px 20px 10px;background:white;border-radius:4px;margin-top:50px;'>  <div style='width:fit-content;margin:auto;padding:10px 10px 10px 10px;'><div class='loader'></div></div> <div style='width:fit-content;margin:auto;'>Processing request</div> </div>")
        $.post(e.target.id === "clckntf902" ? "/notify/xr0x7/?default=unread" : "/notify/xr0x7/",
        {},
        function(dt,s){
            if(s === "success"){
                $("#notify-dsply09x").html(dt);var xk = $(".view-all-notify").html();var height = window.innerHeight;
                if(xk.length > 0 ){ if(height > 500){$(".view-all-notify").css("height", height-200+"px")}}
                
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
        $("#notify-dsply09x").html("<div style='position:fixed;top:0;bottom:0;left:0;right:0;background: rgba(0, 0, 0, 0.5);z-index:100;'><div style='width:300px;margin:auto;padding:20px 10px 20px 10px;background:white;border-radius:4px;margin-top:50px;'>  <div style='width:fit-content;margin:auto;padding:10px 10px 10px 10px;'><div class='loader'></div></div> <div style='width:fit-content;margin:auto;'>Processing request</div> </div>")
        $.post("/notify/xr0x7/xr1/",
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

