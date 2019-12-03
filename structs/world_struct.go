package structs

import (
	"fmt"
	"math/rand"
)

type World struct {
	Name      string
	Cities    []*City
	NumAliens int
}

func NewWorld(worldName string, worldMap map[string][][]string) *World {
	world := World{Name: worldName}
	nameToCity := make(map[string]*City)

	//Create main cities listed
	for cityName, neighborInfo := range worldMap {
		city := NewCity(cityName, neighborInfo)
		world.Cities = append(world.Cities, city)
		nameToCity[cityName] = city
	}

	//Convert from direction from string to City type
	for i := 0; i < len(world.Cities); i++ {
		city := world.Cities[i]
		if city.NorthByName != "" {
			city.North = nameToCity[city.NorthByName]
			if !contains(world.Cities, city.NorthByName) {
				northCity := NewCity(city.NorthByName, nil)
				world.Cities = append(world.Cities, northCity)
				nameToCity[city.NorthByName] = northCity
			}
		}
		if city.SouthByName != "" {
			city.South = nameToCity[city.SouthByName]
			if !contains(world.Cities, city.SouthByName) {
				southCity := NewCity(city.SouthByName, nil)
				world.Cities = append(world.Cities, southCity)
				nameToCity[city.SouthByName] = southCity
			}
		}
		if city.EastByName != "" {
			city.East = nameToCity[city.EastByName]
			if !contains(world.Cities, city.EastByName) {
				eastCity := NewCity(city.EastByName, nil)
				world.Cities = append(world.Cities, eastCity)
				nameToCity[city.SouthByName] = eastCity
			}
		}
		if city.WestByName != "" {
			city.West = nameToCity[city.WestByName]
			if !contains(world.Cities, city.WestByName) {
				westCity := NewCity(city.WestByName, nil)
				world.Cities = append(world.Cities, westCity)
				nameToCity[city.WestByName] = westCity
			}
		}
	}
	return &world
}

func contains(cities []*City, cityName string) bool {
	for i := range cities {
		if cities[i].Name == cityName {
			return true
		}
	}
	return false
}

func (worldX *World) LaunchInvasion(numAliens int) *World {
	worldX.NumAliens = numAliens
	for i := 0; i < numAliens; i++ {
		randIndex := rand.Intn(len(worldX.Cities))
		randCity := worldX.Cities[randIndex]
		a := NewAlien(i)
		randCity.AddAlien(a)
		fmt.Println(randCity.AlienCount, " aliens in ", randCity.Name) //todo: remove later
	}
	return worldX
}

func (worldX World) DestroyCity(city *City) {
	fmt.Println("city being destroyed: ", city.Name)
	var cityIndex []int //list of indices of cities to be destroyed in the worldX.Cities list/array
	for i := range worldX.Cities {
		currCity := worldX.Cities[i]
		if currCity.Name != city.Name { //delete bridges leading into the city to be destroyed
			currCity.DestroyBridge(city)
		} else {
			cityIndex = append(cityIndex, i)
		}
	}

	//delete cities from world
	for j := range cityIndex {
		index := cityIndex[j]
		worldX.Cities = append(worldX.Cities[:index], worldX.Cities[index+1:]...)
	}

	//update number of aliens in the world
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
		if neighbors != "" {
			fmt.Println(city.Name, " ", neighbors)
		}

	}
}
