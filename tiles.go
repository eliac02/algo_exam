package main

import (
	"bufio"
	"os"
)

type piastrella struct {
	x int
	y int
}

type properties struct {
	color          string
	intensity      int
	parent         piastrella
	rank           int
	blockIntensity int
}

type rule struct {
	raw     string
	ruleset map[string]int
	color   string
	usage   int
}

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
