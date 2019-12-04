//Assumption: Each city name is unique. There will be no 2 cities with the same name

package structs

import "fmt"

type City struct {
	Name                                             string   //Name of city
	AlienCount                                       int      //number of aliens in city
	Aliens                                           []*Alien //list of aliens in city
	Directions                                       []string // Bridges leading out from city
	NorthByName, SouthByName, EastByName, WestByName string   //string name of bridge destinations from city
	North, South, East, West                         *City    // bridge city type destinations
}

//Create a new City object
func NewCity(name string, neighborInfo [][]string) *City {
	city := City{Name: name, AlienCount: 0} //initalize alien count in city as 0
	for i := range neighborInfo {
		direction := neighborInfo[i][0]
		adjacentCity := neighborInfo[i][1]
		city.Directions = append(city.Directions, direction)
		//Assign names of destinations of bridges from city
		switch direction {
		case "north":
			city.NorthByName = adjacentCity
		case "south":
			city.SouthByName = adjacentCity
		case "east":
			city.EastByName = adjacentCity
		case "west":
			city.WestByName = adjacentCity
		default:
			fmt.Println("NewCity declaration direction error")
		}
	}
	return &city
}

//Add an alien to city
func (city *City) AddAlien(a *Alien) *City {
	city.AlienCount += 1
	city.Aliens = append(city.Aliens, a)
	return city
}

//Scan city for bridges that may lead to a destination city previously destroyed. If such a bridge is found, destroy it.
func (city *City) DestroyBridge(destroyedCity *City) {
	var directionsIndex []int
	for i := range city.Directions {
		direction := city.Directions[i]
		switch direction {
		case "north":
			adjacentCity := city.North
			if adjacentCity.Name == destroyedCity.Name {
				city.North = nil
				city.NorthByName = ""
				directionsIndex = append(directionsIndex, i)
			}
		case "south":
			adjacentCity := city.South
			if adjacentCity.Name == destroyedCity.Name {
				city.South = nil
				city.SouthByName = ""
				directionsIndex = append(directionsIndex, i)
			}
		case "east":
			adjacentCity := city.East
			if adjacentCity.Name == destroyedCity.Name {
				city.East = nil
				city.EastByName = ""
				directionsIndex = append(directionsIndex, i)
			}
		case "west":
			adjacentCity := city.West
			if adjacentCity.Name == destroyedCity.Name {
				city.West = nil
				city.WestByName = ""
				directionsIndex = append(directionsIndex, i)
			}
		default:
			fmt.Println("DestroyBridge direction error")
		}
	}

	//update bridges leading out from the city
	for j := range directionsIndex {
		index := directionsIndex[j]
		city.Directions = append(city.Directions[:index], city.Directions[index+1:]...)
	}
}

//Remove an alien form a city (usually called when an alien moves from this city to another when wandering around)
func (city City) RemoveAlien(a *Alien) {
	city.AlienCount -= 1
	var alienIndex []int
	for i := range city.Aliens {
		if city.Aliens[i].Id == a.Id {
			alienIndex = append(alienIndex, i)
		}
	}

	for i := range alienIndex {
		index := alienIndex[i]
		city.Aliens = append(city.Aliens[:index], city.Aliens[index+1:]...)
	}
}
