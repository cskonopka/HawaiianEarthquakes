// https://earthquake.usgs.gov/fdsnws/event/1/query?format=geojson&starttime=2014-01-01&endtime=2014-01-02

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type featureCollection struct {
	Features []feature `json:"features"`
}

type feature struct {
	Properties Earthquake `json:"properties"`
}

type Earthquake struct {
	Title string `json:"title"`
}

func main() {
	current := time.Now()
	currentFormat := current.Format("2006-01-02")
	yesterdayTime := time.Now().Add(-24 * time.Hour)
	yesterFormat := yesterdayTime.Format(("2006-01-02"))
	findHawaiianVolcanos := "https://earthquake.usgs.gov/fdsnws/event/1/query?format=geojson&starttime=" + yesterFormat + "&endtime=" + currentFormat

	resp, err := http.Get(findHawaiianVolcanos)
	if err != nil {
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var record featureCollection
	json.Unmarshal(body, &record)

	quakes := make([]Earthquake, 0)
	for _, f := range record.Features {
		quakes = append(quakes, f.Properties)
	}

	responseLength := len(quakes)

	for q := 0; q < responseLength; q++ {
		if strings.Contains(quakes[q].Title, "Hawaii") != false {
			fmt.Println(quakes[q])
		}
	}
}
