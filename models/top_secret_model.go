package models

type TopSecret struct {
	Name     string   `json:"name"`
	Distance float32  `json:"distance"`
	Message  []string `json:"message"`
}

type SatellitePosition struct {
	Name     string
	Position Position `json:"position"`
}

type Position struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}
