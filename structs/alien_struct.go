package structs

import (
	"fmt"
	"math/rand"
)

type Alien struct {
	Id int
}

func NewAlien(id int) *Alien {
	a := Alien{Id: id}
	return &a
}

func (a *Alien) Travel(city *City) *City {
	prevLocation := city
	var newCity *City
	if len(prevLocation.Directions) > 0 {
		direction := prevLocation.Directions[rand.Intn(len(prevLocation.Directions))]
		switch direction {
		case "north":
			newLocation := prevLocation.North
			newCity = newLocation.AddAlien(a)
			fmt.Println("Alien", a.Id, " moved from ", prevLocation.Name, "to ", newLocation.Name)
		case "south":
			newLocation := prevLocation.South
			newCity = newLocation.AddAlien(a)
			fmt.Println("Alien", a.Id, " moved from ", prevLocation.Name, "to ", newLocation.Name)
		case "east":
			newLocation := prevLocation.East
			newCity = newLocation.AddAlien(a)
			fmt.Println("Alien", a.Id, " moved from ", prevLocation.Name, "to ", newLocation.Name)
		case "west":
			newLocation := prevLocation.West
			newCity = newLocation.AddAlien(a) //todo: just checking to see if this could work. maybe remove later, or add more of later
			fmt.Println("Alien ", a.Id, "moved from ", prevLocation.Name, "to ", newCity.Name)
		default:
			fmt.Println("travel direction error")
		}
	}
	if newCity == nil {
		return prevLocation
	} else {
		prevLocation.RemoveAlien(a)
		return newCity
	}
}
