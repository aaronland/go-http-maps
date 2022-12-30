[
    {
        dataLayer:"water",
        symbolizer:new protomaps.PolygonSymbolizer({fill:"#354855"})
    },
    {
	dataLayer: "roads",
	symbolizer: new protomaps.LineSymbolizer({color:"#fff"}),
    },
    {
	dataLayer: "landuse",
        symbolizer:new protomaps.PolygonSymbolizer({fill:"#cccccc"})
    },
    {
	dataLayer: "landuse",
        symbolizer:new protomaps.PolygonSymbolizer({fill:"#999"}),
	filter: (props, ignore) => {
	    
	    if (props["area:aeroway"] == "runway"){
		return true;
	    }
	    
	    if (props["area:aeroway"] == "taxiway"){
		return true;
	    }
	    
	    if (props["aeroway"] == "runway"){
		return true;
	    }
	    
	    if (props["aeroway"] == "aerodrome"){
		return true;
	    }
	    
	    return false;
	}
    },
    {
	dataLayer: "transit",
	symbolizer: new protomaps.LineSymbolizer({color:"#000"}),
	filter: (props, ignore) => {
	    
	    if (props["pmap:kind"] = "aeroway"){
		return true;
	    }
	    
	    return false;
	}
    }
];
