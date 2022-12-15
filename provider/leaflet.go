package provider

import (
	"github.com/aaronland/go-http-leaflet"
	"net/url"
)

func LeafletOptionsFromURL(u *url.URL) (*leaflet.LeafletOptions, error) {

	opts := leaflet.DefaultLeafletOptions()
	return opts, nil
}
