package services

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"
	"strconv"
)

type Coordinates struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

func GetCoord(street string, district string) (*Coordinates, error) {
	baseUrl := "https://nominatim.openstreetmap.org/search"
	encodedQuery := url.QueryEscape(street + ", " + district)

	response, err := http.Get(baseUrl + "?q=" + encodedQuery + "&format=json")
	if err != nil {
		return nil, err

	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(response.Body)

	type results []struct {
		Lat string `json:"lat"`
		Lon string `json:"lon"`
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var result results
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	if len(result) > 0 {
		return &Coordinates{
			Latitude:  result[0].Lat,
			Longitude: result[0].Lon,
		}, nil
	}

	return nil, fmt.Errorf("no coordinates found")
}

type Distance struct {
	Distance   float64 `json:"distance"`
	Latitude1  string  `json:"latitude1"`
	Longitude1 string  `json:"longitude1"`
	Latitude2  string  `json:"latitude2"`
	Longitude2 string  `json:"longitude2"`
}

func toRadians(degrees string) (float64, error) {
	degreesFloat, err := strconv.ParseFloat(degrees, 64)
	if err != nil {
		return 0, err
	}
	return degreesFloat * (math.Pi / 180), nil
}

func GetDistance(lat1 string, lon1 string, lat2 string, lon2 string) (*Distance, error) {
	const EarthRadius = 6371

	lat1Rad, err := toRadians(lat1)
	if err != nil {
		return nil, err
	}
	lon1Rad, err := toRadians(lon1)
	if err != nil {
		return nil, err
	}
	lat2Rad, err := toRadians(lat2)
	if err != nil {
		return nil, err
	}
	lon2Rad, err := toRadians(lon2)
	if err != nil {
		return nil, err
	}

	deltaLat := lat2Rad - lat1Rad
	deltaLon := lon2Rad - lon1Rad

	a := math.Sin(deltaLat/2)*math.Sin(deltaLat/2) +
		math.Cos(lat1Rad)*math.Cos(lat2Rad)*math.Sin(deltaLon/2)*math.Sin(deltaLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	distance := EarthRadius * c

	return &Distance{
		Distance:   distance,
		Latitude1:  lat1,
		Longitude1: lon1,
		Latitude2:  lat2,
		Longitude2: lon2,
	}, nil
}
