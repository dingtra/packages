function Ptag(i, n){
	return [
        "<div class='sellres901' id='btnsell' nx='"+i+"'><div class='img9780'>",
        "<div style='padding:5px 0px 5px 0px;font-size:13px;'><span>Add Title</span></div><div id='img9780input'><input style='height:30px;border-radius:4px;border:1px solid #ccc;padding-left:10px;width:100%;' type='text' name='"+n+"ungogo' placeholder='Product Name'/></div>",
        "<div style='padding:5px 0px 5px 0px;font-size:13px;'><span>Add Price</span></div><div  id='img9780input'><div style='display:flex;'><select name='"+n+"currency'><option>NGN</option><option>USD</option></select><input style='height:30px;border-radius:4px;border:1px solid #ccc;padding-left:10px;width:100%;' type='number' placeholder='Price' name='"+n+"'/></div></div>",
        "</div></div>"
	].join("");
}


function CompileFile(src, i, n){
    var files = document.getElementById("filer092-bssnss").files;var arrange;
    if(files.length > 1){
        arrange = `
        <div id='kngo' style="position:absolute;top:10;left:10;padding:10px 10px 10px 10px;background: rgba(0, 0, 0, 0.5);color:white;" class="imgupljss091">
            <div style="pointer-events:none;">
                <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="feather feather-grid" style="height:15px;width:15px;position:relative;top:1px;" color="white"><rect x="3" y="3" width="7" height="7"></rect><rect x="14" y="3" width="7" height="7"></rect><rect x="14" y="14" width="7" height="7"></rect><rect x="3" y="14" width="7" height="7"></rect></svg>
            </div>
        </div>
        `
    }else{
        arrange = ""
    }
    
	return [
        `
           
            <div id='imgupljss090`+n.replaceAll(".", "_")+`'>
                <div style="position:fixed;overflow:auto;top:0px;left:0px;right:0px;background:red;z-index:10;" class='attachxr09`+n.replace(".", "_") +`'></div>
                <div style="flex-direction:column;">

                    <div style="position:relative;z-index:5;" class='imgupljss090'>
                       <div id="cnclimgx009" file='`+n+`' style="position:absolute;top:10;right:10;padding:10px 10px 10px 10px;background: rgba(0, 0, 0, 0.5);color:white;" class="imgupljss091" x='`+n.replaceAll(".", "_")+`'>
                            <div style="pointer-events:none;">
                                <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="feather feather-x" color="white"><line x1="18" y1="6" x2="6" y2="18"></line><line x1="6" y1="6" x2="18" y2="18"></line></svg>
                            </div>
                       </div>
                       `+arrange+`
                       
                       <div style="pointer-events:none;">
                           <img  src='`+src+`'/>
                       </div>

                       <div  style="position:absolute;bottom:0;left:10;right:10;padding:10px 10px 10px 10px;background: rgba(0, 0, 0, 0.5);color:white;"  class="imgxr090" x='`+n+`' src='`+src+`'>
                            <div style="pointer-events:none;width:fit-content;margin:auto;">
                                <svg xmlns="http://www.w3.org/2000/svg"  viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="feather feather-grid" style="height:100px;width:100px;color:white;stroke-width:1;" color="black"><rect x="3" y="3" width="7" height="7"></rect><rect x="14" y="3" width="7" height="7"></rect><rect x="14" y="14" width="7" height="7"></rect><rect x="3" y="14" width="7" height="7"></rect></svg>
                            </div>
                       </div>
                    </div>
                </div>
                
            </div>  
        `
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
        document.getElementById("fl091281ch").innerHTML += CompileFile(this.result,indx, fle.name)
        $("#fl091281").show();$("#mkdnglabl090").hide();$("#mkdnglabl091").hide();
        tgglr()
    }, false);
    rdr.readAsDataURL(fle);
}



