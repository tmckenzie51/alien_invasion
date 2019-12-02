package structs
import "fmt"

type city struct {
    name string
    alien_count int
    aliens []alien{}
    directions []string
    north, south, east, west city{}
}

//todo: not sure if map[string]city{} will work as a type because that city may not be declared as yet.
func new_city(name string,neighbors map[string]city{}) *city{
  city.name = city_name
  for direction,adjacent_city := neighbors{
    city.directions = append(city.directions,direction)
    city.direction = adjacent_city
  }
}

func add_alien(a alien{}) *city{
  city.alien_count += 1
  city.aliens = append(city.aliens,a)
}

func remove_alien(a alien{}) *city {
  city.alien_count -= 1
  for i := range city.aliens{
    if city.aliens[i] == a{
      city.aliens = append(city.aliens[:i], city.aliens[i+1:]...)
    }
  }
}

func destroy_bridge(curr_city city{}, destroyed_city city{}){
  for i : range curr_city.directions{
    direction = curr_city.directions[i]
    if curr_city.direction == destroyed_city{
      curr_city.direction = nil
    }
  }
}
