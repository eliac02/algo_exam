// Elia Cortesi 01911A
package main

import (
	"bufio"
	"os"
)

// piastrella represents a tile of the system
//
// x y The coordinates of the tile
type piastrella struct {
	x int
	y int
}

// properties represents the properties of a tile
//
// color The color of the tile
// intensity The intensity of the color
// parent The parent the tile is referring to
// rank The number of tiles that refers to the tile
// blockIntensity The sum of the intensity of the tiles that refers to the tile
type properties struct {
	color          string
	intensity      int
	parent         piastrella
	rank           int
	blockIntensity int
}

// ruleset represents the required tiles that need to be adjacents to the tile
//
// color The color of the adjacent tile
// count The number of the adjacents tile of color @color
type ruleset struct {
	color string
	count int
}

// rule represents a rule of the system
//
// raw The raw command given to the program
// ruleset The list of required adjacent tiles
// color The color that the rule gives to teh tile
// usage The number of times the rule has been used
type rule struct {
	raw     string
	ruleset []ruleset
	color   string
	usage   int
}

// piano The system tiles-rules
//
// tiles The set of all the tiles of the system
// rules The list of all the rules of the systems
type piano struct {
	tiles map[piastrella]*properties
	rules *[]rule
}

func main() {
	p := makeSet()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		esegui(p, scanner.Text())
	}
}
