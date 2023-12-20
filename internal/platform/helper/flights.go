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

func SendRequestToGetExactCoordinatesOfAirport(coordinates model.Coordinates, apiKey, url string) (response, error) {
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
		return response{}, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Goog-Api-Key", apiKey)
	req.Header.Add("X-Goog-FieldMask", "places.displayName,places.formattedAddress,places.location,places.addressComponents")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		zap.S().Info(err)
		return response{}, err
	}
	defer resp.Body.Close()

	var r response
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		zap.S().Info(err)
		return response{}, err
	}

	return r, nil
}
