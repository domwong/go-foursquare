package foursquare

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type SearchParams struct {
	Latitude         float64
	Longitude        float64
	Near             string
	LocAccuracy      float64
	Altitude         float64
	AltitudeAccuracy float64
	Query            string
	Limit            int64
	Intent           string
	Radius           int64
	SouthWest        string // TODO should be a lat/lng
	NorthEast        string
	CategoryId       []string
	Url              string
	ProviderId       string
	LinkedId         string
}

type searchRspWrapper struct {
	Meta     meta      `json:"meta"`
	Response searchRsp `json:"response"`
}

type searchRsp struct {
	Venues []*Venue `json:"venues"`
}

func (s *SearchParams) encode() string {
	params := url.Values{}
	if len(s.Near) > 0 {
		params.Add("near", s.Near)
	} else {
		params.Add("ll", fmt.Sprintf("%f,%f", s.Latitude, s.Longitude))
	}
	return params.Encode()
}

// Search invokes /venues/search
func (v *venuesClient) Search(params *SearchParams) ([]*Venue, error) {
	resp, err := http.Get(fmt.Sprintf(baseurl+"search?%s&%s", v.credsString(), params.encode()))
	if err != nil {
		return nil, fmt.Errorf("Error calling search %v", err)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading response from search %v", err)
	}
	var ret searchRspWrapper
	if err := json.Unmarshal(b, &ret); err != nil {
		return nil, fmt.Errorf("Error unmarshalling response from search %v %+v", err, string(b))
	}
	if ret.Meta.Code != 200 {
		return nil, fmt.Errorf("Error searching  %f, %s, %s", ret.Meta.Code, ret.Meta.ErrorType, ret.Meta.ErrorDetail)
	}
	return ret.Response.Venues, nil

}
