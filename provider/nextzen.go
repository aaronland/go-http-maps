package provider

import (
	"context"
	"fmt"
	"github.com/aaronland/go-http-leaflet"
	"github.com/aaronland/go-http-tangramjs"
	tilepack_http "github.com/tilezen/go-tilepacks/http"
	"github.com/tilezen/go-tilepacks/tilepack"
	"net/http"
	"net/url"
)

const NEXTZEN_SCHEME string = "nextzen"

type NextzenProvider struct {
	Provider
	leafletOptions *leaflet.LeafletOptions
	tangramOptions *tangramjs.TangramJSOptions
	mapOptions     *MapOptions
	tilezenOptions *TilezenOptions
}

func init() {
	tangramjs.APPEND_LEAFLET_RESOURCES = false
	tangramjs.APPEND_LEAFLET_ASSETS = false

	ctx := context.Background()
	RegisterProvider(ctx, NEXTZEN_SCHEME, NewNextzenProvider)
}

func TangramJSOptionsFromURL(u *url.URL) (*tangramjs.TangramJSOptions, error) {
	opts := tangramjs.DefaultTangramJSOptions()
	return opts, nil
}

func NewNextzenProvider(ctx context.Context, uri string) (Provider, error) {

	u, err := url.Parse(uri)

	if err != nil {
		return nil, fmt.Errorf("Failed to parse URI, %w", err)
	}

	leaflet_opts, err := LeafletOptionsFromURL(u)

	if err != nil {
		return nil, fmt.Errorf("Failed to create leaflet options, %w", err)
	}

	tangram_opts, err := TangramJSOptionsFromURL(u)

	if err != nil {
		return nil, fmt.Errorf("Failed to create tilezen options, %w", err)
	}

	tilezen_opts, err := TilezenOptionsFromURL(u)

	if err != nil {
		return nil, fmt.Errorf("Failed to create tilezen options, %w", err)
	}

	map_opts, err := MapOptionsFromURL(u)

	if err != nil {
		return nil, fmt.Errorf("Failed to create map options, %w", err)
	}

	p := &NextzenProvider{
		leafletOptions: leaflet_opts,
		tangramOptions: tangram_opts,
		tilezenOptions: tilezen_opts,
		mapOptions:     map_opts,
	}

	return p, nil
}

func (p *NextzenProvider) Scheme() string {
	return NEXTZEN_SCHEME
}

func (p *NextzenProvider) AppendResourcesHandler(handler http.Handler) http.Handler {
	handler = leaflet.AppendResourcesHandler(handler, p.leafletOptions)
	handler = tangramjs.AppendResourcesHandler(handler, p.tangramOptions)
	return handler
}

func (p *NextzenProvider) AppendAssetHandlers(mux *http.ServeMux) error {

	err := leaflet.AppendAssetHandlers(mux)

	if err != nil {
		return fmt.Errorf("Failed to append leaflet asset handler, %w", err)
	}

	err = tangramjs.AppendAssetHandlers(mux)

	if err != nil {
		return fmt.Errorf("Failed to append tangram asset handler, %w", err)
	}

	if p.tilezenOptions.EnableTilepack {

		tilepack_reader, err := tilepack.NewMbtilesReader(p.tilezenOptions.TilepackPath)

		if err != nil {
			return fmt.Errorf("Failed to create tilepack reader, %w", err)
		}

		tilepack_handler := tilepack_http.MbtilesHandler(tilepack_reader)
		mux.Handle(p.tilezenOptions.TilepackURL, tilepack_handler)
	}

	return nil
}