function NormCmpl (fle, indx, itm) {
    var rdr = new FileReader();
    rdr.addEventListener("load", function () {
        var arrange = "";
        var files = document.getElementById("filer092").files;
        if(files.length > 1){
            arrange = `
            <div id='kngo' style="position:absolute;top:10;left:10;padding:10px 10px 10px 10px;background: rgba(0, 0, 0, 0.5);color:white;" class="imgupljss091">
                <div style="pointer-events:none;">
                    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="feather feather-grid" style="height:15px;width:15px;position:relative;top:1px;" color="white"><rect x="3" y="3" width="7" height="7"></rect><rect x="14" y="3" width="7" height="7"></rect><rect x="14" y="14" width="7" height="7"></rect><rect x="3" y="14" width="7" height="7"></rect></svg>
                </div>
            </div>
            `
        }
        var flname = fle.name.replaceAll(" ", "_");
        $("#vdfl091280").trigger('pause');$("#vdfl091280").val("");$("#fl091280").hide();$("#fl091282").hide();
        document.getElementById("fl091281ch").innerHTML += `
            <div id='imgupljss090`+flname.replaceAll(".", "_")+`'><div style="position:relative;" class='imgupljss090'>
               <div id="cnclimgx009" file='`+flname+`' style="position:absolute;top:10;right:10;padding:10px 10px 10px 10px;background: rgba(0, 0, 0, 0.5);color:white;" class="imgupljss091" x='`+flname.replaceAll(".", "_")+`'>
                    <div style="pointer-events:none;">
                        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="feather feather-x" color="white"><line x1="18" y1="6" x2="6" y2="18"></line><line x1="6" y1="6" x2="18" y2="18"></line></svg>
                    </div>
               </div>
               `+arrange+`
               <div class="imgupljss092">
                    <img src='`+this.result+`'/>
               </div>
            </div></div>
            <script>
                feather.replace()
            </script>
         `;
        $("#fl091281").show();$("#mkdnglabl090").hide();$("#mkdnglabl091").hide();
        tgglr()
    }, false);
    rdr.readAsDataURL(fle);
}

$(document).ready(function(){
    $(document).click(function(e){
        if(e.target.id === "kngo"){
            var collect = 0
            var arrs = [];
            var files;
            if(document.querySelector("#filer092")){
                files = document.getElementById("filer092");
            }else if(document.querySelector("#filer092-bssnss")){
                files = document.getElementById("filer092-bssnss");
            }
            let width = window.innerWidth;


            if(files.files.length > 1){
                var cuntfls = $(".setxkjss090").html();
                if(cuntfls.length > 0){
                    $(".setxkjss090").show();
                }else{
                    localStorage.removeItem("files");
                    for (let i = 0; i < files.files.length; i++) {
                        const file = files.files[i];
                        collect += 1
                       
                        var rdr = new FileReader();
                        rdr.addEventListener("load", function () {
        
                            arrs.push(`
                                <div style="position:relative;" class="xrklimgooxl" id='ksp-`+file.name+`'>
                                   <img id="imgbx0090" style="border-radius:4px;" x='`+file.name+`' src='`+this.result+`' />
                                </div>
                            `)
        
                            if(arrs.length == files.files.length) {

                                if(width < 600) {
                                    $(".setxkjss090").html(`
                                        <div style="position:fixed;overflow:auto;background:white;z-index:100;top:0px;left:0;right:0;bottom:0;background:white;border:1px solid #ccc;padding:10px 10px 10px 10px;border-radius:5px;box-shadow: 0 3px 10px rgb(0 0 0 / 0.1)"  class="arrng2jss090">
                                            <div style="display:flex;justify-content:space-between;padding:5px 5px 15px 5px;">
                                                <div style="background:#f2f2f2;border:1px solid #ccc;padding:5px 30px 5px 30px;border-radius:4px;" id="kngo0x0">
                                                    <div style="pointer-events:none;">Restart</div>
                                                </div>
                                                <div id="kngo0x1">
                                                    <div style="pointer-events:none;"><i data-feather='x' color='black'></i></div>
                                                </div>
                                            </div>
                    
                                            <div class="arrng2jss092">
                                                <div style="width:fit-content;margin:auto;" class="arrng2jss093">
                                                  <i style="height:15px;width:15px;position:relative;top:2px;" data-feather='grid' color='black'></i> <span style="font-size:18px;">Click to Arrange files</span>
                                                </div>
                                                <div style="width:`+(width/2+100)+`px;margin:auto;" class="arrng2jss094">
                                                    `+ arrs.join("")+`
                                                </div>
                                            </div>
                                        </div>
                                        <script>
                                            feather.replace()
                                        </script>
                                        
                                    `)
    
                                }else{
                                    $(".setxkjss090").html(`
                                        <div style="position:fixed;overflow:auto;background:white;z-index:100;top:0px;left:0;right:0;bottom:0;background:white;border:1px solid #ccc;padding:20px 20px 20px 20px;border-radius:5px;box-shadow: 0 3px 10px rgb(0 0 0 / 0.1);"  class="arrng2jss090">
                                            <div style="display:flex;justify-content:space-between;padding:5px 5px 15px 5px;">
                                                <div style="background:#f2f2f2;border:1px solid #ccc;padding:5px 30px 5px 30px;border-radius:4px;" id="kngo0x0">
                                                    <div style="pointer-events:none;">Restart</div>
                                                </div>
                                                <div id="kngo0x1">
                                                    <div style="pointer-events:none;"><i data-feather='x' color='black'></i></div>
                                                </div>
                                            </div>
                    
                                            <div style="width:fit-content;margin:auto;" class="arrng2jss092">
                                                <div style="width:fit-content;margin:auto;" class="arrng2jss093">
                                                    <i style="height:15px;width:15px;position:relative;top:1px;" data-feather='grid' color='black'></i> <span style="font-size:18px;">Click to Arrange files</span>
                                                </div>
                                                <div style="display:flex;flex-wrap:wrap;" class="arrng2jss094">
                                                    `+ arrs.join("")+`
                                                </div>
                                            </div>
                                        </div>
                                        <script>
                                            feather.replace()
                                        </script>
                                    `)
                                }
                                $(".setxkjss090").show();
                            }
                        }, false)
                        rdr.readAsDataURL(file);
                    }
                }

            }   
         
        }else if(e.target.id === "kngo0x1" || e.target.id === "kngo0x0"){
            if(e.target.id === "kngo0x0"){
                var itms = localStorage.getItem('files');
                if(itms != null && itms.length > 0 ){
                    var json = JSON.parse(itms)
                    for(var keys in json){
                        var clickedElement = document.getElementById(json[keys].id)
                        var imgElement = clickedElement.querySelector("img");
                        // Get the src attribute value of the img element
                        var srcValue = imgElement.getAttribute("src");

                        document.getElementById(json[keys].id).innerHTML = `
                            <img id="imgbx0090" style="border-radius:4px;" x='`+json[keys].file+`' src='`+srcValue+`' />
                        `
                    }
                }
                localStorage.removeItem("files");    
            }else{
                $(".setxkjss090").hide();
            }
        }else if(e.target.id === "imgbx0090"){
            var objs = {};
            var x = e.target.getAttribute("x");var itms = localStorage.getItem('files');
            if(itms != null && itms.length > 0 ){
                var json = JSON.parse(localStorage.getItem("files"))
                if(json[x] != undefined){
                    console.log(localStorage.getItem("files"))
                }else{
                    document.getElementById("ksp-"+x).innerHTML = `
                        <div style="position:absolute;top:10px;left:10px;padding:10px 10px 10px 10px;width:fit-content;border-radius:100%;background:black;" class="countjxx0">
                           <div style="color:white;" class="countjxx1">
                                <span>`+(Object.keys(json).length + 1)+`</span>
                           </div>
                        </div>
                        `+document.getElementById("ksp-"+x).innerHTML+`
                    `
                    json[x] = {
                        "file":x, "id":"ksp-"+x, "position":(Object.keys(json).length + 1)
                    }
                    localStorage.setItem('files', JSON.stringify(json)); 
                }
                console.log(json)
            }else{
                document.getElementById("ksp-"+x).innerHTML = `
                    <div style="position:absolute;top:10px;left:15px;padding:10px 10px 10px 10px;width:fit-content;border-radius:100%;background:black;" class="countjxx0">
                       <div style="color:white;" class="countjxx1">
                            <span>`+(Object.keys(objs).length + 1)+`</span>
                       </div>
                    </div>
                    `+document.getElementById("ksp-"+x).innerHTML+`
                `

                objs[x] = {
                    "file":x, "id":"ksp-"+x, "position":(Object.keys(objs).length + 1)
                }

                localStorage.setItem('files', JSON.stringify(objs)); 
            }

        }
    })
})


