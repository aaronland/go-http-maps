package provider

import (
	"net/url"
)

type MapOptions struct {
	InitialLatitude  float64
	InitialLongitude float64
	InitialZoom      int
}

func MapOptionsFromURL(u *url.URL) (*MapOptions, error) {

	opts := &MapOptions{
		InitialLatitude:  0.0,
		InitialLongitude: 0.0,
		InitialZoom:      12,
	}

	return opts, nil
}
