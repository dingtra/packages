package makeding



import (
	// 
)


func (let *MakeDingStruct ) GetFiles () string {

	return `
		<div class='flchngxr090' style='display:none;' ><span><input type='file' id='filer092' name='files' accept='video/*,image/*' multiple/></span></div>
		<div class="flxjss090">
			<div class="flchngxr091"> 
				<label for='filer092'>
					<div id='mkdnglabl090' class='mkdnglabelx090'>
						<i data-feather='cloud'></i> <span> Add files here</span>
					</div>
					<div id='mkdnglabl091' class='mkdnglabelx091'>
						<span>Atleast 15 MB for videos & 5mb for photos, we are currently on test stage</span>
					</div>
				</label>
			</div>
	
			<div style='display:none;' id='fl091281'><div style='display:flex;overflow: auto;white-space: nowrap;width:fit-content;margin:auto;' id='fl091281ch'></div></div>
			<div style='display:none;' id='fl091280'><video id='vdfl091280' controls >not supported</video></div>
		</div>
		<div class="setxkjss090"></div>
		<script>
	      feather.replace()
        </script>
	`

}