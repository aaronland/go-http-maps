window.addEventListener("load", function load(event){

    console.log("LOAD");
    
    var map_el = document.getElementById("map");
    var map = aaronland.maps.getMap(map_el);

    var hash = new L.Hash(map);
    var hash_str = location.hash;

    var init_lat = map_el.getAttribute("data-initial-latitude");
    var init_lon = map_el.getAttribute("data-initial-longitude");
    var init_zoom = map_el.getAttribute("data-initial-zoom");
    
    if (hash_str){

	var parsed = aaronland.maps.parseHash(hash_str);

	if (parsed){
	    init_lat = parsed['latitude'];
	    init_lon = parsed['longitude'];
	    init_zoom = parsed['zoom'];
	}
    }

    map.setView([init_lat, init_lon], init_zoom);
    
    // https://leaflet.github.io/Leaflet.draw/docs/leaflet-draw-latest.html
    // https://leaflet.github.io/Leaflet.draw/docs/leaflet-draw-latest.html#l-draw
    
    var drawnItems = new L.FeatureGroup();
    
    map.addLayer(drawnItems);
    
    var drawControl = new L.Control.Draw({
        edit: {
            featureGroup: drawnItems
        }
    });
    
    map.addControl(drawControl);

    map.on(L.Draw.Event.CREATED, function (e){
        var type = e.layerType;
        var layer = e.layer;

	console.log("EDIT", type, layer);
	
        if (type === 'marker') {
            layer.bindPopup('A popup!');
        }
	
        drawnItems.addLayer(layer);

	console.log("GEOJSON", drawnItems.toGeoJSON());
    });

    map.on(L.Draw.Event.EDITED, function (e){
        var layers = e.layers;
	
        var countOfEditedLayers = 0;
	
        layers.eachLayer(function(layer) {

	    console.log("LAYER", layer);
            countOfEditedLayers++;
        });
	
        console.log("Edited " + countOfEditedLayers + " layers");
    });
});