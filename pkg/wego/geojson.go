package wego

// GeoCoordinates is a latitude and longitude
type GeoCoordinates struct {
	Latitude  float64 `json:"Latitude"`
	Longitude float64 `json:"Longitude"`
}

// Location represents a place location.
type Location struct {
	Address        string          `json:"Address"`
	GeoCoordinates *GeoCoordinates `json:"Geo Coordinates"`
}

// Properties represent a WeGo place properties.
type Properties struct {
	CollectionMame string    `json:"Collection Name"`
	CollectionID   string    `json:"Collection Id"`
	WEGOURL        string    `json:"WEGO URL"`
	Location       *Location `json:"Location"`
}

// Geometry represents a WeGo place geometry (eg: a point).
type Geometry struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

// Place represents a WeGo place.
type Place struct {
	Type       string      `json:"type"`
	ID         string      `json:"id"`
	Title      string      `json:"title"`
	Properties *Properties `json:"properties"`
	Geom       *Geometry   `json:"geometry"`
	BBox       []float64   `json:"bbox"`
}

// Collection represents a geoJSON WeGo collection.
type Collection struct {
	Type   string   `json:"type"`
	Places []*Place `json:"features"`
}
