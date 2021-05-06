package provider

import (
	"flag"
	"fmt"
	"github.com/aaronland/go-http-leaflet"
	"github.com/aaronland/go-http-tangramjs"
	"github.com/sfomuseum/go-http-protomaps"
	"net/http"
)

type Provider int

const (
	Unknown Provider = iota
	TangramJS
	Protomaps
)

type MapOptions struct {
	Provider         Provider
	LeafletOptions   *leaflet.LeafletOptions
	NextzenOptions   *tangramjs.NextzenOptions
	ProtomapsOptions *protomaps.ProtomapsOptions
}

func init() {
	tangramjs.APPEND_LEAFLET_RESOURCES = false
	tangramjs.APPEND_LEAFLET_ASSETS = false
	protomaps.APPEND_LEAFLET_RESOURCES = false
	protomaps.APPEND_LEAFLET_ASSETS = false
}

func MapOptionsFromFlagSet(fs *flag.FlagSet) (*MapOptions, error) {

	opts := &MapOptions{}

	return opts, nil
}

func AppendResourcesHandler(handler http.Handler, opts *MapOptions) http.Handler {

	handler = leaflet.AppendResourcesHandler(handler, opts.LeafletOptions)

	switch opts.Provider {
	case Protomaps:

		handler = protomaps.AppendResourcesHandler(handler, opts.ProtomapsOptions)

	case TangramJS:

		tangramjs_opts := tangramjs.DefaultTangramJSOptions()
		handler = tangramjs.AppendResourcesHandler(handler, tangramjs_opts)

	default:
		// pass
	}

	return handler
}

func AppendAssetHandlers(mux *http.ServeMux, opts *MapOptions) error {

	switch opts.Provider {
	case Protomaps:

		err := protomaps.AppendAssetHandlers(mux)

		if err != nil {
			return err
		}

	case TangramJS:

		err := tangramjs.AppendAssetHandlers(mux)

		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("Invalid or unsupporter map provider '%v'", opts.Provider)
	}

	return nil
}
