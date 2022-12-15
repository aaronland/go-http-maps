package server

import (
	"context"
	"flag"
	"fmt"
	"github.com/aaronland/go-http-bootstrap"
	"github.com/aaronland/go-http-maps/http/www"
	"github.com/aaronland/go-http-maps/provider"
	"github.com/aaronland/go-http-maps/templates/html"
	"github.com/aaronland/go-http-server"
	"github.com/sfomuseum/go-flags/flagset"
	"log"
	"net/http"
)

func Run(ctx context.Context, logger *log.Logger) error {

	fs, err := DefaultFlagSet()

	if err != nil {
		return fmt.Errorf("Failed to create default flag set, %w", err)
	}

	return RunWithFlagSet(ctx, fs, logger)
}

func RunWithFlagSet(ctx context.Context, fs *flag.FlagSet, logger *log.Logger) error {

	flagset.Parse(fs)

	t, err := html.LoadTemplates(ctx)

	if err != nil {
		return fmt.Errorf("Failed to parse templates, %v", err)
	}

	provider_uri, err := provider.ProviderURIFromFlagSet(fs)

	if err != nil {
		return fmt.Errorf("Failed to derive provider URI from flagset, %v", err)
	}

	pr, err := provider.NewProvider(ctx, provider_uri)

	if err != nil {
		return fmt.Errorf("Failed to create new provider, %w", err)
	}

	mux := http.NewServeMux()

	err = bootstrap.AppendAssetHandlers(mux)

	if err != nil {
		return fmt.Errorf("Failed to append Bootstrap asset handlers, %v")
	}

	err = www.AppendStaticAssetHandlers(mux)

	if err != nil {
		return fmt.Errorf("Failed to append static asset handlers, %v")
	}

	err = pr.AppendAssetHandlers(mux)

	if err != nil {
		return fmt.Errorf("Failed to append provider asset handlers, %v", err)
	}

	map_opts := &www.MapHandlerOptions{
		Templates:        t,
		MapProvider:      pr,
		InitialLatitude:  initial_latitude,
		InitialLongitude: initial_longitude,
		InitialZoom:      initial_zoom,
	}

	map_handler, err := www.MapHandler(map_opts)

	if err != nil {
		return fmt.Errorf("Failed to create map handler, %v", err)
	}

	bootstrap_opts := bootstrap.DefaultBootstrapOptions()
	map_handler = bootstrap.AppendResourcesHandler(map_handler, bootstrap_opts)

	map_handler = pr.AppendResourcesHandler(map_handler)
	mux.Handle("/", map_handler)

	//

	s, err := server.NewServer(ctx, server_uri)

	if err != nil {
		return fmt.Errorf("Failed to create new server for '%s', %w", server_uri, err)
	}

	log.Printf("Listening on %s\n", s.Address())

	err = s.ListenAndServe(ctx, mux)

	if err != nil {
		return fmt.Errorf("Failed to start server, %w", err)
	}

	return nil
}