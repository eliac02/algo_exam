package tiles

import (
	"bufio"
	"fmt"
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
	tiles map[piastrella]*properties // insieme delle piastrelle (uso puntatore a properties per questioni di addressability)
	rules []rule                     // uso un array per sfruttare un algoritmo di sorting
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	p := makeSet()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		esegui(p, scanner.Text())
	}
}
