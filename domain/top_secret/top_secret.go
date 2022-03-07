package top_secret

type TopSecret struct {
	Name     string   `json:"name"`
	Distance float32  `json:"distance"`
	Message  []string `json:"message"`
}

type SatellitePosition struct {
	Name     string
	Position Position `json:"position"`
}

type Request struct {
	Satellites []TopSecret `json:"satellites"`
}

type RequestSplit struct {
	Name     string   `uri:"satellite_name"`
	Distance float32  `json:"distance"`
	Message  []string `json:"message"`
}

type Position struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}

type Response struct {
	Position Position `json:"position"`
	Message  string   `json:"message"`
}
