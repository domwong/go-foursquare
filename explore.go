package foursquare

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ExploreParams struct {
	Latitude         float64
	Longitude        float64
	Near             string
	LocAccuracy      float64
	Altitude         float64
	AltitudeAccuracy float64
	Radius           int64
	Section          string
	Query            string
	Limit            int64
	Offset           int64
	// TODO https://developer.foursquare.com/docs/venues/explore
}

type exploreRspWrapper struct {
	Meta     meta       `json:"meta"`
	Response exploreRsp `json:"response"`
}

type exploreRsp struct {
	Groups []*RecommendationGroup `json:"groups"`
	// TODO
}

type RecommendationGroup struct {
	Type            string           `json:"type"`
	Name            string           `json:"name"`
	Recommendations []Recommendation `json:"items"`
}

type Recommendation struct {
	Reasons ReasonWrapper `json:"reasons"`
	Venue   Venue         `json:"venue"`
	Tips    []Tip         `json:"tips"`
}

type ReasonWrapper struct {
	Count   int64    `json:"count"`
	Reasons []Reason `json:"items"`
}

type Reason struct {
	Summary string `json:"summary"`
	Type    string `json:"type"`
	Name    string `json:"reasonName"`
}

type Tip struct {
	Id           string `json:"id"`
	CreatedAt    int64  `json:"createdAt"`
	Text         string `json:"text"`
	Type         string `json:"type"`
	CanonicalUrl string `json:"canonicalUrl"`
}

func (s *ExploreParams) encode() string {
	params := url.Values{}
	if len(s.Near) > 0 {
		params.Add("near", s.Near)
	} else {
		params.Add("ll", fmt.Sprintf("%f,%f", s.Latitude, s.Longitude))
	}
	return params.Encode()
}

// Search invokes /venues/search
func (v *venuesClient) Explore(params *ExploreParams) ([]*RecommendationGroup, error) {
	resp, err := http.Get(fmt.Sprintf(baseurl+"explore?%s&%s", v.credsString(), params.encode()))
	if err != nil {
		return nil, fmt.Errorf("Error calling explore %v", err)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading response from explore %v", err)
	}
	var ret exploreRspWrapper
	if err := json.Unmarshal(b, &ret); err != nil {
		return nil, fmt.Errorf("Error unmarshalling response from explore %v %+v", err, string(b))
	}
	if ret.Meta.Code != 200 {
		return nil, fmt.Errorf("Error exploring %f, %s, %s", ret.Meta.Code, ret.Meta.ErrorType, ret.Meta.ErrorDetail)
	}
	return ret.Response.Groups, nil

}
