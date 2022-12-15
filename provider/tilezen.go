package provider

import (
	"net/url"
)

type TilezenOptions struct {
	EnableTilepack bool
	TilepackPath   string
	TilepackURL    string
}

func TilezenOptionsFromURL(u *url.URL) (*TilezenOptions, error) {

	opts := &TilezenOptions{
		EnableTilepack: false,
		TilepackPath:   "",
		TilepackURL:    "",
	}

	return opts, nil
}
