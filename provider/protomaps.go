package provider

import (
	"context"
	"fmt"
	"github.com/aaronland/go-http-leaflet"
	"github.com/sfomuseum/go-http-protomaps"
	"net/http"
	"net/url"
)

const PROTOMAPS_SCHEME string = "protomaps"

type ProtomapsProvider struct {
	Provider
	leafletOptions   *leaflet.LeafletOptions
	protomapsOptions *protomaps.ProtomapsOptions
}

func init() {
	protomaps.APPEND_LEAFLET_RESOURCES = false
	protomaps.APPEND_LEAFLET_ASSETS = false

	ctx := context.Background()
	RegisterProvider(ctx, PROTOMAPS_SCHEME, NewProtomapsProvider)
}

func ProtomapsOptionsFromURL(u *url.URL) (*protomaps.ProtomapsOptions, error) {
	opts := protomaps.DefaultProtomapsOptions()
	return opts, nil
}

func NewProtomapsProvider(ctx context.Context, uri string) (Provider, error) {

	u, err := url.Parse(uri)

	if err != nil {
		return nil, fmt.Errorf("Failed to parse URI, %w", err)
	}

	leaflet_opts, err := LeafletOptionsFromURL(u)

	if err != nil {
		return nil, fmt.Errorf("Failed to create leaflet options, %w", err)
	}

	protomaps_opts, err := ProtomapsOptionsFromURL(u)

	if err != nil {
		return nil, fmt.Errorf("Failed to create protomaps options, %w", err)
	}

	p := &ProtomapsProvider{
		leafletOptions:   leaflet_opts,
		protomapsOptions: protomaps_opts,
	}

	return p, nil
}

func (p *ProtomapsProvider) Scheme() string {
	return PROTOMAPS_SCHEME
}

func (p *ProtomapsProvider) AppendResourcesHandler(handler http.Handler) http.Handler {
	handler = leaflet.AppendResourcesHandler(handler, p.leafletOptions)
	handler = protomaps.AppendResourcesHandler(handler, p.protomapsOptions)
	return handler
}

func (p *ProtomapsProvider) AppendAssetHandlers(mux *http.ServeMux) error {

	err := leaflet.AppendAssetHandlers(mux)

	if err != nil {
		return fmt.Errorf("Failed to append leaflet asset handler, %w", err)
	}

	err = protomaps.AppendAssetHandlers(mux)

	if err != nil {
		return fmt.Errorf("Failed to append protomaps asset handler, %w", err)
	}

	return nil
}
