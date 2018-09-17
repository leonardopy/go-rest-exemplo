package models

import "encoding/json"

func UnmarshalErro(data []byte) (Erro, error) {
	var r Erro
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Erro) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Erro struct {
	Detail string `json:"detail"`
}
