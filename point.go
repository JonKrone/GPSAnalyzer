package GPSAnalyzer

import (
	"fmt"
	"time"
)

type Point struct {
	Time time.Time
	Lat float64
	Lng float64
	Bearing float64
	Speed float64
}

func (p *Point) String() string {
	return fmt.Sprintf("Time: %s, Latitude: %f, Longitute: %f, Bearing: %f, Speed: %f",
		p.Time,p.Lat,p.Lng,p.Bearing,p.Speed)
}
