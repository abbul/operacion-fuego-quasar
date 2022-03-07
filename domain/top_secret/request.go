package top_secret

type Satellite struct {
	Name     string   `json:"name"`
	Distance float32  `json:"distance"`
	Message  []string `json:"message"`
}

type Request struct {
	Satellites []Satellite `json:"satellites"`
}
