package makeding



import (
	"strings"
	"html/template"
	"encoding/json"
)


func (let *MakeDingStruct ) GetFiles () {

	Result := map[string]template.HTML{}

	Result["xro9polx89"] = template.HTML(strings.Join([]string{
		`
		<div id='fl091282'><label for='filer092'><div id='create-new-ding-class-files01'><i data-feather='image'></i> <span> Add files here</span></div></label></div>
		<div style='display:none;' id='fl091281'><div style='display:flex;overflow: auto;white-space: nowrap;' id='fl091281ch'></div></div>

		<div style='display:none;' id='fl091280'><video id='vdfl091280' controls >not supported</video></div>
	
		<div style='display:none;' id='rmfile001'><div style='pointer-events:none;' class='rmfile'><div id='rmfile'><span>Remove file</span></div></div></div>

		<script>
	      feather.replace()
        </script>
		`,
	}, ""))

	json, err := json.Marshal(Result)

	if err != nil {
		panic(err)
	}

	let.JsonVal = string(json)

}