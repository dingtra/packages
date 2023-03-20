document.addEventListener("click", function(e){
    if(e.target.id === "cnclimgx009"){
        var files;var x = e.target.getAttribute("x");var filesToKeep = [];var filetype;

        if(document.querySelector("#filer092")){
            files = document.getElementById("filer092");
            filetype = true
        }else if(document.querySelector("#filer092-bssnss")){
            files = document.getElementById("filer092-bssnss");
            filetype = false
        }

        for (var i = 0; i < files.files.length; i++) {
            // exclude clicked img
            if(files.files[i].name !=  e.target.getAttribute("file")){
                filesToKeep.push(files.files[i]);
            }
        }
        // create a new FileList object using DataTransfer
        var dt = new DataTransfer();

        if(filesToKeep.length > 0 ){
            for (var i = 0; i < filesToKeep.length; i++) {
              dt.items.add(filesToKeep[i]);
            }
        }

        if(filesToKeep.length > 0 ){
            // this triggers when the files are morethan one
            // if the item is clicked remaining item get pushed to filesToKeep[] array
            // then this block of condition becomes true 

            // remove exixting file
            $("#imgupljss090"+x).html("")
            files.files = dt.files;

            // emty class and remove saved files
            $(".setxkjss090").html("");
            localStorage.removeItem("files");
        }else{
            // when item is clicked and no other item to appened to filesToKeep[] array
            // this block becomes valid and modified files becomes empty

            // add modified Files
            files.files = dt.files;

            $.post("/xoz/qr1/optnsxr090",
            {data:filetype === true ? "files" : "bussiness",},
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
        

    }
})


document.addEventListener("click", function(e){
    if(e.target.id === "xl09bjj090"){
        $(".flxjss090 ").toggle();
    }
})