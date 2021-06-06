package provider

import (
	"flag"
	"fmt"
	"github.com/aaronland/go-http-leaflet"
	"github.com/aaronland/go-http-tangramjs"
	"github.com/sfomuseum/go-http-protomaps"
	_ "log"
	"strings"
)

const MapProviderFlag string = "map-provider"

var map_provider string

const LeafletEnableHashFlag string = "leaflet-enable-hash"

var leaflet_enable_hash bool

const LeafletEnableFullscreenFlag string = "leaflet-enable-fullscreen"

var leaflet_enable_fullscreen bool

const LeafletEnableDrawFlag string = "leaflet-enable-draw"

var leaflet_enable_draw bool

const NextzenAPIKeyFlag string = "nextzen-apikey"

var nextzen_apikey string

const NextzenStyleURLFlag string = "nextzen-style-url"

var nextzen_style_url string

const NextzenTileURLFlag string = "nextzen-tile-url"

var nextzen_tile_url string

const ProtomapsTileURLFlag string = "protomaps-tile-url"

var protomaps_tile_url string

const TilezenEnableTilepack string = "tilezen-enable-tilepack"

var tilezen_enable_tilepack bool

const TilezenTilepackPath string = "tilezen-tilepack-path"

var tilezen_tilepack_path string

const InitialLatitude string = "initial-latitude"

var initial_latitude float64

const InitialLongitude string = "initial-longitude"

var initial_longitude float64

const InitialZoom string = "initial-zoom"

var initial_zoom int

func AppendProviderFlags(fs *flag.FlagSet) error {

	fs.StringVar(&map_provider, MapProviderFlag, "", "...")

	fs.BoolVar(&leaflet_enable_hash, LeafletEnableHashFlag, true, "Enable the Leaflet.Hash plugin.")
	fs.BoolVar(&leaflet_enable_fullscreen, LeafletEnableFullscreenFlag, false, "Enable the Leaflet.Fullscreen plugin.")
	fs.BoolVar(&leaflet_enable_draw, LeafletEnableDrawFlag, false, "Enable the Leaflet.Draw plugin.")

	fs.StringVar(&nextzen_apikey, NextzenAPIKeyFlag, "", "A valid Nextzen API key")
	fs.StringVar(&nextzen_style_url, NextzenStyleURLFlag, "/tangram/refill-style.zip", "A valid URL for loading a Tangram.js style bundle.")
	fs.StringVar(&nextzen_tile_url, NextzenTileURLFlag, tangramjs.NEXTZEN_MVT_ENDPOINT, "A valid Nextzen tile URL template for loading map tiles.")

	fs.StringVar(&protomaps_tile_url, ProtomapsTileURLFlag, "", "A valid Protomaps .pmtiles URL for loading map tiles.")

	fs.BoolVar(&tilezen_enable_tilepack, TilezenEnableTilepack, false, "Enable to use of Tilezen MBTiles tilepack for tile-serving.")
	fs.StringVar(&tilezen_tilepack_path, TilezenTilepackPath, "", "The path to the Tilezen MBTiles tilepack to use for serving tiles.")

	fs.Float64Var(&initial_latitude, InitialLatitude, 37.778008, "A valid latitude for the map's initial view.")
	fs.Float64Var(&initial_longitude, InitialLongitude, -122.431272, "A valid longitude for the map's initial view.")
	fs.IntVar(&initial_zoom, InitialZoom, 15, "A valid zoom level for the map's initial view.")

	return nil
}

func ProviderOptionsFromFlagSet(fs *flag.FlagSet) (*ProviderOptions, error) {

	opts := &ProviderOptions{
		InitialLatitude:  initial_latitude,
		InitialLongitude: initial_longitude,
		InitialZoom:      initial_zoom,
	}

	leaflet_opts := leaflet.DefaultLeafletOptions()

	if leaflet_enable_hash {
		leaflet_opts.EnableHash()
	}

	if leaflet_enable_fullscreen {
		leaflet_opts.EnableFullscreen()
	}

	if leaflet_enable_draw {
		leaflet_opts.EnableDraw()
	}

	opts.LeafletOptions = leaflet_opts

	switch strings.ToLower(map_provider) {
	case "protomaps":

		pm_opts := protomaps.DefaultProtomapsOptions()
		pm_opts.TileURL = protomaps_tile_url

		opts.ProtomapsOptions = pm_opts
		opts.Provider = ProtomapsProvider

	case "tangramjs":

		tangramjs_opts := tangramjs.DefaultTangramJSOptions()
		tangramjs_opts.NextzenOptions.APIKey = nextzen_apikey
		tangramjs_opts.NextzenOptions.StyleURL = nextzen_style_url
		tangramjs_opts.NextzenOptions.TileURL = nextzen_tile_url

		opts.TangramJSOptions = tangramjs_opts
		opts.Provider = TangramJSProvider

	default:
		return nil, fmt.Errorf("Unknown or unsupported map provider '%s'", map_provider)
	}

	if tilezen_enable_tilepack {
		opts.TilezenEnableTilepack = true
		opts.TilezenTilepackPath = tilezen_tilepack_path
		opts.TilezenTilepackURL = "/tilezen/"

		if strings.ToLower(map_provider) == "tangramjs" {
			opts.TangramJSOptions.NextzenOptions.TileURL = "/tilezen/vector/v1/512/all/{z}/{x}/{y}.mvt"
		}
	}

	return opts, nil
}
