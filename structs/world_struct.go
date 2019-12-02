package structs
import (
    "fmt"
    "math/rand"
)

type World struct {
    Name string
    Cities []City
    NumAliens int
}

func NewWorld(worldName string, worldMap map[string][]string) *World {
    world := World{Name: worldName}
    nameToCity := make(map[string]City)
    for cityName, neighbors := range worldMap {
        city = City.NewCity(cityName, neighbors)
        world.Cities = append(world.cities, city)
        nameToCity[city_name] = city
    }
    for city := range world.cities {
        if city.NorthByName != nil {
            city.North = nameToCity[city.NorthByName]
        }
        if city.SouthByName != nil {
            city.South = nameToCity[city.SouthByName]
        }
        if city.EastByName != nil {
            city.East = nameToCity[city.EastByName]
        }
        if city.WestByName != nil {
            city.West = nameToCity[city.WestByName]
        }
    }
}



func LaunchInvasion(numAliens int) *World{
    world := World{NumAliens: numAliens} //refernce to world giving issues
    for i := range numAliens{
        randCity = World.Cities[rand.Intn(len(World.Cities))]
        a = Alien.NewAlien(i,randCity)
        randCity.AlienCount += 1
        randCity.Aliens = append(randCity.Aliens,a)
        rand_city.add_alien(a)
    }
    return &world
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
