package foursquare

import (
	"fmt"
)

const (
	baseurl = "https://api.foursquare.com/v2/venues/"
)

type VenuesClient interface {
	Search(*SearchParams) ([]*Venue, error)
	Explore(*ExploreParams) ([]*RecommendationGroup, error)
	Venue(string) (*Venue, error)
}

type venuesClient struct {
	id      string // Foursquare provided
	secret  string // Foursquare provided
	version string
}

func NewVenuesClient(id, secret, version string) VenuesClient {
	if len(version) == 0 {
		version = "20150702"
	}
	return &venuesClient{
		id:      id,
		secret:  secret,
		version: version,
	}
}

func (v *venuesClient) credsString() string {
	return fmt.Sprintf("client_id=%s&client_secret=%s&v=%s", v.id, v.secret, v.version)
}

type meta struct {
	Code        int64  `json:"code"`
	ErrorType   string `json:"error_type"`
	ErrorDetail string `json:"error_detail"`
}

// TODO add omitempty where appropriate
type Venue struct {
	Id         string     `json:"id"`
	Name       string     `json:"name"`
	Contact    Contact    `json:"contact"`
	Location   Location   `json:"location"`
	Categories []Category `json:"categories"`
	Verified   bool       `json:"verified"`
	Stats      Stats      `json:"stats"`
	Url        string     `json:"url"`
	Hours      Hours      `json:"hours"`
	Popular    Hours      `json:"popular"`
	Menu       Menu       `json:"menu"`
	Price      Price      `json:"price"`
	Rating     float64    `json:"rating"`
	// TODO https://developer.foursquare.com/docs/responses/venue
}

type Contact struct {
	Twitter        string `json:"twitter"`
	Phone          string `json:"phone"`
	FormattedPhone string `json:"formattedPhone"`
}

type Location struct {
	Address     string  `json:"address"`
	CrossStreet string  `json:"crossStreet"`
	City        string  `json:"city"`
	State       string  `json:"state"`
	PostalCode  string  `json:"postalCode"`
	Country     string  `json:"country"`
	Latitude    float64 `json:"lat"`
	Longitude   float64 `json:"lng"`
	Distance    float64 `json:"distance"`
	IsFuzzed    bool    `json:"isFuzzed"`
}

type Category struct {
	Id         string     `json:"id"`
	Name       string     `json:"name"`
	PluralName string     `json:"pluralName"`
	ShortName  string     `json:"shortName"`
	Icon       Icon       `json:"icon"`
	Categories []Category `json:"categories"`
}

type Icon struct {
	Prefix string `json:"prefix"`
	Suffix string `json:"suffix"`
}

type Stats struct {
	Checkins int64 `json:"checkinsCount"`
	Users    int64 `json:"usersCount"`
	Tips     int64 `json:"tipCount"`
}

type Hours struct {
	Status string `json:"status"`
	IsOpen bool   `json:"isOpen"`
	// TODO https://developer.foursquare.com/docs/responses/hoursformatted.html
}

type Menu struct {
	Url       string `json:"url"`
	MobileUrl string `json:"hours"`
}

type Price struct {
	Tier        int64  `json:"tier"`
	Description string `json:"message"`
}
