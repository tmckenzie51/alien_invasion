package structs

import "math/rand"

type Alien struct {
	Id       int
	Location City
}

func NewAlien(id int, location City) *Alien {
	a := Alien{Id: id, Location: location}
	return &a
}

func Travel(a Alien) *Alien {
	prevLocation := a.Location
	directions := prevLocation.Directions
	direction := directions[rand.Intn(len(directions))] //todo: fix the direction part
	newLocation := prevLocation.*direction
	City.RemoveAlien(prevLocation,a) //todo: fix these
	City.AddAlien(newLocation,a)
	a.Location = newLocation
	return &a
}
