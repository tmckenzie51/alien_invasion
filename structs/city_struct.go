//Assumption: Each city name is unique. There will be no 2 cities with the same name

package structs

import "fmt"

type City struct {
	Name                                             string
	AlienCount                                       int
	Aliens                                           []*Alien
	Directions                                       []string
	NorthByName, SouthByName, EastByName, WestByName string
	North, South, East, West                         *City
}

func NewCity(name string, neighborInfo [][]string) *City {
	city := City{Name: name, AlienCount: 0}
	for i := range neighborInfo {
		direction := neighborInfo[i][0] //example direction : "north". Note: direction is not capitalized
		adjacentCity := neighborInfo[i][1]
		city.Directions = append(city.Directions, direction)
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

//todo: there might be some issues with alien tracking because currently not using pointers to reference aliens
//todo: maybe return city.aliens here instead
func (city *City) AddAlien(a *Alien) *City {
	city.AlienCount += 1
	city.Aliens = append(city.Aliens, a)

	if len(city.Aliens) > 1 {
		fmt.Println("added alien to ", city.Name, ". Aliens =  ")
		for i := range city.Aliens {
			fmt.Println(city.Aliens[i].Id)
		}
	}

	return city

}

func (city *City) DestroyBridge(destroyedCity *City) {
	fmt.Println("Destroying bridges from", city.Name, " to ", destroyedCity.Name)
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

	//update directions/bridges leading out from the city
	for j := range directionsIndex {
		index := directionsIndex[j]
		city.Directions = append(city.Directions[:index], city.Directions[index+1:]...)
	}
}

//todo: there might be some issues with alien tracking because currently not using pointers to reference aliens
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
