package main

import (
	"fmt"
	"log"

	"github.com/gitfuf/misc/funcoptions/api"
)

func main() {
	house, err := api.NewHouse("Fuf")
	if err != nil {
		log.Fatal(err)
	}
	print(house)

	house1, err := api.NewHouse("Maria", api.Toilets(4), api.Rooms(10))
	if err != nil {
		log.Fatal(err)
	}
	print(house1)
}

func print(house *api.House) {
	fmt.Printf("Was created house %s:\n rooms: %d, toilets: %d, garden=%t\n ", house.Name, house.Rooms, house.Toilets, house.HasGarden)
}
