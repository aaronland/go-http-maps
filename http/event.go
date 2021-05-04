package http

import (
	"errors"
	"html/template"
	gohttp "net/http"
)

type EventHandlerOptions struct {
	Templates        *template.Template
	InitialLatitude  float64
	InitialLongitude float64
	InitialZoom      int
}

type EventHandlerVars struct {
	InitialLatitude  float64
	InitialLongitude float64
	InitialZoom      int
}

func EventHandler(opts *EventHandlerOptions) (gohttp.Handler, error) {

	t := opts.Templates.Lookup("event")

	if t == nil {
		return nil, errors.New("Missing 'event' template")
	}

	fn := func(rsp gohttp.ResponseWriter, req *gohttp.Request) {

		vars := EventHandlerVars{
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
