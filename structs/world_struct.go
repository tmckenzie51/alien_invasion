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

func NewWorld(worldName string, worldMap map[string][]string) *World { //todo: why is this function unused?
    world := World{Name: worldName}
    nameToCity := make(map[string]City)
    for cityName, neighbors := range worldMap {
        city := NewCity(cityName, neighbors) //todo: unresolved references to all city fields
        world.Cities = append(world.Cities, city)
        nameToCity[cityName] = city
    }
    for city := range world.Cities {
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
    return &world
}


func (worldX World) LaunchInvasion(numAliens int) *World{
    worldX.NumAliens = numAliens
    for i := 1;  i<=numAliens; i++ {
        randCity := worldX.Cities[rand.Intn(len(worldX.Cities))]
        a = Alien.NewAlien(i,randCity) //todo: whyyy unresolved reference?
        randCity.AlienCount += 1
        randCity.Aliens = append(randCity.Aliens,a)
        rand_city.add_alien(a)
    }
    return &world
}

func (worldX World) destroyCity(city City) *World{
    for i, curr_city := range world.cities{
      if curr_city == city{ //delete city
        world.cities = append(world.cities[:i], world.cities[i+1:]...)
      }
      else { //delete bridges leading into the city to be destroyed

        curr_city.destroy_bridge(curr_city,city)

        //Destroy bridges leading into the destroyed city
        for i := range curr_city.Directions{
            direction := curr_city.Directions[i] //todo
            if curr_city.direction.Name == city.Name{
                curr_city.direction = nil
                dir = direction + "ByName"
                curr_city.dir = nil
                curr_city.Directions = append(curr_city.Directions[:i], curr_city.Directions[i+1:]...) //remove direction from directions
            }
        }
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
