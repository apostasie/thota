package away

// A List represents a collection of places.
type List struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Places      []*Place `json:"places"`
}

// A Place represents a location with its details.
type Place struct {
	Name      string  `json:"name"`
	Address   string  `json:"address"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Note      string  `json:"notes,omitempty"`
}
