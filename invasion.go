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

func fightAndDestroy() {
	cities = world_x.cities
	for i := range cities {
		curr_city = cities[i]
		num_aliens = curr_city.alien_count
		aliens = curr_city.aliens()
		if num_aliens >= 2 {
			world_x.destroy(curr_city)
			curr_city.fight()
		}
		fmt.Printf("%d has been destroyed by aliens %d\n", curr_city, aliens)
	}
}

func wander() {
	cities = world_x.cities
	for i, curr_city := range cities {
		a = curr_city.aliens
		a[0].travel()
	}
}

func read(fileName string) *map[string]string {
	worldMap := make(map[string]string)
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	file.Close()

	for _, eachline := range txtlines {
		cityInfo = strings.Split(eachline, " ")
		for i := range cityInfo {
			if i == 0 {
				cityName = cityInfo[i]
			} else {
				niehgborInfo = strings.Split(cityInfo[i], "=")
				worldMap[cityName] = neighborInfo
			}
		}
	}
	return &worldMap
}

func main() {
	worldMap := read("/X.txt")
	numAliens, _ := strconv.Atoi(os.Args[3])
	worldX := structs.World.NewWorld('X', worldMap)
	structs.World.worldX.LaunchInvasion(numAliens) //todo: not sure what's the right way to go about this
	for i := range 10000 {
		fightAndDestroy()
		numAliens = worldX.AlienCount
		if numAliens <= 0 {
			worldX.PrintMap()
			return
		} else {
			wander()
		}
	}
}
