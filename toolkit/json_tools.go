package toolkit

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// JSONStruct is the type used for sending JSON around
type JSONStruct struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func (t *Tools) ReadJSON(w http.ResponseWriter, r *http.Request,
	data interface{}) error {
	maxBytes := 1024 * 1024
	if t.MaxJSONSize != 0 {
		maxBytes = t.MaxJSONSize
	}

	r.Body = http.MaxBytesReader(w, r.Body,
		int64(maxBytes))

	dec := json.NewDecoder(r.Body)

	if !t.AllowUnknownFields {
		dec.DisallowUnknownFields()
	}

	err := dec.Decode(data)
	if err != nil {
		return err
	}

	err = dec.Decode(&struct {
	}{})
	if err != io.EOF {
		return errors.New("Body must contains only one JSON value")

	}

	return nil
}
