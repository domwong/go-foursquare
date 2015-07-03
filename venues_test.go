package foursquare

import (
	"os"
	"testing"
)

func TestSearch(t *testing.T) {
	vc := NewVenuesClient(os.Getenv("FOURSQ_ID"), os.Getenv("FOURSQ_SECRET"), "")
	venues, err := vc.Search(&SearchParams{
		Near: "Trafalgar Square",
	})
	if err != nil {
		t.Errorf("Error calling search, %v", err)
		t.FailNow()
	}
	if len(venues) == 0 {
		t.Error("Failed to find any venues")
		t.FailNow()
	}
}

func TestExplore(t *testing.T) {
	vc := NewVenuesClient(os.Getenv("FOURSQ_ID"), os.Getenv("FOURSQ_SECRET"), "")
	venues, err := vc.Explore(&ExploreParams{
		//Near: "Trafalgar Square",
		Latitude:  51.522472,
		Longitude: -0.042689,
	})
	if err != nil {
		t.Errorf("Error calling explore, %v", err)
		t.FailNow()
	}
	if len(venues) == 0 {
		t.Error("Failed to find any recommendations")
		t.FailNow()
	}
}

func TestVenue(t *testing.T) {
	vc := NewVenuesClient(os.Getenv("FOURSQ_ID"), os.Getenv("FOURSQ_SECRET"), "")
	v, err := vc.Venue("4ac518cdf964a520e6a520e3")
	if err != nil {
		t.Errorf("Error calling venue, %v", err)
		t.FailNow()
	}
	if v.Name != "National Gallery" {
		t.Error("Failed to get correct venue, expected \"National Gallery\" but got \"%s\"", v.Name)
		t.FailNow()
	}
}
