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
