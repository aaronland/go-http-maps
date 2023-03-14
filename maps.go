package maps

import (
	"io"
	"log"
	gohttp "net/http"

	"github.com/aaronland/go-http-maps/provider"
	"github.com/aaronland/go-http-maps/static"
	aa_static "github.com/aaronland/go-http-static"
)

// MapsOptions provides a list of JavaScript and CSS link to include with HTML output.
type MapsOptions struct {
	JS             []string
	CSS            []string
	DataAttributes map[string]string
	// AppendJavaScriptAtEOF is a boolean flag to append JavaScript markup at the end of an HTML document
	// rather than in the <head> HTML element. Default is false
	AppendJavaScriptAtEOF bool
	RollupAssets          bool
	Prefix                string
	Logger                *log.Logger
}

// Return a *MapsOptions struct with default paths and URIs.
func DefaultMapsOptions() *MapsOptions {

	logger := log.New(io.Discard, "", 0)

	opts := &MapsOptions{
		CSS: []string{
			"/css/aaronland.maps.css",
		},
		JS: []string{
			"/javascript/aaronland.maps.js",
		},
		DataAttributes: make(map[string]string),
		Logger:         logger,
	}

	return opts
}

func AppendResourcesHandlerWithProvider(next gohttp.Handler, map_provider provider.Provider, maps_opts *MapsOptions) gohttp.Handler {
	next = map_provider.AppendResourcesHandler(next)
	next = AppendResourcesHandler(next, maps_opts)
	return next
}

// AppendResourcesHandler will rewrite any HTML produced by previous handler to include the necessary markup to load Maps JavaScript and CSS files and related assets.
func AppendResourcesHandler(next gohttp.Handler, opts *MapsOptions) gohttp.Handler {

	static_opts := aa_static.DefaultResourcesOptions()
	static_opts.CSS = opts.CSS
	static_opts.JS = opts.JS
	static_opts.AppendJavaScriptAtEOF = opts.AppendJavaScriptAtEOF

	return aa_static.AppendResourcesHandlerWithPrefix(next, static_opts, opts.Prefix)
}

// Append all the files in the net/http FS instance containing the embedded Maps assets to an *http.ServeMux instance.
func AppendAssetHandlers(mux *gohttp.ServeMux, opts *MapsOptions) error {

	return aa_static.AppendStaticAssetHandlersWithPrefix(mux, static.FS, opts.Prefix)
}
