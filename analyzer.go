package GPSAnalyzer

import (
	"fmt"
	"math"
	"time"
)

// Analyze returns some key performance indicators using the information given
func Analyze(points []Point) (KPI) {
	var kpi KPI
	start := true
	stopped := true
	var stopTime time.Time
	var prevLat float64
	var prevLng float64
	var prevTime time.Time
	var distance float64
	var tme float64
	for _,p := range points {
		if !start {
			// wait for the car to start moving
			distance = distance + distanceLngLat(prevLat, prevLng, p.Lat, p.Lng)
			tme = tme + p.Time.Sub(prevTime).Minutes()
		}
		if p.Speed > kpi.maxSpeed {
			kpi.maxSpeed = p.Speed
		}
		if p.Speed == 0 {
			if !stopped {
				stopped = true
				stopTime = p.Time
			}
		} else {
			if start {
				// wait for the car to start moving
				distance = distance + distanceLngLat(prevLat, prevLng, p.Lat, p.Lng)
				tme = tme + p.Time.Sub(prevTime).Minutes()
				start = false
				stopped = false
			}
			// only add time and distance to kpi when the car is moving to avoid adding the time when the car stops for
			// the final time.
			kpi.distanceTravelled = distance
			kpi.timeTravelled = tme
			if stopped {
				stopped = false
				// if the stop was between 5 seconds and 3 minutes it was probably a red light.
				// A better way to check would be to use a map that contains traffic lights.
				if p.Time.Sub(stopTime) > 5*time.Second  && p.Time.Sub(stopTime) < 3*time.Minute {
					kpi.redLights += 1
					kpi.redLightDuration += p.Time.Sub(stopTime).Minutes()
				}else if p.Time.Sub(stopTime) > 3*time.Minute { // if stop was longer than 3 minutes
					kpi.stops += 1
					kpi.stopDuration += p.Time.Sub(stopTime).Minutes()
				}
			}
		}
		prevTime = p.Time
		prevLat = p.Lat
		prevLng = p.Lng
	}
	kpi.averageSpeed = kpi.distanceTravelled / (kpi.timeTravelled/60)
	return kpi
}

// distanceLngLat returns the distance between two coordinates in km.
//
// The distance is calculated using Haversine formula. Because the
// earth is not a perfect sphere there is a 0.5% error possible.
func distanceLngLat(lat1, lng1, lat2, lng2 float64)(float64){
	const R = 6371 // Radius of the earth
	dLat := deg2rad(lat2 - lat1)
	dLng := deg2rad(lng2 - lng1)
	a := math.Sin(dLat/2) * math.Sin(dLat/2) +
		math.Cos(deg2rad(lat1)) * math.Cos(deg2rad(lat2)) *
		math.Sin(dLng/2) * math.Sin(dLng/2)
	c := 2 * math.Atan2(math.Sqrt(a),math.Sqrt(1-a))
	return R * c
}

func deg2rad(deg float64)(float64) {
	return deg * (math.Pi/180)
}

type KPI struct {
	distanceTravelled  float64
	timeTravelled      float64
	stops              int
	stopDuration       float64
	redLights          int
	redLightDuration   float64
	averageSpeed       float64
	maxSpeed           float64
}

func (kpi *KPI) String() string {
	return fmt.Sprintf("Distance travelled: %.2f km, Time travelled: %.1f min, Nr. of Stops: %d, "+
		"Duration of Stops: %.1f min, Nr. of red lights: %d, Duration of red lights: %.1f min, "+
		"Average Speed: %.1f km/h, Max Speed: %.1f km/h",kpi.distanceTravelled,
		kpi.timeTravelled, kpi.stops, kpi.stopDuration, kpi.redLights, kpi.redLightDuration,
		kpi.averageSpeed, kpi.maxSpeed)
}
