package GPSAnalyzer

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
	"time"
	"strconv"
)

func Read(path string) ([]Point, error) {
	csvFile, _ := os.Open(path)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var points []Point
	// skip first line
	_, err := reader.Read()
	if err != nil {
		return points, err
	}
	for {
		line, err := reader.Read()
		// End of file
		if err == io.EOF {
			break
		} else if err != nil {
			return points, err
		}
		points = append(points, Point{
			Time: ParseTime(line[0]),
			Lat:  ParseFloat(line[2]),
			Lng: ParseFloat(line[3]),
			Bearing: ParseFloat(line[5]),
			Speed: ParseFloat(line[6]),
		})
	}
	return points, nil
}

func ParseTime(timeString string) (time.Time) {
	layout := "2006-01-02 15:04:05 -0700"
	tme,err := time.Parse(layout,timeString)
	if err != nil{
		panic(err)
	}
	return tme
}

func ParseFloat(floatString string)(float64) {
	f,err := strconv.ParseFloat(floatString,64)
	if err != nil {
		panic(err)
	}
	return f
}
