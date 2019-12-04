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
		case "south":
			newLocation := prevLocation.South
			newCity = newLocation.AddAlien(a)
		case "east":
			newLocation := prevLocation.East
			newCity = newLocation.AddAlien(a)
		case "west":
			newLocation := prevLocation.West
			newCity = newLocation.AddAlien(a)
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
