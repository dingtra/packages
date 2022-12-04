document.addEventListener("keyup", function(e){
    if(e.target.id === "frmpst27890editable"){
        var v=e.target.value

        if(v.trim().length > 0){
            $(".nwdngxr091-textr09x").show();

            if(v.length > 1000) {
                $("#nwdngxr091-text092x").html("<div class='counter-div-class'><span style='color:red;'>1000</span></div>")
                e.target.value = e.target.value.substr(0, 1000)
            }else{
                $("#nwdngxr091-text092x").html("<div class='counter-div-class'><span>"+(1000-v.length) +"/1000</span></div>")
            }
        }else{
            $(".nwdngxr091-textr09x").hide();
            
        }
    }
})

document.addEventListener("click", function(e){
    if(e.target.id != "frmpst27890editable"){
        $(".nwdngxr091-textr09x").hide();
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
        let v = e.target.value;let ch = e.target.checked;let c = Number(e.target.getAttribute("kp"));$("#brx0x9x1").val(v)
        for (let i = 0; i < c; i++) {
            if(ch === true){
                if(i == e.target.id.substr(e.target.id.length-1)){
                    continue;
                }
                $("#"+e.target.id.substr(0, e.target.id.length-1)+i).prop("checked", false);
            }else{
                $("#brx0x9x1").val("");
            }
        }
    }
})

document.addEventListener("click", function(e){
    if(e.target.id === "xl09br5409"){
        var x= e.target.getAttribute("x");
        $(".xro9polx89").html("<div style='width:fit-content;margin:auto;padding:10px 10px 10px 10px;'> <div class='loader'></div> </div>")
        $.post(x,
        {},
        function(d,s){
            if(s === "success"){
                var jsn = JSON.parse(d);
                for( const key in jsn){$("."+key).html(jsn[key])}
                
            }else{
                alert("something went wrong")
            }
       });
    }
})


function tgglr(){
    var v = $("#frmpst27890editable").val();var x = document.getElementById("filer092").files;
    console.log("F ",v.length, " sec ",x.length)
    if(v.trim().length > 0 || x.length > 0){$("#frmpst27890submtbtn").show();}else{$("#frmpst27890submtbtn").hide();}
}
