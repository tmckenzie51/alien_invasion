package main

import (
	"alien_invasion/structs"
	"testing"
)

//Test Skeleton/Driver Code
func testInvade(fileName string, numAliens int) (*structs.World, []*structs.City, int, int, int, int) {
	worldMap := read(fileName)
	world := structs.NewWorld("SymmetryCombo", worldMap)
	world = world.LaunchInvasion(numAliens)
	aftermathWorld, destroyedCities, numMoves := invade(world, numAliens)
	alienCount := aftermathWorld.NumAliens
	numTraps := numTrapCities(aftermathWorld)
	numCities := len(aftermathWorld.Cities)
	return aftermathWorld, destroyedCities, numMoves, alienCount, numTraps, numCities
}

//Test Symmetric World for invasions of sizes 0,1,3,4
func TestSymmetric(t *testing.T) {
	fileName := "./tests/symmetric.txt"
	var numberOfAliens = [...]int{0, 1, 3, 4}
	for k := range numberOfAliens {
		numAliens := numberOfAliens[k]
		aftermathWorld, destroyedCities, numMoves, alienCount, numTraps, numCities := testInvade(fileName, numAliens)
		//check for errors in city destruction
		for i := range aftermathWorld.Cities {
			for j := range destroyedCities {
				if aftermathWorld.Cities[i].Name == destroyedCities[j].Name {
					t.Error("numAliens = ", numAliens, " : destroyed city ", destroyedCities[j].Name, " should not be in final worldMap")
				}
			}
		}

		//check for errors in ending of program
		if alienCount != 0 && numMoves != numIters && numCities != numTraps {
			t.Error("numAliens = ", numAliens, " : ")
			t.Errorf("program ended prematurely. Got: alienCount = %d, numMoves = %d, numCities = %d, and numTraps = %d. Expected: alienCount = 0, or numMoves = numIters, or numCities = numTraps ", alienCount, numMoves, numCities, numTraps)
		}
	}
}

//Test Asymmetric World for invasions of sizes 0,2,3,5
func TestAsymmetric(t *testing.T) {
	fileName := "./tests/asymmetric.txt"
	var numberOfAliens = [...]int{0, 2, 3, 5}
	for k := range numberOfAliens {
		numAliens := numberOfAliens[k]
		aftermathWorld, destroyedCities, numMoves, alienCount, numTraps, numCities := testInvade(fileName, numAliens)
		//check for errors in city destruction
		for i := range aftermathWorld.Cities {
			for j := range destroyedCities {
				if aftermathWorld.Cities[i].Name == destroyedCities[j].Name {
					t.Error("destroyed city ", destroyedCities[j].Name, " should not be in final worldMap")
				}
			}
		}

		//check for errors in ending of program
		if alienCount != 0 && numMoves != numIters && numCities != numTraps {
			t.Errorf("program ended prematurely. Got: alienCount = %d, numMoves = %d, numCities = %d, and numTraps = %d. Expected: alienCount = 0, or numMoves = numIters, or numCities = numTraps ", alienCount, numMoves, numCities, numTraps)
		}
	}
}

//Test Symmetry Combination World for invasions of sizes 0,2,5,7
func TestSymmetryCombo(t *testing.T) {
	fileName := "./tests/symmetryCombo.txt"
	var numberOfAliens = [...]int{0, 2, 5, 7}
	for k := range numberOfAliens {
		numAliens := numberOfAliens[k]
		aftermathWorld, destroyedCities, numMoves, alienCount, numTraps, numCities := testInvade(fileName, numAliens)
		//check for errors in city destruction
		for i := range aftermathWorld.Cities {
			for j := range destroyedCities {
				if aftermathWorld.Cities[i].Name == destroyedCities[j].Name {
					t.Error("destroyed city ", destroyedCities[j].Name, " should not be in final worldMap")
				}
			}
		}

		//check for errors in ending of program
		if alienCount != 0 && numMoves != numIters && numCities != numTraps {
			t.Errorf("program ended prematurely. Got: alienCount = %d, numMoves = %d, numCities = %d, and numTraps = %d. Expected: alienCount = 0, or numMoves = numIters, or numCities = numTraps ", alienCount, numMoves, numCities, numTraps)
		}
	}
}

//Test One-Liner World for invasions of sizes 0,1,2
func TestOneLiner(t *testing.T) {
	fileName := "./tests/oneLiner.txt"
	var numberOfAliens = [...]int{0, 1, 2}
	for k := range numberOfAliens {
		numAliens := numberOfAliens[k]
		aftermathWorld, destroyedCities, numMoves, alienCount, numTraps, numCities := testInvade(fileName, numAliens)
		//check for errors in city destruction
		for i := range aftermathWorld.Cities {
			for j := range destroyedCities {
				if aftermathWorld.Cities[i].Name == destroyedCities[j].Name {
					t.Error("destroyed city ", destroyedCities[j].Name, " should not be in final worldMap")
				}
			}
		}

		//check for errors in ending of program
		if alienCount != 0 && numMoves != numIters && numCities != numTraps {
			t.Errorf("program ended prematurely. Got: alienCount = %d, numMoves = %d, numCities = %d, and numTraps = %d. Expected: alienCount = 0, or numMoves = numIters, or numCities = numTraps ", alienCount, numMoves, numCities, numTraps)
		}
	}
}

//Test Stress/Complex World for invasions of sizes 0,6,11,14
func TestStress(t *testing.T) {
	fileName := "./tests/stress.txt"
	var numberOfAliens = [...]int{0, 6, 11, 14}
	for k := range numberOfAliens {
		numAliens := numberOfAliens[k]
		aftermathWorld, destroyedCities, numMoves, alienCount, numTraps, numCities := testInvade(fileName, numAliens)
		//check for errors in city destruction
		for i := range aftermathWorld.Cities {
			for j := range destroyedCities {
				if aftermathWorld.Cities[i].Name == destroyedCities[j].Name {
					t.Error("destroyed city ", destroyedCities[j].Name, " should not be in final worldMap")
				}
			}
		}

		//check for errors in ending of program
		if alienCount != 0 && numMoves != numIters && numCities != numTraps {
			t.Errorf("program ended prematurely. Got: alienCount = %d, numMoves = %d, numCities = %d, and numTraps = %d. Expected: alienCount = 0, or numMoves = numIters, or numCities = numTraps ", alienCount, numMoves, numCities, numTraps)
		}
	}
}
