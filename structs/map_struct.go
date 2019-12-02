package structs
import "fmt"
//todo: do I have to do package main and import fmt here?

type world struct {
    city string
    neighbors map[string]city{}

}

type city struct{
  name, num_bridges,north,south,east,west string
}


func new_world(m) *world *city{
  //todo: set alien count to 0 in new world
  for city_name, neighbors := range m{
        world.city = city
        city.name = city_name
        world.neighbors =
        for neighbor in neighbors{
          direction = neigbor[0]
          neighboring_city = neighbor[1]
          world.*direction = neighboring_city
        }
    }

func destroy_city(curr_city string) *world{
    world.*curr_city
    a.location := get_rand_loc
    aleins := append(aliens, a)
    return &a
}


//randomly assign aliens to their initial locations
def launch_invasion(num_aliens,world_map):
    world_x = world.new_world(world_map)
    for i := range num_aliens {
      alien.launch_invasion(num_aliens,world_map)
      num_cities = world.num_cities()
      city_index = rand.Intn(num_cities)
      assigned_city = cities[city_index]
      alien_tracker[assigned_city].append(i)
    }


    def destroy(city,world_map):
        world_map.remove(city)
        for curr_city,adjacent_cities in world_map.iteritems():
            for location in adjacent_cities:
                direction = location[0]
                adj_city = location[1]
                if adj_city == city:
                    world_map[curr_city].remove(location)
        return world_map

        func fight(a1,a2) {
          for i, v := range aliens {
            if v == a1 or v==a2:
              aliens = append(aliens[:i], aliens[i+1:]...)
              //todo : remember to print out that these aliens have died
        	}
        }

        // todo: not sure if this will work
        func alien_inventory(){
          return aliens
        }

        //a.location := get_rand_loc
        //aleins := append(aliens, a)
        //return &a
        //var aliens := []int{}




        //city things
        prev_location.alien_count -= 1
        prev_location.remove_alien(a)
        new_location.alien_count += 1
        new_location.add_alien(a)
