package http

import (
	"errors"
	"html/template"
	gohttp "net/http"
)

type MapHandlerOptions struct {
	Templates        *template.Template
	InitialLatitude  float64
	InitialLongitude float64
	InitialZoom      int
	MapRenderer string
}

type MapHandlerVars struct {
	InitialLatitude  float64
	InitialLongitude float64
	InitialZoom      int
	MapRenderer string	
}

func MapHandler(opts *MapHandlerOptions) (gohttp.Handler, error) {

	t := opts.Templates.Lookup("map")

	if t == nil {
		return nil, errors.New("Missing 'map' template")
	}

	fn := func(rsp gohttp.ResponseWriter, req *gohttp.Request) {

		vars := MapHandlerVars{
			MapRenderer: opts.MapRenderer,
			InitialLatitude:  opts.InitialLatitude,
			InitialLongitude: opts.InitialLongitude,
			InitialZoom:      opts.InitialZoom,
		}

		rsp.Header().Set("Content-Type", "text/html; charset=utf-8")

		err := t.Execute(rsp, vars)

		if err != nil {
			gohttp.Error(rsp, err.Error(), gohttp.StatusInternalServerError)
			return
		}

		return
	}

	h := gohttp.HandlerFunc(fn)
	return h, nil
}
