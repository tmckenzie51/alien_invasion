# Alien Invasion
****What is this ?**** 

This is my solution to the Tendermint Alien Invasion Technical Challenge, briefly described here:

*"Mad aliens are about to invade the earth and you are tasked with simulating the
  invasion. These aliens start out at random places on the map, and wander around randomly,
 following links. Each iteration, the aliens can travel in any of the directions
 leading out of a city. When two aliens end up in the same place, they fight, and in the process kill
 each other and destroy the city. When a city is destroyed, it is removed from
 the map, and so are any roads that lead into or out of it. Once a city is destroyed, aliens can no longer travel to or through it. You should create a program that reads in the world map, creates N aliens, and
 unleashes them. The program should run until all the aliens have been destroyed, or each alien has moved at least 10,000 times."*
   

****Input****           
The city and each of the pairs are separated by a single space, and the
directions are separated from their respective cities with an equals (=) sign.    
For example:

Foo north=Bar west=Baz south=Qu-ux                                                                                                                       
Bar south=Foo west=Bee                                                                                      
                                                                                          
                                                                                                                                                           
****Architecture / Design****
 
 This program consists of 3 structs namely: city,world, and alien structs, each of which has methods and fields that are used/referenced in the main invasion.go program file.
 
 ****Tests****
 
Test files can be found in invasion_test.go and in the tests folder. The tests folder contains all of the made-up test worlds, namely:
1. *Symmetric World* :  
In this world, all bridges are bidirectional.
2. *Asymmetric World*    
In this world, there are no bidirectional bridges. 
3. *Symmetry Combination World*   
In this world there is a combination of bidirectional and single directional bridges. 
4. *One-Liner World*
In this world there is only one "non-trap" city (A trap city is one which has no bridges leading out of the city)
5. *Stress World*:   
In this world there is a combination of single directional and bidirectional birdges on larger number of cities in a more complex arrangement. 

Each of these worlds are tested with launch parties of varying sizes listed below:
1. 0
2. a size less than the number of cities in the world to be tested
3. the same number of aliens as there are cities in the world to be tested
4. more aliens than there are cities in teh world to be tested. 

 ****How To Use****
 1. Fork/Clone Repository
 2. To run tests, run the following on the command line:   
 *test go -run NameOfTest*
 3. To run program, run the following on the command line:   
 *go run invasion.go << NumberOfAliens >>*  
 where  *<< NumberOfAliens >>* is your desired alien launch party size. The default world is the Symmetry Combination World. This may be modified by changing the input filepath in the read function in invasion.go 
  