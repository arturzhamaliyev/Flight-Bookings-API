package model

import (
	"errors"
	"strconv"
	"strings"
)

type Coordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func (c *Coordinates) Scan(value interface{}) error {
	src, ok := value.(string)
	if !ok {
		return errors.New("incompatible type")
	}

	for _, old := range []string{"\\\"", "\"", "{", "}", "(", ")"} {
		src = strings.ReplaceAll(src, old, "")
	}

	rawData := strings.Split(src, ",")

	lat, err := strconv.ParseFloat(rawData[0], 64)
	if err != nil {
		return err
	}

	long, err := strconv.ParseFloat(rawData[1], 64)
	if err != nil {
		return err
	}

	c.Latitude = lat
	c.Longitude = long

	return nil
}
