package foursquare

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type venueRspWrapper struct {
	Meta     meta     `json:"meta"`
	Response venueRsp `json:"response"`
}

type venueRsp struct {
	Venue Venue `json:"venue"`
}

// Search invokes /venues/search
func (v *venuesClient) Venue(id string) (*Venue, error) {
	resp, err := http.Get(fmt.Sprintf(baseurl+"%s?%s", id, v.credsString()))
	if err != nil {
		return nil, fmt.Errorf("Error calling venue %v", err)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading response from venue %v", err)
	}
	var ret venueRspWrapper
	if err := json.Unmarshal(b, &ret); err != nil {
		return nil, fmt.Errorf("Error unmarshalling response from venue %v %+v", err, string(b))
	}
	if ret.Meta.Code != 200 {
		return nil, fmt.Errorf("Error getting venue %f, %s, %s", ret.Meta.Code, ret.Meta.ErrorType, ret.Meta.ErrorDetail)
	}
	return &ret.Response.Venue, nil

}
