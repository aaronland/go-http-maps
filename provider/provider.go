package provider

import (
	"fmt"
	"github.com/aaronland/go-http-leaflet"
	"github.com/aaronland/go-http-tangramjs"
	"github.com/sfomuseum/go-http-protomaps"
	tilepack_http "github.com/tilezen/go-tilepacks/http"
	"github.com/tilezen/go-tilepacks/tilepack"
	_ "log"
	"net/http"
)

const (
	UnknownProvider Provider = iota
	TangramJSProvider
	ProtomapsProvider
)

type Provider int

func (p Provider) String() string {
	switch p {
	case TangramJSProvider:
		return "tangramjs"
	case ProtomapsProvider:
		return "protomaps"
	default:
		return "unknown"
	}
}

type ProviderOptions struct {
	Provider              Provider
	LeafletOptions        *leaflet.LeafletOptions
	TangramJSOptions      *tangramjs.TangramJSOptions
	ProtomapsOptions      *protomaps.ProtomapsOptions
	TilezenEnableTilepack bool
	TilezenTilepackPath   string
	TilezenTilepackURL    string
	InitialLatitude       float64
	InitialLongitude      float64
	InitialZoom           int
}

func init() {
	tangramjs.APPEND_LEAFLET_RESOURCES = false
	tangramjs.APPEND_LEAFLET_ASSETS = false
	protomaps.APPEND_LEAFLET_RESOURCES = false
	protomaps.APPEND_LEAFLET_ASSETS = false
}

func AppendResourcesHandler(handler http.Handler, opts *ProviderOptions) http.Handler {

	handler = leaflet.AppendResourcesHandler(handler, opts.LeafletOptions)

	switch opts.Provider {
	case ProtomapsProvider:

		handler = protomaps.AppendResourcesHandler(handler, opts.ProtomapsOptions)

	case TangramJSProvider:

		handler = tangramjs.AppendResourcesHandler(handler, opts.TangramJSOptions)

	default:
		// pass
	}

	return handler
}

func AppendAssetHandlers(mux *http.ServeMux, opts *ProviderOptions) error {

	err := leaflet.AppendAssetHandlers(mux)

	if err != nil {
		return err
	}

	switch opts.Provider {
	case ProtomapsProvider:

		err := protomaps.AppendAssetHandlers(mux)

		if err != nil {
			return err
		}

	case TangramJSProvider:

		err := tangramjs.AppendAssetHandlers(mux)

		if err != nil {
			return err
		}

		if opts.TilezenEnableTilepack {

			tilepack_reader, err := tilepack.NewMbtilesReader(opts.TilezenTilepackPath)

			if err != nil {
				return fmt.Errorf("Failed to create tilepack reader, %w", err)
			}

			tilepack_handler := tilepack_http.MbtilesHandler(tilepack_reader)
			mux.Handle(opts.TilezenTilepackURL, tilepack_handler)
		}

	default:
		return fmt.Errorf("Invalid or unsupporter map provider '%v'", opts.Provider)
	}

	return nil
}
