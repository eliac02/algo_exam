package tiles

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type piastrella struct {
	x      int
	y      int
}

type rule struct {
	rule    string
	consumo int
}

type piano struct {
	tiles map[piastrella][2]string
	rules []rule
    adj []piastrella
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
			colora(p, istruzione[1], istruzione[2], istruzione[3])
		case "S":
			spegni(p, istruzione[1], istruzione[2])
		case "r":
			regola(p, r)
		case "?":
			stato(p, istruzione[1], istruzione[2])
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
