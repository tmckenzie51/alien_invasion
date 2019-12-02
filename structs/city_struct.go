package structs
import "fmt"

//todo issues: equality issue, directions reference

type City struct {
    Name string
    AlienCount int
    Aliens []Alien
    Directions []string
    NorthByName,SouthByName, EastByName, WestByName string
    North, South, East, West City
}

func NewCity(name string,neighbors map[string]string){
  city := City{Name: name, AlienCount:0}
  for direction,adjacentCity := range neighbors{
    city.Directions = append(city.Directions,direction)
    dir := direction + "ByName" //todo
    city.*dir = adjacentCity
  }
}

func AddAlien(city City,a Alien) {
  city.AlienCount += 1
  city.Aliens = append(City.Aliens,a)
}

func RemoveAlien(city City, a Alien) {
  city.AlienCount -= 1
  for i := range city.Aliens{
    if city.Aliens[i] = a{ //todo
      city.Aliens = append(city.Aliens[:i], city.Aliens[i+1:]...)
    }
  }
}

func DestroyBridge(city City, destroyedCity City){
    for i := range city.Directions{
        direction := city.Directions[i]
        if city.direction == destroyedCity{ //todo
            city.direction = nil  //todo
        }
    }
}

