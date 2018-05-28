package GPSAnalyzer

import (
	"fmt"
	"testing"
)

func TestAnalyze(t *testing.T) {
	// Todo: check results
	fmt.Println("Route 1")
	points, err := Read("testdata\\Route-1.csv")
	if err != nil {
		t.Fatal(err)
	}
	kpi := Analyze(points)
	fmt.Println(kpi.String())

	fmt.Println("Route 2")
	points2, err := Read("testdata\\Route-2.csv")
	if err != nil {
		t.Fatal(err)
	}
	kpi2 := Analyze(points2)
	fmt.Println(kpi2.String())
}
