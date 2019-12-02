package structs

import "math/rand"

type alien struct {
	id       int
	location City
}

func new_alien(id int, location string) *alien {
	a := alien{id: id, location: location}
	return &a
}

func travel(a) *alien {
	prev_location = a.location
	directions = prev_location.directions
	direction = directions[rand.Intn(len(directions))]
	new_location = prev_location.direction
	prev_location.remove_alien(a)
	new_location.add_alien(a)
	a.location = new_location
}
