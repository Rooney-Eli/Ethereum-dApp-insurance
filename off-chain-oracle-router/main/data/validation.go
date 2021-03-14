package data

import (
	"encoding/json"
	"io"
)

type Validation struct {
	Validation bool `json:"valid"`
}


func (v *Validation) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(v)
}

func (v *Validation) ToJSON(r io.Writer) error {
	e := json.NewEncoder(r)
	return e.Encode(v)
}



