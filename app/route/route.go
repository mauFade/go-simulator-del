package route

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

type Route struct {
	ID        string
	ClientID  string
	Positions []Position
}

type Position struct {
	Latitude  float64
	Longitude float64
}

func (route *Route) LoadPositions() error {
	if route.ID == "" {
		return errors.New("Missing route ID.")
	}

	file, err := os.Open("destinations/" + route.ID + ".txt")

	if err != nil {
		return err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		data := strings.Split(scanner.Text(), ",")

		latitude, err := strconv.ParseFloat(data[0], 64)

		if err != nil {
			return nil
		}

		longitude, err := strconv.ParseFloat(data[1], 64)

		if err != nil {
			return nil
		}

		route.Positions = append(route.Positions, Position{
			Latitude:  latitude,
			Longitude: longitude,
		})
	}

	return nil
}
