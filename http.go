package maps

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

// MapConfigHandler returns a new `http.Handler` that will return a JSON-encoded version of 'cfg'.
func MapConfigHandler(cfg *MapConfig) http.Handler {

	fn := func(rsp http.ResponseWriter, req *http.Request) {

		rsp.Header().Set("Content-type", "application/json")

		enc := json.NewEncoder(rsp)
		err := enc.Encode(cfg)

		if err != nil {
			slog.Error("Failed to encode map config", "error", err)
			http.Error(rsp, "Internal server error", http.StatusInternalServerError)
		}

		return
	}

	return http.HandlerFunc(fn)
}
