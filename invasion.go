//Assumption 1:
//For each direction (north, sotuh,east, west), there will be at most 1 city in that direction.
//in other words, there can not be a road in any direction that leads to mutiple cities
//in terms of our input, this means that "Foo north=Bar,Bee" is not a valid input

//Assumption 2:
//all inputed world_maps meet the afformentioned standards/requirements of a the world, such as no multiple destinations from the same road,
//cities names are all string and do not contain numbers (int or float), and the syntactical requirements of the input is met
//as specified in the problem specifics detailed in the alien_invasion pdf document.

package main
import (
    "alien_invasion/structs"
    "fmt"
)

func fight_and_destroy() *structs.World{
  cities = world_x.cities
  for i, curr_city := range cities {
    num_aliens = curr_city.alien_count
    aliens = curr_city.aliens()
    if num_aliens >= 2{
      world_x.destroy(structs.city)
      curr_city.fight()
    }
    fmt.Printf("%d has been destroyed by aliens %d\n",curr_city,aliens)
  }
}

func wander(){
  cities = world_x.cities
  for i,curr_city := range cities{
    a = curr_city.aliens
    a[0].travel()
  }
}

func invasion(num_aliens int){
  world_x.launch_invasion(num_aliens)
  for i := range(10000){
      fight_and_destroy()
      num_aliens = world_x.alien_count
      if num_aliens <=0 {
        world_x.print_map()
        return
      }
      else{
        wander()
      }
  }
}

num_aliens = 3
invasion(num_aliens)
world_map = file.read("X.txt").tomap() //todo: go
var world_x = structs.world.new_world('x',world_map)
