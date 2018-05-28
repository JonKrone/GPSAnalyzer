package GPSAnalyzer

import (
	"testing"
)

func TestRead(t *testing.T) {
	path := "testdata\\Route-test.csv"
	got, err := Read(path)
	if err != nil {
		t.Fatal(err)
	}
	want := []Point{{ParseTime("2018-05-23 00:00:07 +0200"),48.6525803,8.8251185,0.0,0.0},
		{ParseTime("2018-05-23 00:00:17 +0200"),48.6525802,8.8251185,0.0,0.0}}
	if len(want) != len(got) {
		t.Fatalf("got = %d points, want = %d", len(got), len(want))
	}
	for i := range got {
		if !got[i].Time.Equal(want[i].Time){
			t.Errorf("Time: got[%d] = %s, want[%d] = %s", i, got[i].Time, i, want[i].Time)
		}
		if got[i].Lat != want[i].Lat {
			t.Errorf("Latitude: got[%d] = %f, want[%d] = %f", i, got[i].Lat, i, want[i].Lat)
		}
		if got[i].Lng != want[i].Lng {
			t.Errorf("Longitude: got[%d] = %f, want[%d] = %f", i, got[i].Lng, i, want[i].Lng)
		}
		if got[i].Bearing != want[i].Bearing {
			t.Errorf("got[%d] = %f, want[%d] = %f", i, got[i].Bearing, i, want[i].Bearing)
		}
		if got[i].Speed != want[i].Speed {
			t.Errorf("got[%d] = %f, want[%d] = %f", i, got[i].Speed, i, want[i].Speed)
		}
	}
}