package helper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"go.uber.org/zap"

	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/model"
)

type response struct {
	Places []struct {
		FormattedAddress  string `json:"formattedAddress"`
		AddressComponents []struct {
			LongText     string   `json:"longText"`
			ShortText    string   `json:"shortText"`
			Types        []string `json:"types"`
			LanguageCode string   `json:"languageCode"`
		} `json:"addressComponents"`
		Location model.Coordinates `json:"location"`
		Name     struct {
			Text         string `json:"text"`
			LanguageCode string `json:"languageCode"`
		} `json:"displayName"`
	} `json:"places"`
}

const url = `https://places.googleapis.com/v1/places:searchNearby`

func FindAirportByCoordinates(coordinates model.Coordinates) (model.Airport, error) {
	body := []byte(
		fmt.Sprintf(`
	{
		"includedTypes": ["airport"],
		"maxResultCount": 1,
		"locationRestriction": {
		  "circle": {
			"center": {
			  "latitude": %f,
			  "longitude": %f
			},
			"radius": 5000.0
		  }
		}
	}
	`,
			coordinates.Latitude,
			coordinates.Longitude,
		),
	)

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		zap.S().Info(err)
		return model.Airport{}, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Goog-Api-Key", "AIzaSyBxx7ig2mlKVAhB8UTCKtvbYbilJLXmqKQ")
	req.Header.Add("X-Goog-FieldMask", "places.displayName,places.formattedAddress,places.location,places.addressComponents")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		zap.S().Info(err)
		return model.Airport{}, err
	}
	defer resp.Body.Close()

	var r response
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		zap.S().Info(err)
		return model.Airport{}, err
	}

	var city, country string
	for _, component := range r.Places[0].AddressComponents {
		if component.Types[0] == "administrative_area_level_1" {
			city = component.LongText
		} else if component.Types[0] == "country" {
			country = component.LongText
		}
	}

	// fmt.Println(r)
	return model.Airport{
		Name:        r.Places[0].Name.Text,
		City:        city,
		Country:     country,
		Coordinates: r.Places[0].Location,
	}, nil
}
