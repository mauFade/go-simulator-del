package route

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Route struct {
	ID        string     `json:"routeId"`
	ClientID  string     `json:"clientId"`
	Positions []Position `json:"position"`
}

type Position struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type PartialRoutePosition struct {
	ID       string    `json:"routeId"`
	ClientId string    `json:"clientId"`
	Position []float64 `json:"position"`
	Finished bool      `json:"finished"`
}

func NewRoute() *Route {
	return &Route{}
}

func (route *Route) LoadPositions() error {
	if route.ID == "" {
		return errors.New("Missing route ID.")
	}

	file, err := os.Open("app/destinations/" + route.ID + ".txt")

	if err != nil {
		fmt.Println(err.Error())
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

func (r *Route) ExportJsonPositions() ([]string, error) {
	var route PartialRoutePosition
	var result []string
	total := len(r.Positions)

	for key, value := range r.Positions {
		route.ID = r.ID
		route.ClientId = r.ClientID
		route.Position = []float64{value.Latitude, value.Longitude}
		route.Finished = false

		if total-1 == key {
			route.Finished = true
		}

		jsonRoute, err := json.Marshal(route)

		if err != nil {
			return nil, err
		}

		result = append(result, string(jsonRoute))

	}

	return result, nil
}
