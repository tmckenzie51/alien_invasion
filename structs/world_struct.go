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

//Create a new World object  using the world map read in from the input file
func NewWorld(worldName string, worldMap map[string][][]string) *World {
	world := World{Name: worldName}
	nameToCity := make(map[string]*City)

	//Create main cities listed, wher emain cities have their own line in the input file (i.e. main cities are not trap cities)
	for cityName, neighborInfo := range worldMap {
		city := NewCity(cityName, neighborInfo)
		world.Cities = append(world.Cities, city)
		nameToCity[cityName] = city //save cityName as it corresponds to the relevant city pointer.
	}

	//Convert from neighbor of main city from the string type to the city type
	for i := 0; i < len(world.Cities); i++ {
		city := world.Cities[i]
		if city.NorthByName != "" {
			//add trap city to world
			if !contains(world.Cities, city.NorthByName) { //in the case that the city has a neighboring city that was not included in the world's cities when main cities were being added (i.e. in the case that the city has a trap-city neighbor)
				northCity := NewCity(city.NorthByName, nil)
				world.Cities = append(world.Cities, northCity) //add city's trap-city neighbor to the world's cities
				nameToCity[city.NorthByName] = northCity
			}
			city.North = nameToCity[city.NorthByName]
		}
		if city.SouthByName != "" {
			//add trap city to world
			if !contains(world.Cities, city.SouthByName) {
				southCity := NewCity(city.SouthByName, nil)
				world.Cities = append(world.Cities, southCity)
				nameToCity[city.SouthByName] = southCity
			}
			city.South = nameToCity[city.SouthByName]
		}
		if city.EastByName != "" {
			if !contains(world.Cities, city.EastByName) {
				eastCity := NewCity(city.EastByName, nil)
				world.Cities = append(world.Cities, eastCity)
				nameToCity[city.SouthByName] = eastCity
			}
			city.East = nameToCity[city.EastByName]
		}
		if city.WestByName != "" {
			if !contains(world.Cities, city.WestByName) {
				westCity := NewCity(city.WestByName, nil)
				world.Cities = append(world.Cities, westCity)
				nameToCity[city.WestByName] = westCity
			}
			city.West = nameToCity[city.WestByName]
		}
	}
	return &world
}

//Check to seee if a city is already included in the list of the world's cities - return true/false
func contains(cities []*City, cityName string) bool {
	for i := range cities {
		if cities[i].Name == cityName {
			return true
		}
	}
	return false
}

//Launch inital alien deplyment/invasion on cities by randomly selecting a city for each of the alien in the invasion party
func (worldX *World) LaunchInvasion(numAliens int) *World {
	worldX.NumAliens = numAliens
	for i := 0; i < numAliens; i++ {
		randIndex := rand.Intn(len(worldX.Cities))
		randCity := worldX.Cities[randIndex]
		a := NewAlien(i)     //create new alien
		randCity.AddAlien(a) //add new alien to the randomly selected city
	}
	return worldX
}

//Scan through the world's cities and destroy the city specified in the parameter, as 'city' and call on destroyBridge method to
//also destroy all bridges leading to this city.
func (worldX *World) DestroyCity(city *City) {
	var cityIndex []int //list of indices of cities to be destroyed in the worldX.Cities list/array
	for i := range worldX.Cities {
		currCity := worldX.Cities[i]
		if currCity.Name != city.Name { //delete bridges leading into the city to be destroyed
			currCity.DestroyBridge(city)
		} else {
			cityIndex = append(cityIndex, i)
		}
	}

	//delete citiy from world
	for j := range cityIndex {
		index := cityIndex[j]
		worldX.Cities = append(worldX.Cities[:index], worldX.Cities[index+1:]...)
	}

	//update number of aliens in the world to reflect the alien deaths in the alien fight that took place in the city
	numCityAliens := city.AlienCount
	worldX.NumAliens -= numCityAliens
}

//Print the current world map in the same format as is given in the input test files, with each non-trap city on its own line
// followed by whitespace and the bridge-destination city pairs (which are connected with an '=' sign'
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
			case "north":
				//adjacentCity = city.NorthByName
				adjacentCity = city.North.Name
			case "south":
				//adjacentCity = city.SouthByName
				adjacentCity = city.South.Name
			case "east":
				//adjacentCity = city.EastByName
				adjacentCity = city.East.Name
			case "west":
				//adjacentCity = city.WestByName
				adjacentCity = city.West.Name
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
