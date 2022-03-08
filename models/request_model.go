package models

type Request struct {
	Satellites []TopSecret `json:"satellites"`
}

type RequestSplit struct {
	Name     string   `uri:"satellite_name"`
	Distance float32  `json:"distance"`
	Message  []string `json:"message"`
}
