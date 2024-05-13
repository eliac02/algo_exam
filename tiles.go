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

type pair struct {
	number int
	letter string
}

type rule struct {
	ruleset map[string]int
	color   string
}

type piano struct {
	tiles map[piastrella][2]string
	rules []rule // uso un array per sfruttare un algoritmo di sorting
	adj   []piastrella
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	p := piano{tiles: make(map[piastrella][2]string), rules: make([]rule, 0), adj: make([]piastrella, 0)}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		istruzione := strings.Split(scanner.Text(), " ")
		op := istruzione[0]
		switch op {
		case "C":
			x, _ := strconv.Atoi(istruzione[1])
			y, _ := strconv.Atoi(istruzione[2])
			colora(p, x, y, istruzione[3], 0)
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
		case "B":
		case "p":
		case "P":
		case "o":
		case "q":
			return
		}
	}
}
