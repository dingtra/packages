function Ptag(i, n){
	return [
        "<div class='sellres901' id='btnsell' nx='"+i+"'><div class='img9780'>",
        "<div style='padding:5px 0px 5px 0px;font-size:13px;'><span>Add Title</span></div><div id='img9780input'><input style='height:30px;border-radius:4px;border:1px solid #ccc;padding-left:10px;width:100%;' type='text' name='"+n+"ungogo' placeholder='Product Name'/></div>",
        "<div style='padding:5px 0px 5px 0px;font-size:13px;'><span>Add Price</span></div><div  id='img9780input'><div style='display:flex;'><select name='"+n+"currency'><option>NGN</option><option>USD</option></select><input style='height:30px;border-radius:4px;border:1px solid #ccc;padding-left:10px;width:100%;' type='number' placeholder='Price' name='"+n+"'/></div></div>",
        "</div></div>"
	].join("");
}

function VPLYR (s){
    var l = s.split("/");var ls = l.pop();
    return ls;
}

function Cmpl(fle, indx) {
    var rdr = new FileReader();
    rdr.addEventListener("load", function () {
        $("#vdfl091280").trigger('pause');$("#vdfl091280").val("");$("#fl091280").hide();$("#fl091282").hide();
        document.getElementById("fl091281ch").innerHTML += "<div class='imgbg09' style='position:relative;'><img src='"+this.result+"' id='img091' nx='"+indx+"' />  <div id='mrkt"+indx+"' style='display:none;position:absolute;top:0;bottom:0;left:0;right:0;background:white;'><div style='height:fit-content;width:100%;padding:10px 5px 10px 5px;background:#f2f2f2;border-radius:4px;'>"+Ptag(indx, fle.name)+"</div> </div> </div>";
        $("#fl091281").show(); $("#fl091281ch").show();
        $("#rmfile001").show();
        tgglr()
    }, false);
    rdr.readAsDataURL(fle);
}



$(document).ready(function(){
    $(document).click(function(e){
        if(e.target.id === "img091"){
           var t = e.target.getAttribute("nx");$("#mrkt"+t).show();
        }else if(e.target.id.substr(0,4) === "mrkt"){
            $("#"+e.target.id).hide();
        }
    })
})


$(document).ready(function(){
    $(document).click(function(e){
        if(e.target.id === "rmfile001"){
            $("#filer092").val("");$("#fl091280").hide();$("#fl091281").hide();$("#fl091281ch").html("");$("#rmfile001").hide();$("#fl091282").show();tgglr();
        }
    })
})