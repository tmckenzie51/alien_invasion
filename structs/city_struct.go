//Assumption: Each city name is unique. There will be no 2 cities with the same name

package structs

import "fmt"

type City struct {
	Name                                             string
	AlienCount                                       int
	Aliens                                           []Alien
	Directions                                       []string
	NorthByName, SouthByName, EastByName, WestByName string
	North, South, East, West                         *City
}

func NewCity(name string, neighborInfo [][]string) (*City, City) {
	city := City{Name: name, AlienCount: 0}
	for i := range neighborInfo {
		direction := neighborInfo[i][0] //example direction : "north". Note: direction is not capitalized
		adjacentCity := neighborInfo[i][1]
		city.Directions = append(city.Directions, direction)
		switch direction {
		case "north": //todo(REFINE) : declare const directions, and iterate over the const directions in switch cases
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
	return &city, city
}

func (city City) AddAlien(a Alien) {
	city.AlienCount += 1
	city.Aliens = append(city.Aliens, a)
}

func (city City) DestroyBridge(destroyedCity City) {
	for i := range city.Directions {
		direction := city.Directions[i]
		adjacentCity := ""
		switch direction {
		case "north": //todo(REFINE) : declare const directions, and iterate over the const directions in switch cases
			adjacentCity = city.NorthByName
			if adjacentCity == city.Name {
				city.North = nil
				city.NorthByName = ""
				city.Directions = append(city.Directions[:i], city.Directions[i+1:]...)
			}
		case "south":
			adjacentCity = city.SouthByName
			if adjacentCity == city.Name {
				city.South = nil
				city.SouthByName = ""
				city.Directions = append(city.Directions[:i], city.Directions[i+1:]...)
			}
		case "east":
			adjacentCity = city.EastByName
			if adjacentCity == city.Name {
				city.East = nil
				city.EastByName = ""
				city.Directions = append(city.Directions[:i], city.Directions[i+1:]...)
			}
		case "west":
			adjacentCity = city.WestByName
			if adjacentCity == city.Name {
				city.West = nil
				city.WestByName = ""
				city.Directions = append(city.Directions[:i], city.Directions[i+1:]...)
			}
		default:
			fmt.Println("DestroyBridge direction error")
		}
	}
}

func (city City) RemoveAlien(a Alien) {
	city.AlienCount -= 1
	for i := range city.Aliens {
		if city.Aliens[i].Id == a.Id {
			city.Aliens = append(city.Aliens[:i], city.Aliens[i+1:]...)
		}
	}
}
