package toolkit

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

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

func (t *Tools) WriteJSON(w http.ResponseWriter, status int,
	data interface{}, headers ...http.Header) error {
	out, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_, err = w.Write(out)
	if err != nil {
		return err
	}

	return nil
}

func (t *Tools) ErrorJSON(w http.ResponseWriter,
	err error, status ...int) error {
	statusCode := http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]
	}

	var payload JSONResponse
	payload.Error = true
	payload.Message = err.Error()

	return t.WriteJSON(w, statusCode, payload)
}
