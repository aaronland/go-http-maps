package maps

// MapConfig defines common configuration details for maps.
type MapConfig struct {
	// A valid map provider label.
	Provider string `json:"provider"`
	// A valid Leaflet tile layer URI.
	TileURL string `json:"tile_url"`
	// Optional Protomaps configuration details
	Protomaps     *ProtomapsConfig `json:"protomaps,omitempty"`
	Leaflet       *LeafletConfig   `json:"leaflet,omitempty"`
	InitialView   [2]float64       `json:"initial_view,omitempty"`
	InitialZoom   int              `json:"initial_zoom,omitempty"`
	InitialBounds [4]float64       `json:"initial_view,omitempty"`
}
