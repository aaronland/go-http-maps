{{ define "rules" }}
var aaronland = aaronland || {};
var aaronland.maps = aaronland.maps || {};
var aaronland.maps.protomaps = aaronland.maps.protomaps || {};

aaronland.maps.protomaps.rules = (function(){

    var self = {
	'paintRules': function(){
	    return {{ .PaintRules }};
	},

	'labelRules': function(){
	    return {{ .LabelRules }};
	},
    };

    return self
}
{{ end }}
