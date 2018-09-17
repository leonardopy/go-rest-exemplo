package models

import "encoding/json"

func UnmarshalSwapiListPlanet(data []byte) (SwapiListPlanet, error) {
	var r SwapiListPlanet
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SwapiListPlanet) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type SwapiListPlanet struct {
	Count    int64    `json:"count"`
	Next     string   `json:"next"`
	Previous string   `json:"previous"`
	Results  []Result `json:"results"`
}

type Result struct {
	Name           string   `json:"name"`
	RotationPeriod string   `json:"rotation_period"`
	OrbitalPeriod  string   `json:"orbital_period"`
	Diameter       string   `json:"diameter"`
	Climate        string   `json:"climate"`
	Gravity        string   `json:"gravity"`
	Terrain        string   `json:"terrain"`
	SurfaceWater   string   `json:"surface_water"`
	Population     string   `json:"population"`
	Residents      []string `json:"residents"`
	Films          []string `json:"films"`
	Created        string   `json:"created"`
	Edited         string   `json:"edited"`
	URL            string   `json:"url"`
}
