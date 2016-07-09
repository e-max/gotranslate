package dictionary

import "encoding/json"

type Word struct {
	Title  string `json:"title,omitempty"`
	Locale string `json:"locale,omitempty"`
}

// Create json from word
func (w *Word) Marshal() ([]byte, error) {
	return json.Marshal(w)
}
