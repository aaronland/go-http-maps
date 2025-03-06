package maps

// MapConfig defines common configuration details for maps.
type MapConfig struct {
	// A valid map provider label.
	Provider string `json:"provider"`
	// A valid Leaflet tile layer URI.
	TileURL string `json:"tile_url"`
	// Optional Protomaps configuration details
	Protomaps       *ProtomapsConfig `json:"protomaps,omitempty"`
	Style           *LeafletStyle    `json:"style,omitempty"`
	PointStyle      *LeafletStyle    `json:"point_style,omitempty"`
	LabelProperties []string         `json:"label_properties"`
}
