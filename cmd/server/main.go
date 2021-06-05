package main

import (
	"context"
	"github.com/aaronland/go-http-bootstrap"
	"github.com/aaronland/go-http-maps/http"
	"github.com/aaronland/go-http-maps/provider"
	"github.com/aaronland/go-http-maps/templates"
	"github.com/aaronland/go-http-server"
	"github.com/sfomuseum/go-flags/flagset"
	"log"
	gohttp "net/http"
)

func main() {

	fs := flagset.NewFlagSet("map")

	server_uri := fs.String("server-uri", "http://localhost:8080", "A valid aaronland/go-http-server URI")

	err := provider.AppendProviderFlags(fs)

	if err != nil {
		log.Fatalf("Failed to append map provider flags, %v", err)
	}

	flagset.Parse(fs)

	ctx := context.Background()

	s, err := server.NewServer(ctx, *server_uri)

	if err != nil {
		log.Fatalf("Failed to create new server for '%s', %v", *server_uri, err)
	}

	t, err := templates.LoadHTMLTemplates()

	if err != nil {
		log.Fatalf("Failed to parse templates, %v", err)
	}

	provider_opts, err := provider.ProviderOptionsFromFlagSet(fs)

	if err != nil {
		log.Fatalf("Failed to create map options from flagset, %v", err)
	}

	mux := gohttp.NewServeMux()

	err = bootstrap.AppendAssetHandlers(mux)

	if err != nil {
		log.Fatalf("Failed to append Bootstrap asset handlers, %v")
	}

	err = http.AppendStaticAssetHandlers(mux)

	if err != nil {
		log.Fatalf("Failed to append static asset handlers, %v")
	}

	err = provider.AppendAssetHandlers(mux, provider_opts)

	if err != nil {
		log.Fatalf("Failed to append Leaflet asset handlers, %v", err)
	}

	map_opts := &http.MapHandlerOptions{
		Templates:        t,
		MapProvider:      provider_opts.Provider,
		InitialLatitude:  provider_opts.InitialLatitude,
		InitialLongitude: provider_opts.InitialLongitude,
		InitialZoom:      provider_opts.InitialZoom,
	}

	map_handler, err := http.MapHandler(map_opts)

	if err != nil {
		log.Fatalf("Failed to create map handler, %v", err)
	}

	bootstrap_opts := bootstrap.DefaultBootstrapOptions()
	map_handler = bootstrap.AppendResourcesHandler(map_handler, bootstrap_opts)

	map_handler = provider.AppendResourcesHandler(map_handler, provider_opts)
	mux.Handle("/", map_handler)

	//

	log.Printf("Listening on %s\n", s.Address())
	err = s.ListenAndServe(ctx, mux)

	if err != nil {
		log.Fatalf("Failed to start server, %v", err)
	}
}
