package structs

import (
	"fmt"
	"math/rand"
)

type Alien struct {
	Id int
}

func NewAlien(id int) Alien {
	a := Alien{Id: id}
	return a
}

func (a Alien) Travel(city *City) {
	prevLocation := city
	prevLocation.RemoveAlien(a)
	if len(prevLocation.Directions) > 0 {
		direction := prevLocation.Directions[rand.Intn(len(prevLocation.Directions))]
		switch direction {
		case "north":
			newLocation := prevLocation.North
			newLocation.AddAlien(a)
			fmt.Println("Alien moved from ", prevLocation.Name, "to ", newLocation.Name)
		case "south":
			newLocation := prevLocation.South
			newLocation.AddAlien(a)
			fmt.Println("Alien moved from ", prevLocation.Name, "to ", newLocation.Name)
		case "east":
			newLocation := prevLocation.East
			newLocation.AddAlien(a)
			fmt.Println("Alien moved from ", prevLocation.Name, "to ", newLocation.Name)
		case "west":
			newLocation := prevLocation.West
			newLocation.AddAlien(a)
			fmt.Println("Alien moved from ", prevLocation.Name, "to ", newLocation.Name)
		default:
			fmt.Println("travel direction error")
		}
	}
}
