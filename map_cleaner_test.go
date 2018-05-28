package GPSAnalyzer

import (
	"fmt"
	"testing"
)

func TestMatchMap(t *testing.T) {
	// Todo: check results
	fmt.Println("Route 1")
	points, err := Read("testdata\\Route-1.csv")
	if err != nil {
		t.Fatal(err)
	}
	MatchMap(points)

	fmt.Println("Route 2")
	points2, err := Read("testdata\\Route-2.csv")
	if err != nil {
		t.Fatal(err)
	}
	MatchMap(points2)
}
