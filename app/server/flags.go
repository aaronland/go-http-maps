package server

import (
	"flag"
	"fmt"
	"github.com/aaronland/go-http-maps/provider"
	"github.com/sfomuseum/go-flags/flagset"
)

var server_uri string
var initial_latitude float64
var initial_longitude float64
var initial_zoom int

func DefaultFlagSet() (*flag.FlagSet, error) {

	fs := flagset.NewFlagSet("map")

	fs.StringVar(&server_uri, "server-uri", "http://localhost:8080", "A valid aaronland/go-http-server URI")

	fs.Float64Var(&initial_latitude, "initial-latitude", 0.0, "...")
	fs.Float64Var(&initial_longitude, "initial-longitude", 0.0, "...")
	fs.IntVar(&initial_zoom, "initial-zoom", 12, "...")

	err := provider.AppendProviderFlags(fs)

	if err != nil {
		return nil, fmt.Errorf("Failed to append map provider flags, %v", err)
	}

	return fs, nil
}