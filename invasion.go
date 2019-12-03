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
	fmt.Println("just created this world")
	worldX.PrintMap()
	worldX = worldX.LaunchInvasion(numAliens)
	fmt.Println("world after launch: ")
	worldX.PrintMap()
	for i := 0; i < 10000; i++ {
		worldX = fightAndDestroy(worldX)
		numAliens = worldX.NumAliens
		if numAliens == 0 || i == 9999 {
			worldX.PrintMap()
			return
		} else {
			worldX = wander(worldX)
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
	for i := range cities { //cities is getting confused because elements are being destroyed while in the loop
		currCity := cities[i]
		numAliens := currCity.AlienCount
		if numAliens >= 2 {
			citiesToDestroy = append(citiesToDestroy, currCity)
		}
	}

	for j := range citiesToDestroy {
		currCity := citiesToDestroy[j]
		worldX.DestroyCity(currCity)
		aliens := currCity.Aliens
		alienNames := getNames(aliens)
		destroyStatement := currCity.Name + " has been destroyed by aliens: " + alienNames
		fmt.Println(destroyStatement)

	}
	return worldX
}

func getNames(aliens []structs.Alien) string {
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
		aliens := currCity.Aliens
		if len(aliens) > 0 {
			a := aliens[0]
			a.Travel(currCity) //todo: maybe return the updatede world.city or something like that
		}
	}
	return worldX
}
