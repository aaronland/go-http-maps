package main

import (
	"context"
	"github.com/aaronland/go-http-bootstrap"
	"github.com/aaronland/go-http-leaflet"
	"github.com/aaronland/go-http-maps/http"
	"github.com/aaronland/go-http-maps/templates/html"
	"github.com/aaronland/go-http-server"
	"github.com/aaronland/go-http-tangramjs"
	"github.com/sfomuseum/go-flags/flagset"
	"github.com/sfomuseum/go-http-protomaps"
	"html/template"
	"log"
	gohttp "net/http"
)

func init() {

	tangramjs.APPEND_LEAFLET_RESOURCES = false
	tangramjs.APPEND_LEAFLET_ASSETS = false
	protomaps.APPEND_LEAFLET_RESOURCES = false
	protomaps.APPEND_LEAFLET_ASSETS = false
}

func main() {

	fs := flagset.NewFlagSet("privatezen")

	server_uri := fs.String("server-uri", "http://localhost:8080", "A valid aaronland/go-http-server URI")

	map_renderer := fs.String("map-renderer", "", "Valid options are: protomaps, tangramjs")

	nextzen_apikey := fs.String("nextzen-apikey", "", "A valid Nextzen API key")
	nextzen_style_url := fs.String("nextzen-style-url", "/tangram/refill-style.zip", "A valid URL for loading a Tangram.js style bundle.")
	nextzen_tile_url := fs.String("nextzen-tile-url", tangramjs.NEXTZEN_MVT_ENDPOINT, "A valid Nextzen tile URL template for loading map tiles.")

	initial_latitude := fs.Float64("initial-latitude", 37.61799, "A valid latitude for the map's initial view.")
	initial_longitude := fs.Float64("initial-longitude", -122.370943, "A valid longitude for the map's initial view.")
	initial_zoom := fs.Int("initial-zoom", 15, "A valid zoom level for the map's initial view.")

	protomaps_tile_url := fs.String("protomaps-tile-url", "", "A valid Protomaps .pmtiles URL for loading map tiles.")

	flagset.Parse(fs)

	ctx := context.Background()

	s, err := server.NewServer(ctx, *server_uri)

	if err != nil {
		log.Fatalf("Failed to create new server for '%s', %v", *server_uri, err)
	}

	t := template.New("geotag").Funcs(template.FuncMap{
		//
	})

	t, err = t.ParseFS(html.FS, "*.html")

	if err != nil {
		log.Fatalf("Failed to parse templates, %v", err)
	}

	mux := gohttp.NewServeMux()

	err = bootstrap.AppendAssetHandlers(mux)

	if err != nil {
		log.Fatalf("Failed to append Bootstrap asset handlers, %v")
	}

	err = http.AppendStaticAssetHandlers(mux)

	if err != nil {
		log.Fatalf("Failed to append static asset handlers, %v")
	}

	err = leaflet.AppendAssetHandlers(mux)

	if err != nil {
		log.Fatalf("Failed to append Leaflet asset handlers, %v", err)
	}

	map_opts := &http.MapHandlerOptions{
		Templates:        t,
		MapRenderer:      *map_renderer,
		InitialLatitude:  *initial_latitude,
		InitialLongitude: *initial_longitude,
		InitialZoom:      *initial_zoom,
	}

	map_handler, err := http.MapHandler(map_opts)

	if err != nil {
		log.Fatalf("Failed to create map handler, %v", err)
	}

	bootstrap_opts := bootstrap.DefaultBootstrapOptions()
	map_handler = bootstrap.AppendResourcesHandler(map_handler, bootstrap_opts)

	leaflet_opts := leaflet.DefaultLeafletOptions()
	map_handler = leaflet.AppendResourcesHandler(map_handler, leaflet_opts)

	switch *map_renderer {
	case "protomaps":

		pm_opts := protomaps.DefaultProtomapsOptions()
		pm_opts.TileURL = *protomaps_tile_url

		map_handler = protomaps.AppendResourcesHandler(map_handler, pm_opts)

		err = protomaps.AppendAssetHandlers(mux)

		if err != nil {
			log.Fatalf("Failed to append Protomaps asset handlers, %v")
		}

	case "tangramjs":

		tangramjs_opts := tangramjs.DefaultTangramJSOptions()
		tangramjs_opts.Nextzen.APIKey = *nextzen_apikey
		tangramjs_opts.Nextzen.StyleURL = *nextzen_style_url
		tangramjs_opts.Nextzen.TileURL = *nextzen_tile_url

		map_handler = tangramjs.AppendResourcesHandler(map_handler, tangramjs_opts)

		err = tangramjs.AppendAssetHandlers(mux)

		if err != nil {
			log.Fatalf("Failed to append TangramJS asset handlers, %v")
		}

	default:
		log.Fatalf("Invalid or unsupporter map renderer '%s'", *map_renderer)
	}

	mux.Handle("/", map_handler)

	//

	log.Printf("Listening on %s\n", s.Address())
	err = s.ListenAndServe(ctx, mux)

	if err != nil {
		log.Fatalf("Failed to start server, %v", err)
	}
}
