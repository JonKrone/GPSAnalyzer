package GPSAnalyzer

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/antonholmquist/jason"
)

const my_access_token = "pk.eyJ1IjoiYmFzdGlhbnRoZW9uZSIsImEiOiJjamhwanRpdGowdXFkMzdwdjNocWxnajd1In0.rDrepWFNDCNEzDEjLpe1fA"

// MatchMap takes the raw data points and matches them to the OpenStreetMap using https://www.mapbox.com.
// It returns the matched points.
// It prints the matched points in the following format:
// 	Time: Longitude, Latitude
func MatchMap(points []Point)([]Point) {
	// FIXME don't print out the result but return them and return error if necessary
	ps, ts := convert(points)
	result := make([]Point,0)
	for i,coordinates := range ps {
		response, err := http.Get("https://api.mapbox.com/matching/v5/mapbox/driving-traffic/" + coordinates + "?tidy=true&timestamps="+ts[i]+"&access_token=" + my_access_token)
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			v, _ := jason.NewObjectFromBytes(data)
			ok,_ := v.GetString("code")
			if ok != "Ok"{
				fmt.Println("Couldn't match point",i*100,"until point",(i+1)*100)
				continue
			}
			tracepoints, err := v.GetValueArray("tracepoints")
			if err != nil {
				fmt.Println(err)
			}else{
				for j,t := range tracepoints {
					// t is null if the raw data point is seen as an outlier or cluster
					if t.Null() == nil{
						continue
					}
					o,err := t.Object()
					if err != nil{
						fmt.Println(err)
						continue
					}
					loc,err := o.Map()["location"].Array()
					if err != nil{
						fmt.Println(err)
						continue
					}
					lng,err := loc[0].Float64()
					if err != nil{
						fmt.Println(err)
						continue
					}
					lat,err := loc[1].Float64()
					if err != nil{
						fmt.Println(err)
						continue
					}
					idx := i*100 + j
					fmt.Println(points[idx].Time.String(),":",lng,",",lat)
					result = append(result,Point{Lat:lat,Lng:lng,Time:points[idx].Time})
				}
			}
		}
	}
	return result
}

// convert converts the point into a format that is valid for the API.
// The Api supports at max 100 coordinates per request. It returns a string
// array of the coordinates and a string array of the unix time.
func convert(points []Point)([]string, []string){
	s := make([]string,0)
	t := make([]string,0)
	j := -1
	for i,p := range points {
		if i%100 == 0{
			j++
			s = append(s,"")
			t = append(t,"")
		}
		if i%100 != 0 {
			s[j] += ";"
			t[j] += ";"
		}
		s[j] += fmt.Sprintf("%f,%f",p.Lng, p.Lat)
		t[j] += fmt.Sprintf("%d",p.Time.Unix())
	}
	return s,t
}
