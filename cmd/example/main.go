package main

import (
	"flag"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"net/url"
	"strings"

	"github.com/aaronland/go-http-maps/v2"
	"github.com/aaronland/go-http-maps/v2/static/www"
)

func main() {

	var verbose bool

	var host string
	var port int

	var map_provider string
	var map_tile_uri string
	var protomaps_theme string

	var style string

	flag.StringVar(&host, "host", "localhost", "The host to listen for requests on")
	flag.IntVar(&port, "port", 8080, "The port number to listen for requests on")
	flag.StringVar(&map_provider, "map-provider", "leaflet", "Valid options are: leaflet, protomaps")
	flag.StringVar(&map_tile_uri, "map-tile-uri", maps.LEAFLET_OSM_TILE_URL, "A valid Leaflet tile layer URI. See documentation for special-case (interpolated tile) URIs.")
	flag.StringVar(&protomaps_theme, "protomaps-theme", "white", "A valid Protomaps theme label.")

	flag.StringVar(&style, "style", "", "A custom Leaflet style definition for geometries. This may either be a JSON-encoded string or a path on disk.")

	flag.BoolVar(&verbose, "verbose", false, "Enable verbose (debug) logging.")

	flag.Parse()

	if verbose {
		slog.SetLogLoggerLevel(slog.LevelDebug)
		slog.Debug("Verbose logging enabled")
	}

	// ctx := context.Background()

	mux := http.NewServeMux()

	map_cfg := &maps.MapConfig{
		Provider: map_provider,
		TileURL:  map_tile_uri,
		// Style:           style,
	}

	if map_provider == "protomaps" {

		u, err := url.Parse(map_tile_uri)

		if err != nil {
			log.Fatalf("Failed to parse Protomaps tile URL, %v", err)
		}

		switch u.Scheme {
		case "file":

			mux_url, mux_handler, err := maps.ProtomapsFileHandlerFromPath(u.Path, "")

			if err != nil {
				log.Fatalf("Failed to determine absolute path for '%s', %v", map_tile_uri, err)
			}

			mux.Handle(mux_url, mux_handler)
			map_cfg.TileURL = mux_url

		case "api":
			key := u.Host
			map_cfg.TileURL = strings.Replace(maps.PROTOMAPS_API_TILE_URL, "{key}", key, 1)
		}

		map_cfg.Protomaps = &maps.ProtomapsConfig{
			Theme: protomaps_theme,
		}
	}

	map_cfg_handler := maps.MapConfigHandler(map_cfg)

	mux.Handle("/map.json", map_cfg_handler)

	www_fs := http.FS(www.FS)
	www_handler := http.FileServer(www_fs)

	mux.Handle("/", www_handler)

	addr := fmt.Sprintf("%s:%d", host, port)
	slog.Info("Listening for requests", "address", addr)

	err := http.ListenAndServe(addr, mux)

	if err != nil {
		log.Fatalf("Failed to serve requests, %v", err)
	}

}
