package dashboard

import (
	"encoding/json"

	"github.com/bitly/go-simplejson"
)

func ValidateDashJSON(dash *simplejson.Json) (*Dashboard, error) {
	d := Dashboard{}
	dashBytes, err := dash.MarshalJSON()
	if err != nil {
		return nil, err
	}
	json.Unmarshal(dashBytes, &d)
	return &d, d.Validate()
}

func ValidateDash(dash *Dashboard) (*Dashboard, error) {
	return dash, dash.Validate()
}

func ValidateBytesDash(dash []byte) (*Dashboard, error) {
	d := Dashboard{}
	json.Unmarshal(dash, &d)
	return &d, d.Validate()
}
