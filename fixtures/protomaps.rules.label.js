[
	{
            dataLayer: "places",
            symbolizer: new protomaps.CenteredTextSymbolizer({
                label_props:["name:en", "name"],
                fill:"black",
		stroke:"white",
		width:2,
                font:"500 14px sans-serif",
		lineHeight:1.3,
            }),
            filter: (z,f) => { 

		// console.log(z, f.props["pmap:kind"]);

		if ((z >= 3) && (f.props["pmap:kind"] == "country")){
		    return true;
		}

		if ((z >= 5) && (f.props["pmap:kind"] == "city")){

		    if ((z <= 10) && (f.props["population"] < 500000)){
			return false;
		    }
		    return true
		}
		
		if ((z >= 12) && (f.props["pmap:kind"] == "town")){
		    return true;
		}

		return false;
	    }
        },
	{
            dataLayer: "landuse",
            symbolizer: new protomaps.CenteredTextSymbolizer({
                label_props:["name:en", "name"],
                fill:"black",
		stroke:"white",
		width:2,
                font:"500 14px sans-serif",
		lineHeight:1.5,
            }),
            filter: (z,f) => {
		if (f.props["pmap:kind"] != "aerodrome"){
		    return false;
		}
		if (f.props["name"] == "San Francisco International Airport"){
		    return false;
		}
		return true;
	    }
        }

    ]
