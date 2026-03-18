window.addEventListener("load", function load(event){

    const null_island = [ 0.0, 0.0 ];
    
    fetch("/map.json")
        .then((rsp) => rsp.json())
        .then((cfg) => {

	    var leaflet_map;
	    var maplibre_map;
	    
	    console.debug("Got map config", cfg);
	    
            switch (cfg.provider) {
                case "leaflet":
		    
		    leaflet_map = L.map('map').setView(null_island, 1);    		    
                    var tile_url = cfg.tile_url;

                    var tile_layer = L.tileLayer(tile_url, {
                        maxZoom: 19,
                    });

                    tile_layer.addTo(leaflet_map);
                    break;

		case "protomaps":

		    leaflet_map = L.map('map').setView(null_island, 1);

		    var tile_args = {
			url: cfg.tile_url,
		    };

		    if (cfg.protomaps.theme){
			tile_args.flavor = cfg.protomaps.theme;
		    }

		    if (cfg.protomaps.paintRules){
			tile_args.paintRules = cfg.protomaps.paintRules;
		    }

		    if (cfg.protomaps.labelRules){
			tile_args.labelRules = cfg.protomaps.labelRules;
		    }
		    
		    var tile_layer = protomapsL.leafletLayer(tile_args);
		    
		    tile_layer.addTo(leaflet_map);
		    break;


                case "protomaps-raster":
		    
		    leaflet_map = L.map('map').setView(null_island, 1);    		    		    
                    var tile_url = cfg.tile_url;

		    const p = new pmtiles.PMTiles(
			cfg.tile_url,
		    );
		    
		    p.getHeader().then((h) => {
			
			let tile_layer = pmtiles.leafletRasterLayer(p, {
			    maxNativeZoom: h.maxZoom
			});
			
			tile_layer.addTo(leaflet_map);
		    });
			
		    break;

		case "protomaps-ml":

		    let protocol = new pmtiles.Protocol({metadata: true});
		    maplibregl.addProtocol("pmtiles", protocol.tile);

		    var center = null_island;

		    var zoom = 3;
		    
		    if (cfg.initial_zoom){
			center = cfg.initial_view;
		    }

		    if (cfg.initial_zoom){
			zoom = cfg.initial_zoom;
		    }
		    
		    maplibre_map = new maplibregl.Map({
			container: "map",
			zoom: zoom,
			center: center,
			style: {
			    version: 8,
			    sources: {
				example_source: {
				    type: "vector",
				    // For standard Z/X/Y tile APIs or Z/X/Y URLs served from go-pmtiles, replace "url" with "tiles" and remove all the pmtiles-related client code.
				    // tiles: ["https://example.com/{z}/[x}/{y}.mvt"],
				    // see https://maplibre.org/maplibre-style-spec/sources/#vector
				    url: "pmtiles://" + location.toString().trimEnd("/") + cfg.tile_url,
				},
			    },
			    layers: [
				{
				    id: "water",
				    source: "example_source",
				    "source-layer": "water",
				    filter: ["==",["geometry-type"],"Polygon"],
				    type: "fill",
				    paint: {
					"fill-color": "#80b1d3",
				    },
				},
				{
				    id: "buildings",
				    source: "example_source",
				    "source-layer": "buildings",
				    type: "fill",
				    paint: {
					"fill-color": "#d9d9d9",
				    },
				},
				{
				    id: "roads",
				    source: "example_source",
				    "source-layer": "roads",
				    type: "line",
				    paint: {
					"line-color": "#fc8d62",
				    },
				},
				{
				    id: "pois",
				    source: "example_source",
				    "source-layer": "pois",
				    type: "circle",
				    paint: {
					"circle-color": "#ffffb3",
				    },
				},
			    ],
			},
		    });
		    
		    // map.showTileBoundaries = true;
		    break;
		    
                default:
                    console.error("Uknown or unsupported map provider");
                    return;
	    }

	    if (leaflet_map){
		
		if (cfg.initial_view) {

		    var zm = leaflet_map.getZoom();
		    
		    if (cfg.initial_zoom){
			zm = cfg.initial_zoom;
		    }
		    
		    leaflet_map.setView([cfg.initial_view[1], cfg.initial_view[0]], zm);
		    
		} else if (cfg.initial_bounds){
		    
		    var bounds = [
			[ cfg.initial_bounds[1], cfg.initial_bounds[0] ],
			[ cfg.initial_bounds[3], cfg.initial_bounds[2] ],
		    ];
		    
		    leaflet_map.fitBounds(bounds);
		}
	    }
	    
	    console.debug("Finished map setup");
	    
        }).catch((err) => {
	    console.error("Failed to derive map config", err);
	    return;
	});    
    
});
