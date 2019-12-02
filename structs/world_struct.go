package structs
import (
    "fmt"
    "math/rand"
)
//todo: do I have to do package main and import fmt here?

type World struct {
    name string
    cities []city{}
    num_aliens int
}

func new_world(world_name string, world_map map[string][]string) *world{
  world.name = world_name
  //todo: figure out how to add neighbors
  city_neighbors_map := map[city{}][]city{}
  for city_name, neighbors := range m{
        curr_city = city.new_city(city_name)
        world.cities = append(world.cities,curr_city)

  }
  for city_name,neighbors := range m{

  }

func launch_invasion(num_aliens int) *world{
  world.num_aliens = num_aliens
  for i := range num_aliens{
    a = alien.new_alien(i)
    rand_city = world.cities[rand.Intn(len(cities))]
    rand_city.add_alien(a)
  }
}

func destroy_city(city city {}) *world{
    for i, curr_city := range world.cities{
      if curr_city == city{ //delete city
        world.cities = append(world.cities[:i], world.cities[i+1:]...)
      }
      else { //delete bridges leading into the city to be destroyed
        curr_city.destroy_bridge(curr_city,city)
          }
    }
    num_city_aliens = city.alien_count
    world.num_aliens -= num_city_aliens
}

func print_map(){
  cities = world.cities
  for i := range cities{
    city = cities[i]
    directions = world.cities[i].directions
    neighbors = ""
    for j := range directions{
      direction = directions[j]
      neighbors += direction + "=" + city.direction.name + " "
    }
    fmt.Println(city, " ", neighbors)
  }
}