$(document).ready(function(){
    $(document).click(function(e){
        if(e.target.classList.contains("imgxr090")){
            if( $(".attachxr09"+e.target.getAttribute("x").replace(".", "_")).html().length > 0) {
                $(".attachxr09"+e.target.getAttribute("x").replace(".", "_")).show();
            }else{
                $(".attachxr09"+e.target.getAttribute("x").replace(".", "_")).html("<div style='width:fit-content;margin:auto;height:25px;'> <div style='margin-top:2px;height:20px;width:20px;' class='loader'></div> </div>")
                $.post("/xoz/qr1/bssnss090",
                {
                    name:e.target.getAttribute("x"),
                    id:e.target.getAttribute("x").replace(".", "_"),
             
                },
                function(d,s){
                    if(s === "success"){
                        var jsn = JSON.parse(d);
                        for( const key in jsn){$("."+key).html(jsn[key])}
                        var res = $(".attachxr09"+e.target.getAttribute("x").replace(".", "_")).html()
                        var xr = res.replace(e.target.getAttribute("x")+":main", `<img style="border-radius:4px;" src='`+e.target.getAttribute("src")+`'/>`)
                        $(".attachxr09"+e.target.getAttribute("x").replace(".", "_")).html(xr)
                        ImgHeights()

                    }else{
                        alert("something went wrong")
                    }
               });
            }

        }else  if(e.target.classList.contains("imgxr090cncl")){
            $(".attachxr09"+e.target.getAttribute("x")).hide();
        }
 
    })
})


function ImgHeights(){
    let w = window.innerHeight;
    let h = window.innerWidth;
    if(h < 400) {
        $(".bssnssxr090-chld").css("height",w-100+"px")
    }else{
        $(".bssnssxr090-chld").css("height",w-70+"px")
    }
}
