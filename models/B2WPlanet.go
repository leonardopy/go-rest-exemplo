package models

import "encoding/json"

func UnmarshalB2WPlanet(data []byte) (B2WPlanet, error) {
	var r B2WPlanet
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *B2WPlanet) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type B2WPlanet struct {
	Nome      string `json:"ID"`
	Clima     string `json:"clima"`
	Terreno   string `json:"terreno"`
	QtdFilmes int    `json:"qtdFilmes"`
}

type B2WPlanetPost struct {
	Nome    string `json:"ID"`
	Clima   string `json:"clima"`
	Terreno string `json:"terreno"`
}
