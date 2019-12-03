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
	worldMap := read("./tests/X.txt") //todo: check if this works
	numAliens, _ := strconv.Atoi(os.Args[0])
	worldX := structs.NewWorld("X", worldMap)
	worldX.LaunchInvasion(numAliens)
	for i := 0; i < 10000; i++ { // todo(REFINE): save num_iters as a const
		cities := worldX.Cities
		fightAndDestroy(worldX)
		numAliens = worldX.NumAliens
		if numAliens <= 0 {
			worldX.PrintMap()
			return
		} else {
			wander(cities)
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

func fightAndDestroy(worldX structs.World) {
	cities := worldX.Cities
	for i := range cities {
		currCity := cities[i]
		numAliens := currCity.AlienCount
		aliens := currCity.Aliens
		if numAliens >= 2 {
			worldX.DestroyCity(currCity)
			alienNames := getNames(aliens)
			fmt.Printf(currCity.Name, " has been destroyed by aliens: %d\n", alienNames)
		}
	}
}

func getNames(aliens []structs.Alien) string {
	alienNames := ""
	for i := range aliens {
		currAlien := aliens[i]
		if i == len(aliens)-1 {
			alienNames += strconv.Itoa(currAlien.Id)
		} else if i == len(aliens)-2 {
			alienNames += strconv.Itoa(currAlien.Id) + "and "
		} else {
			alienNames += strconv.Itoa(currAlien.Id) + ", "
		}
	}
	return alienNames
}

func wander(cities []structs.City) {
	for i := range cities {
		currCity := cities[i]
		aliens := currCity.Aliens
		a := aliens[0]
		a.Travel(currCity)
	}
}
