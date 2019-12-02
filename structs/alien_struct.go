package structs

import (
	"fmt"
	"math/rand"
)

//todo: Issues: direction issue, unused function issue
type Alien struct {
	Id       int
	Location City
}

func NewAlien(id int, location City) Alien {
	a := Alien{Id: id, Location: location}
	return a
}

func (a Alien) Travel() *Alien {
	prevLocation := a.Location
	directions := prevLocation.Directions
	direction := directions[rand.Intn(len(directions))]
	switch direction { //todo : complete switch statement
	case "North": //todo: declare const directions, and iterate over the const directions in switch cases
		newLocation := prevLocation.North
	case "South":
		newLocation := prevLocation.South
	case "East":
		newLocation := prevLocation.East
	case "West":
		newLocation := prevLocation.West
	default:
		fmt.Println("travel direction error")
	}

	//Remove Alien from previous city
	prevLocation.AlienCount -= 1
	for i := range prevLocation.Aliens {
		if prevLocation.Aliens[i].Id == a.Id {
			prevLocation.Aliens = append(prevLocation.Aliens[:i], prevLocation.Aliens[i+1:]...)
		}
	}

	//Add Alien to new city
	newLocation.AlienCount += 1
	newLocation.Aliens = append(newLocation.Aliens, a)

	a.Location = newLocation

	return &a
}
