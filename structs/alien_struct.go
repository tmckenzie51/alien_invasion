package structs

import (
	"fmt"
	"math/rand"
)

//todo: Issues: direction issue, unused function issue
type Alien struct {
	Id int
}

func NewAlien(id int) Alien {
	a := Alien{Id: id}
	return a
}

func (a Alien) Travel(city City) {
	prevLocation := city
	direction := prevLocation.Directions[rand.Intn(len(prevLocation.Directions))]
	prevLocation.RemoveAlien(a)
	switch direction {
	case "north": //todo (REFINE): declare const directions, and iterate over the const directions in switch cases
		newLocation := prevLocation.North
		newLocation.AddAlien(a)
	case "south":
		newLocation := prevLocation.South
		newLocation.AddAlien(a)
	case "east":
		newLocation := prevLocation.East
		newLocation.AddAlien(a)
	case "west":
		newLocation := prevLocation.West
		newLocation.AddAlien(a)
	default:
		fmt.Println("travel direction error")
	}

}
