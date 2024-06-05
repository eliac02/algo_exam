package tiles

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
		istruzione := strings.Split(scanner.Text(), " ")
		op := istruzione[0]
		switch op {
		case "C":
			x, _ := strconv.Atoi(istruzione[1])
			y, _ := strconv.Atoi(istruzione[2])
			intensity, _ := strconv.Atoi(istruzione[4])
			colora(p, x, y, istruzione[3], intensity)
		case "S":
			x, _ := strconv.Atoi(istruzione[1])
			y, _ := strconv.Atoi(istruzione[2])
			spegni(p, x, y)
		case "r":
			_, r, _ := strings.Cut(scanner.Text(), " ")
			regola(p, r)
		case "?":
			x, _ := strconv.Atoi(istruzione[1])
			y, _ := strconv.Atoi(istruzione[2])
			stato(p, x, y)
		case "s":
			stampa(p)
		case "b":
			x, _ := strconv.Atoi(istruzione[1])
			y, _ := strconv.Atoi(istruzione[2])
			blocco(p, x, y)
		case "B":
			x, _ := strconv.Atoi(istruzione[1])
			y, _ := strconv.Atoi(istruzione[2])
			bloccoOmog(p, x, y)
		case "p":
			x, _ := strconv.Atoi(istruzione[1])
			y, _ := strconv.Atoi(istruzione[2])
			propaga(p, x, y)
		case "P":
			x, _ := strconv.Atoi(istruzione[1])
			y, _ := strconv.Atoi(istruzione[2])
			propagaBlocco(p, x, y)
		case "o":
			res := make([]rule, 0)
			res = ordina(p)
			p.rules = res
		case "q":
			return
		}
	}
}
