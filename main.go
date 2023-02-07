package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	route2 "github.com/mauFade/go-simulator-del/app/route"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err.Error())
	}
}

func main() {
	route := route2.Route{
		ID:       "1",
		ClientID: "1",
	}

	route.LoadPositions()

	stringJson, _ := route.ExportJsonPositions()

	fmt.Println(stringJson[0])
}
