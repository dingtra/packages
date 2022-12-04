package makeding


import (
	"strings"
	"html/template"
	"encoding/json"
)


func (let *MakeDingStruct ) Polls ()  {
	Result := map[string]template.HTML{}

	details := []string{}

	// ------------------------------- polls
	details = append(details, "<div class='pollsxy91x89'>")

	details = append(details, "<div class='pollsxy92x89' id='xl09br5409' x='/xoz/qr1/sk0x'><div style='pointer-events:none;' class='pollxr09'><i data-feather='x'></i></div></div>")

	// 
	details = append(details, "<div class='pollsxy93x89'>")


	// -------------
	details = append(details, "<div class='pollsxy93x89xr'>")
	details = append(details, "<span><input type='text' name='poll1' placeholder='Option 1' class='pollx-box'></span>")
	details = append(details, "</div>")
	// ------------

	// -------------
	details = append(details, "<div class='pollsxy93x89xr'>")
	details = append(details, "<span><input type='text' name='poll2' placeholder='Option 2' class='pollx-box'></span>")
	details = append(details, "</div>")
	// ------------

	// -------------
	details = append(details, "<div class='pollsxy93x89xr'>")
	details = append(details, "<span><input type='text' name='poll3' placeholder='Option 3' class='pollx-box'></span>")
	details = append(details, "</div>")
	// ------------

	// -------------
	details = append(details, "<div class='pollsxy93x89xr'>")
	details = append(details, "<span><input type='text' name='poll4' placeholder='Option 4' class='pollx-box'></span>")
	details = append(details, "</div>")
	// ------------


	details = append(details, "</div>")
	// 


	// length
	details = append(details, "<div class='pollsxy94x89'>")

	details = append(details, "<div class='pollsxy94x89-polls'><span>Poll timing</div>")

	details = append(details, "<div class='pollsxy94x89-box'>"+PollSelection1()+"</div>")

	details = append(details, "</div>")
	// length



	details = append(details, "</div>"+
	`
	<script>
	  feather.replace()
    </script>
	`)
	// ------------------------------ polls

	Result["xro9polx89"] = template.HTML(strings.Join(details, ""))

	json, err := json.Marshal(Result)

	if err != nil {
		panic(err)
	}

	let.JsonVal = string(json)
}



func PollSelection1 () string {

	return strings.Join([]string{
		`
		<div class='pollsxy94x89xr'>
			<div class='selector-x01-label'>
			    <div class='selector-x01'><span>Days</span></div>
				<select name="days">
				   <option value='0'>0</option>
				   <option value='1'>1</option>
				   <option value='2'>2</option>
				   <option value='3'>3</option>
				   <option value='4'>4</option>
				   <option value='5'>5</option>
				   <option value='6'>6</option>
				   <option value='7'>7</option>
				   </select>
			</div>
		</div>

		<div class='pollsxy94x89xr'>
			<div class='selector-x01-label'>
			    <div class='selector-x01'><span>Hours</span></div>
				<select name="hours">
					<option value='0'>0</option>
					<option value='1'>1</option>
					<option value='2'>2</option>
					<option value='3'>3</option>
					<option value='4'>4</option>
					<option value='5'>5</option>
					<option value='6'>6</option>
					<option value='7'>7</option>
					<option value='8'>8</option>
					<option value='9'>9</option>
					<option value='10'>10</option>
					<option value='11'>11</option>
					<option value='12'>12</option>
					<option value='13'>13</option>
					<option value='14'>14</option>
					<option value='15'>15</option>
					<option value='16'>16</option>
					<option value='17'>17</option>
					<option value='18'>18</option>
					<option value='19'>19</option>
					<option value='20'>20</option>
					<option value='21'>21</option>
					<option value='22'>22</option>
					<option value='23'>23</option>
					<option value='24'>24</option>
				</select>
			</div>
		</div>


		<div class='pollsxy94x89xr'>
			<div class='selector-x01-label'>
			    <div class='selector-x01'><span>Minutes</span></div>
				<select name="minutes">
					<option value='0'>0</option>
					<option value='1'>1</option>
					<option value='2'>2</option>
					<option value='3'>3</option>
					<option value='4'>4</option>
					<option value='5'>5</option>
					<option value='6'>6</option>
					<option value='7'>7</option>
					<option value='8'>8</option>
					<option value='9'>9</option>
					<option value='10'>10</option>
					<option value='11'>11</option>
					<option value='12'>12</option>
					<option value='13'>13</option>
					<option value='14'>14</option>
					<option value='15'>15</option>
					<option value='16'>16</option>
					<option value='17'>17</option>
					<option value='18'>18</option>
					<option value='19'>19</option>
					<option value='20'>20</option>
					<option value='21'>21</option>
					<option value='22'>22</option>
					<option value='23'>23</option>
					<option value='24'>24</option>
					<option value='25'>25</option>
					<option value='26'>26</option>
					<option value='27'>27</option>
					<option value='28'>28</option>
					<option value='29'>29</option>
					<option value='30'>30</option>
					<option value='31'>31</option>
					<option value='32'>32</option>
					<option value='33'>33</option>
					<option value='34'>34</option>
					<option value='35'>35</option>
					<option value='36'>36</option>
					<option value='37'>37</option>
					<option value='38'>38</option>
					<option value='39'>39</option>
					<option value='40'>40</option>
					<option value='41'>41</option>
					<option value='42'>42</option>
					<option value='43'>43</option>
					<option value='44'>44</option>
					<option value='45'>45</option>
					<option value='46'>46</option>
					<option value='47'>47</option>
					<option value='48'>48</option>
					<option value='49'>49</option>
					<option value='50'>50</option>
					<option value='51'>51</option>
					<option value='52'>52</option>
					<option value='53'>53</option>
					<option value='54'>54</option>
					<option value='55'>55</option>
					<option value='56'>56</option>
					<option value='57'>57</option>
					<option value='58'>58</option>
					<option value='59'>59</option>
				</select>
			</div>
		</div>
			
		`,
	}, "")
}
