document.addEventListener("keyup", function(e){
    if(e.target.id === "frmpst27890editable"){
        var v=e.target.value

        if(v.trim().length > 0){
            $(".phsxr092-desc").show();

            if(v.length > 1000) {
                $("#phsxr094-desc").html("<div class='counter-div-class'><span style='color:red;'>1000</span></div>")
                e.target.value = e.target.value.substr(0, 1000)
            }else{
                $("#phsxr094-desc").html("<div class='counter-div-class'><span>"+(1000-v.length) +"/1000</span></div>")
            }
        }else{
            $(".phsxr092-desc").hide();
            
        }
    }
})

document.addEventListener("keyup", function(e){
    if(e.target.id === "lkbrnchxr09x"){
        $("#lkupxr09x").html("<div style='width:fit-content;margin:auto;height:25px;'> <div style='margin-top:2px;height:10px;width:10px;' class='loader'></div> </div>")
        if(e.target.value.length > 0){
            $.post("/xoz/qr1/tg1x",
            {
                name:e.target.value,
         
            },
            function(d,s){
                if(s === "success"){
                    var jsn = JSON.parse(d);
                    for( const key in jsn){$("#"+key).html(jsn[key])}
                }else{
                    alert("something went wrong")
                }
           });
        }else{
            $("#lkupxr09x").html("");
        }
    }
})


document.addEventListener("click", function(e){
    // look up name - btn clicks
    if (e.target.id === "tglkupxr0x"){
        var x = e.target.getAttribute("mx");
        $("#refxr09").html("<div style='width:fit-content;margin:auto;height:15px;'> <div style='margin-top:2px;height:10px;width:10px;' class='loader'></div> </div>");

        $.post("/xoz/qr1/lkbrx1/",
        {
            data:x,
        },
        function(d,s){
            if(s === "success"){
                var jsn = JSON.parse(d);
                // console.log(d)
                for( const key in jsn){$("."+key).html(jsn[key])}
                
            }else{
                alert("something went wrong")
            }
       });
    }
})

document.addEventListener("click", function(e){
    if(e.target.id != "frmpst27890editable"){
        $(".nwdngxr091-textr09x").hide();
    }
})

document.addEventListener("click", function(e){
    if(e.target.id === "iconsx00"){
        document.getElementById(e.target.getAttribute("data")).style.display = "block";
    }

})


document.addEventListener("click", function(e){
    if(e.target.id === "brnc12"){
        $("#brnchjax090res").html("<div style='width:fit-content;margin:auto;padding:10px 10px 10px 10px;'> <div class='loader'></div> </div>")
        $.post("/xoz/qr1/xk1/",
        {
            data:e.target.value,
            x:e.target.getAttribute("x"),
        },
        function(d,s){
            if(s === "success"){
                var jsn = JSON.parse(d);
                for( const key in jsn){$("#"+key).html(jsn[key])}
                
            }else{
                alert("something went wrong")
            }
       });
    }else if(e.target.id.substr(0,12) === "inbranch-res"){
        let el = e.target.id.substr(12);let v = e.target.value;let ch = e.target.checked;let sum = Number(e.target.getAttribute("kp"));$("#brx0x9x1").val(v);
        for (let i = 1; i < sum+1; i++) {
            if (ch === true){
                if(i == el) {
                    continue
                }
                $("#"+e.target.id.substr(0, 12)+i).prop("checked", false);
            }
        }

    }else if(e.target.id.substr(0,12) === "refxr09"){
        $("#refxr09").html("<div style='width:fit-content;margin:auto;height:15px;'> <div style='margin-top:2px;height:10px;width:10px;' class='loader'></div> </div>"); var url = e.target.getAttribute("url");
        $.post("/xoz/qr1/"+url,
        {},
        function(d,s){
            if(s === "success"){
                var jsn = JSON.parse(d);
                for( const key in jsn){$("#"+key).html(jsn[key])}
                $("#brx0x9x1").val("")
            }else{
                alert("something went wrong")
            }
       });
    }
})

document.addEventListener("click", function(e){
    if(e.target.id === "xl09br5409"){
        var x= e.target.getAttribute("x");
        $.post("/xoz/qr1/optnsxr090",
        {data:x,},
        function(d,s){
            if(s === "success"){
                var jsn = JSON.parse(d);
                for( const key in jsn){$("."+key).html(jsn[key])}
                tgglr();
            }else{
                alert("something went wrong")
            }
       });
        
    }
})


function tgglr(){
    var v = $("#frmpst27890editable").val();var x;

    if(document.querySelector("#filer092")){
        x = document.getElementById("filer092").files;
    }else if(document.querySelector("#filer092-bssnss")){
        x = document.getElementById("filer092-bssnss").files;
    }
    if(v.trim().length > 0 || x.length > 0){$("#frmpst27890submtbtn").show();}else{$("#frmpst27890submtbtn").hide();}
}

