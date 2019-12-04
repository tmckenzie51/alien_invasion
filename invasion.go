//Assumption 1:
//For each direction (north, south,east, west), there will be at most 1 city in that direction.
//in other words, there can not be a road in any direction that leads to multiple cities
//in terms of our input, this means that "Foo north=Bar,Bee" is not a valid input

//Assumption 2:
//all inputted world_maps meet the aforementioned standards/requirements of a the world, such as no multiple destinations from the same road,
//cities names are all string and do not contain numbers (int or float), and the syntactical requirements of the input is met
//as specified in the problem specifics detailed in the alien_invasion pdf document.

//Assumption 3:
//All inputs make geographical sense, meaning that, for example: there will be no input such that Foo.North = Bar, and Bar.North = Foo.

//Assumption 4:
//No need to filter for invalid input. We will assume that all input is valid in terms or character/string capitalization, white spaces, equal signs, etc.
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

const numIters = 10000

//Driver Code
func main() {
	worldMap := read("./tests/symmetryCombo.txt")
	numAliens, _ := strconv.Atoi(os.Args[1])
	worldX := structs.NewWorld("SymmetryCombo", worldMap)
	worldX = worldX.LaunchInvasion(numAliens)
	invade(worldX, numAliens)
}

//after initial launch/alien deployment into the cities, we now carry on with the invasion, by destroying the appropriate cities,
//accounting for alien deaths in fights, and do the necessary city, etc. assignments when aliens move/wander around.
func invade(worldX *structs.World, numAliens int) (*structs.World, []*structs.City, int) {
	var destroyedCities []*structs.City
	numMoves := 0
	for i := 0; i <= numIters; i++ {
		var destroyed []*structs.City
		worldX, destroyed = fightAndDestroy(worldX)
		for j := range destroyed {
			destroyedCities = append(destroyedCities, destroyed[j])
		}
		numAliens = worldX.NumAliens
		numTraps := numTrapCities(worldX)
		//trap city optimization: when all the cities lef tin the world are trap cities we can terminate the program early, because all surviving aliens are trapped
		//in their respective cities
		if numAliens == 0 || i == numIters || numTraps == len(worldX.Cities) {
			worldX.PrintMap()
			return worldX, destroyedCities, numMoves
		} else {
			worldX.Cities = wander(worldX)
			numMoves += 1
		}
	}
	return worldX, destroyedCities, numMoves
}

//Get the number of cities that do not have any bridges leading out of them. These are called 'trap cities'.
func numTrapCities(worldX *structs.World) int {
	numTraps := 0
	for i := range worldX.Cities {
		city := worldX.Cities[i]
		if len(city.Directions) == 0 {
			numTraps += 1
		}
	}
	return numTraps
}

//Read in the input world files
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

	file.Close()

	for _, eachLine := range txtLines {
		cityInfo := strings.Split(eachLine, " ")
		cityName := cityInfo[0]
		for i := range cityInfo {
			if i > 0 {
				neighborInfo := strings.Split(cityInfo[i], "=")
				worldMap[cityName] = append(worldMap[cityName], neighborInfo)
			}
		}
	}
	return worldMap
}

//Scan the existing cities to see which have more than 1 alien, and destroy these cities and all bridges leading to them.
//Alien count is adjusted when cities are destroyed as they kill each other in the process.
func fightAndDestroy(worldX *structs.World) (*structs.World, []*structs.City) {
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
		aliens := currCity.Aliens
		alienNames := getNames(aliens)
		destroyStatement := currCity.Name + " has been destroyed by aliens: " + alienNames
		fmt.Println(destroyStatement)
	}
	return worldX, citiesToDestroy
}

//Get the ID of the aliens remaining in a specific city
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

//Aliens roam around randomly, moving from city to city
func wander(worldX *structs.World) []*structs.City {
	cities := worldX.Cities
	for i := range cities {
		currCity := cities[i]
		aliens := currCity.Aliens
		if len(aliens) > 0 && len(currCity.Directions) > 0 {
			a := aliens[0]
			a.Travel(currCity)
		}
	}
	return cities
}
