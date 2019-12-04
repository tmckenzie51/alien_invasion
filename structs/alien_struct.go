package structs

import (
	"fmt"
	"math/rand"
)

type Alien struct {
	Id int
}

//Create a new Alien object
func NewAlien(id int) *Alien {
	a := Alien{Id: id}
	return &a
}

//move alien from current location city to a new neighboring city, that has a bridge leading to it.
func (a *Alien) Travel(city *City) *City {
	prevLocation := city
	var newCity *City
	if len(prevLocation.Directions) > 0 { //only consider moving alien, when we confirm that the prevLocation is not a trap-city
		direction := prevLocation.Directions[rand.Intn(len(prevLocation.Directions))]
		switch direction {
		case "north":
			newLocation := prevLocation.North
			newCity = newLocation.AddAlien(a) //add alien to it's new city location
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
		return prevLocation //in the case that there was an error, or in the case that tried to move a trapped alien, return prevLocation as aliens current location
	} else {
		prevLocation.RemoveAlien(a) //remove alien from prev location and return alien's new location - newCity
		return newCity
	}
}
