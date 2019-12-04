//Assumption 1:
//For each direction (north, south,east, west), there will be at most 1 city in that direction.
//in other words, there can not be a road in any direction that leads to multiple cities
//in terms of our input, this means that "Foo north=Bar,Bee" is not a valid input

//Assumption 2:
//all inputted world_maps meet the aforementioned standards/requirements of a the world, such as no multiple destinations from the same road,
//cities names are all string and do not contain numbers (int or float), and the syntactical requirements of the input is met
//as specified in the problem specifics detailed in the alien_invasion pdf document.

package main

import (
	"alien_invasion/structs"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	worldMap := read("./tests/X.txt")
	numAliens, _ := strconv.Atoi(os.Args[1])
	worldX := structs.NewWorld("X", worldMap)

	//check status of world
	fmt.Println("checking status of world after creation")
	worldX.PrintMap()

	worldX = worldX.LaunchInvasion(numAliens)

	for i := 0; i < 10000; i++ { //todo: optimize for trap cities
		worldX = fightAndDestroy(worldX)
		numAliens = worldX.NumAliens
		if numAliens == 0 || i == 9999 {
			fmt.Println("at program end")
			worldX.PrintMap() //todo: print "world empty or something. Maybe also print trap cities as well
			return
		} else {
			worldX = wander(worldX)

			//checking if the corect worldX is returned
			if len(worldX.Cities) > 0 {
				for j := range worldX.Cities {
					currCity := worldX.Cities[j]
					if currCity.Name == "Baz" {
						fmt.Println("checking if the worldX returned from wander is correct")
						for k := range currCity.Aliens {
							fmt.Println(currCity.Aliens[k])
						}
					}
				}
			}

		}
	}
}

func read(fileName string) map[string][][]string {
	worldMap := make(map[string][][]string)
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtLines []string

	for scanner.Scan() {
		txtLines = append(txtLines, scanner.Text())
	}

	file.Close() //todo: unhandled error?

	for _, eachLine := range txtLines { //example eachLine:  "Foo north=Bar west=Baz South=Qu-ux"
		cityInfo := strings.Split(eachLine, " ") //example cityInfo : ["Foo", "north=Bar", "west=Baz", "South=Qu-ux"]
		cityName := cityInfo[0]                  //example: "Foo"
		for i := range cityInfo {
			if i > 0 {
				neighborInfo := strings.Split(cityInfo[i], "=")               // example neighborInfo : ["north","bar"]
				worldMap[cityName] = append(worldMap[cityName], neighborInfo) //worldMap["Foo"] = [["north","Bar"], ["west",Baz"], ["South=Qu-ux"]
			}
		}
	}
	return worldMap
}

func fightAndDestroy(worldX *structs.World) *structs.World {
	cities := worldX.Cities
	var citiesToDestroy []*structs.City

	//get indices of cities to be destroyed
	for i := range cities {
		currCity := cities[i]
		numAliens := currCity.AlienCount
		if numAliens >= 2 {
			citiesToDestroy = append(citiesToDestroy, currCity)
		}
	}
	//Destroy cities at specified indices
	for j := range citiesToDestroy {
		currCity := citiesToDestroy[j]
		worldX.DestroyCity(currCity)
		aliens := currCity.Aliens //todo: fix: not the correct list of aliens at currcity
		alienNames := getNames(aliens)
		destroyStatement := currCity.Name + " has been destroyed by aliens: " + alienNames
		fmt.Println(destroyStatement)
	}
	return worldX
}

//todo: fix - the aliens list here is inaccurate
func getNames(aliens []*structs.Alien) string {
	alienNames := ""
	for i := range aliens {
		currAlien := aliens[i]
		if i == len(aliens)-1 {
			alienNames += strconv.Itoa(currAlien.Id)
		} else if (i == len(aliens)-2) && (len(aliens) > 1) {
			alienNames += strconv.Itoa(currAlien.Id) + " and "
		} else {
			alienNames += strconv.Itoa(currAlien.Id) + ", "
		}
	}
	return alienNames
}

func wander(worldX *structs.World) *structs.World {
	cities := worldX.Cities
	for i := range cities {
		currCity := cities[i]
		//fmt.Println("wandering from city: ", currCity.Name)
		aliens := currCity.Aliens
		if len(aliens) > 0 {
			a := aliens[0]
			//a.Travel(currCity)
			newCity := a.Travel(currCity)
			replaceWith(worldX.Cities, newCity)
		}
	}

	//check that cities are correct
	for i := range worldX.Cities {
		city := cities[i]
		if city.Name == "Baz" {
			fmt.Println("check cities in wander")
			for j := range city.Aliens {
				fmt.Println(city.Aliens[j].Id)
			}
		}
	}

	worldX.Cities = cities
	return worldX //todo: is this returning old or new world X?
}

//TODO: REMOVE FUNCTION LATER. BOTH OLD AND NEW CITIES ARE CORRECT
func replaceWith(cities []*structs.City, newCity *structs.City) {
	for i := range cities {
		city := cities[i]
		fmt.Println(city.Name)
		fmt.Println(newCity.Name)

		fmt.Println(city.Aliens)
		if city.Name == newCity.Name {

			//print aliens at city to see if it is correct or no
			fmt.Println("check old city")
			for j := range city.Aliens {
				fmt.Println(city.Aliens[j].Id)
			}

			//print aliens in new city to see if correct or no
			fmt.Println("check new City")
			for j := range newCity.Aliens {
				fmt.Println(newCity.Aliens[j].Id)
			}

			//update cities
			cities[i] = newCity

			//print aliens in  city to see if updated correctly  or no
			fmt.Println("check update")
			for j := range cities[i].Aliens {
				fmt.Println(cities[i].Aliens[j].Id)
			}
		}
	}
}
