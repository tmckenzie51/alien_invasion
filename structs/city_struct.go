//Assumption: Each city name is unique. There will be no 2 cities with the same name

package structs

//todo issues: directions reference

type City struct {
	//Parent *City  : doesn't work to fix recursive issue
	Name                                             string
	AlienCount                                       int
	Aliens                                           []Alien
	Directions                                       []string
	NorthByName, SouthByName, EastByName, WestByName string
	North, South, East, West                         *City
}

func NewCity(name string, neighbors []string) *City {
	city := City{Name: name, AlienCount: 0}
	for i := range neighbors{ //todp: parse string
	    direction  =
    }
	for direction, _ := range neighbors {
		city.Directions = append(city.Directions, direction) //todo: is direction capitalized?
		dir := direction + "ByName"                          //todo
		city.dir = adjacentCity
	}
	return &city
}

//todo: add functions (add, remove, destroybridges etc.