package structs

import (
	"fmt"
	"math/rand"
)

type World struct {
	Name      string
	Cities    []City
	NumAliens int
}

func NewWorld(worldName string, worldMap map[string][][]string) World {
	world := World{Name: worldName}
	nameToCity := make(map[string]*City)
	for cityName, neighborInfo := range worldMap {
		pointerToCity, city := NewCity(cityName, neighborInfo)
		world.Cities = append(world.Cities, city)
		nameToCity[cityName] = pointerToCity
	}
	for i := 0; i <= len(world.Cities); i++ {
		city := world.Cities[i]
		if city.NorthByName != "" {
			city.North = nameToCity[city.NorthByName]
		}
		if city.SouthByName != "" {
			city.South = nameToCity[city.SouthByName]
		}
		if city.EastByName != "" {
			city.East = nameToCity[city.EastByName]
		}
		if city.WestByName != "" {
			city.West = nameToCity[city.WestByName]
		}
	}
	return world
}

func (worldX World) LaunchInvasion(numAliens int) {
	worldX.NumAliens = numAliens
	for i := 0; i <= numAliens; i++ {
		randCity := worldX.Cities[rand.Intn(len(worldX.Cities))]
		a := NewAlien(i, randCity)
		randCity.AddAlien(a)
	}
}

func (worldX World) DestroyCity(city City) {
	for i := range worldX.Cities {
		currCity := worldX.Cities[i]
		if currCity.Name == city.Name { //delete city
			worldX.Cities = append(worldX.Cities[:i], worldX.Cities[i+1:]...)
		} else { //delete bridges leading into the city to be destroyed
			currCity.DestroyBridge(city)
		}
	}
	numCityAliens := city.AlienCount
	worldX.NumAliens -= numCityAliens
}

func (worldX World) PrintMap() {
	cities := worldX.Cities
	for i := range cities {
		city := cities[i]
		directions := cities[i].Directions
		neighbors := ""
		for j := range directions {
			direction := directions[j]
			adjacentCity := ""
			switch direction {
			case "north": //todo(REFINE) : declare const directions, and iterate over the const directions in switch cases
				adjacentCity = city.NorthByName
			case "south":
				adjacentCity = city.SouthByName
			case "east":
				adjacentCity = city.EastByName
			case "west":
				adjacentCity = city.WestByName
			default:
				fmt.Println("PrintMap direction error")
			}
			neighbors += direction + "=" + adjacentCity + " "
		}
		fmt.Println(city, " ", neighbors)
	}
}
