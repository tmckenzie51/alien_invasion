package main //todo: check if this actually works

import (
	"alien_invasion/structs"
	"testing"
)

var fileNames = [...]string{"./tests/symmetric.txt", "./tests/asymmetric.txt", "./tests/symmetryCombo.txt", "./tests/stress.txt", "./tests/oneLiner.txt"}

func TestSymmetryCombo(t *testing.T) {
	fileName := "./tests/symmetryCombo.txt"
	worldMap := read(fileName)
	world := structs.NewWorld("SymmetryCombo", worldMap)
	numAliens := 7
	world = world.LaunchInvasion(numAliens)
	world = invade(world, numAliens)

	//got := Abs(-1)
	if got != 1 {
		t.Errorf("Abs(-1) = %d; want 1", got)
	}
}

//1. save fileNames in main
//2. run several unit tests, namely:
//get worldmap from creating newworld
// get stateOfWorld after Launch invasion (with preferred numAliens) --> get aliens in each city & numAliens in each city
//get destroyed cities
//get killed aliens from :  destruction statements, for alien tracking
//check for errrors?
