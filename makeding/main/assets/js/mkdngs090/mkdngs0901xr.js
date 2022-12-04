document.addEventListener("keyup", function(e){
    if(e.target.id === "frmpst27890editable"){
        var v = e.target.value;var k = e.target.selectionStart;var s = [];

        if(v.trim().length > 0 ){
            $("#frmpst27890submtbtn").show();
            if(v[k-1] === "@"){
                s.push(k-1)
            }
            
            var kl = v.substr(s[0], k);var sp = kl.split(" ");var res = sp[sp.length-1]
            
            if(res.length > 0) {
                if(res[0] === "@"){
                    var xp = res.substr(1, res.length);
    
                    if (xp.length > 0 ){
                        $(".nwdngxr092").html("<div style='width:fit-content;margin:auto;padding:10px 10px 10px 10px;'> <div class='loader'></div> </div>")
                        $.post("/xoz/qr1/tg0x",
                        {
                            name:xp,
                     
                        },
                        function(d,s){
                            if(s === "success"){
                                var jsn = JSON.parse(d);
                                for( const key in jsn){$("."+key).html(jsn[key])}
                            }else{
                                alert("something went wrong")
                            }
                       });
                    }
                }else{
                    $(".nwdngxr092").html("");
                }
    
            }else{
                $(".nwdngxr092").html("");
            }
        }else{
            $("#frmpst27890submtbtn").hide();
        }
        
    
    }
})
      

document.addEventListener("click", function(e){
    // alert(e.target.id)
    if(e.target.id === "tgall091") {
        var tg = e.target.getAttribute("gx")
        var v = document.getElementById("frmpst27890editable").selectionStart;
        var b = $("#frmpst27890editable").val();
        
        var t = [];
        
        for(var i=b.substr(0, v).length - 1 ; i >= 0; i--){
            
            if(b[i] === "@"){
                t.push(i)

            }
        }
        
        if(t.length > 0 ){
            var tag = "@"+tg
            var x1 = b.substr(0, v);
            var p1 = x1.substr(0, t[0]);
            var p2 = b.substr(v);
            $("#frmpst27890editable").val(p1+tag+p2)
            StCrt("frmpst27890editable", t[0]+tag.length)
            $(".nwdngxr092").html("")
        }
    }
})


function StCrt(id, pos) {
    var e = document.getElementById(id);
    if (e != null) {
        if(e.createTextRange){

            var r = e.createTextRange();
            r.move("character", pos);
            r.select();
        }else{
            if(e.selectionStart) {
                e.focus();
                e.setSelectionRange(pos, pos);
            }else{
                e.focus();
            }
        }
    }
}

// Add an event listener for the "focus" event
document.addEventListener('click', (e) => {
    if(e.target.id === "tgall091-"){
        var textarea = document.getElementById("frmpst27890editable");var v = textarea.value; var k = textarea.selectionStart;var s = [];
        
        if(v[k-1] === "@"){
            s.push(k-1)
        }
        
        var kl = v.substr(s[0], k);var sp = kl.split(" ");var res = sp[sp.length-1]
        
        if(res.length > 0) {
            if(res[0] === "@"){
                var xp = res.substr(1, res.length);

                if (xp.length > 0 ){
                    var asp = e.target.getAttribute("gx")
                    textarea.value = textarea.value.substr(0, k-res.length)+"@"+asp+textarea.value.substr(k);textarea.focus();
                    textarea.setSelectionRange(2,6);
                }
            }
        }
    }
});


                

// document.addEventListener("keyup", function(e){
//     if(e.target.id === "frmpst27890editabl"){
//         var v = e.target.value;var k = e.target.selectionStart;var s = [];var rgx = v.match(/\n/g);
        
//         if(v[k-1] === "@"){
//             s.push(k-1)
//         }
        
//         var kl = v.substr(s[0], k);var sp = kl.split(" ");

//         var sk;

//         if (sp[0][0] === "@" && sp[sp.length-1][0] === "@"){
//             sk = sp[0];
//         }else{
//             sk = sp[sp.length-1];
//         }
        

//         if(sk.length > 0 ){
//             if(sk[0] === "@"){

//                 var xp = sk.substr(1, sk.length);

//                 if (xp.length > 0 ){
//                     $(".nwdngxr092").html("<div style='width:fit-content;margin:auto;padding:10px 10px 10px 10px;'> <div class='loader'></div> </div>")
//                     $.post("/xoz/qr1/tg0x",
//                     {
//                         name:xp,
                 
//                     },
//                     function(d,s){
//                         if(s === "success"){
//                             var jsn = JSON.parse(d);
//                             for( const key in jsn){$("."+key).html(jsn[key])}
//                         }else{
//                             alert("something went wrong")
//                         }
//                    });
//                 }
//             }else{
//                 $(".nwdngxr092").html("");
//             }

//         }else{
//             $(".nwdngxr092").html("");
//         }
    
//     }
// })
