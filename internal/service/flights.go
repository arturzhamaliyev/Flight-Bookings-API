package service

type FlightsService struct{}

func NewFlightsService() *FlightsService {
	return &FlightsService{}
}

// type response struct {
// 	Places []struct {
// 		FormattedAddress string            `json:"formattedAddress"`
// 		Location         model.Coordinates `json:"location"`
// 		Name             struct {
// 			Text         string `json:"text"`
// 			LanguageCode string `json:"languageCode"`
// 		} `json:"displayName"`
// 	} `json:"places"`
// }

// const url = `https://places.googleapis.com/v1/places:searchNearby`

// func (f *FlightsService) FindAirportByCoordinates(coordinates model.Coordinates) (model.Airport, error) {
// 	body := []byte(`
// 	{
// 		"includedTypes": ["airport"],
// 		"maxResultCount": 1,
// 		"locationRestriction": {
// 		  "circle": {
// 			"center": {
// 			  "latitude": 51.15491135401816,
// 			  "longitude": -51.537861551602454
// 			},
// 			"radius": 500.0
// 		  }
// 		}
// 	}
// 	`)

// 	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
// 	if err != nil {
// 		zap.S().Info(err)
// 		return model.Airport{}, err
// 	}

// 	req.Header.Add("Content-Type", "application/json")
// 	req.Header.Add("X-Goog-Api-Key", "AIzaSyBxx7ig2mlKVAhB8UTCKtvbYbilJLXmqKQ")
// 	req.Header.Add("X-Goog-FieldMask", "places.displayName,places.formattedAddress,places.location")

// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		zap.S().Info(err)
// 		return model.Airport{}, err
// 	}
// 	defer resp.Body.Close()

// 	var places response
// 	err = json.NewDecoder(resp.Body).Decode(&places)
// 	if err != nil {
// 		zap.S().Info(err)
// 		return model.Airport{}, err
// 	}

// 	fmt.Println(places)
// 	return model.Airport{}, nil
// }
