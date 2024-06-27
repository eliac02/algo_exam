//Elia Cortesi 01911A

package main

import (
	"strconv"
	"strings"
)

// esegui executes the command that's been given
//
// @param p The system tiles-rules
// @param s The command given to the program
func esegui(p piano, s string) {
	tempSlice := strings.Split(s, " ")
	command := tempSlice[0]
	switch command {
	case "s":
		stampa(p)
	case "q":
		return
	case "C":
		x, err := strconv.Atoi(tempSlice[1])
		if err != nil {
			return
		}
		y, err := strconv.Atoi(tempSlice[2])
		if err != nil {
			return
		}
		i, err := strconv.Atoi(tempSlice[4])
		if err != nil {
			return
		}
		colora(p, x, y, tempSlice[3], i)
    case "L":
        x1, err := strconv.Atoi(tempSlice[1])
		if err != nil {
			return
		}
		y1, err := strconv.Atoi(tempSlice[2])
		if err != nil {
			return
		}
		x2, err := strconv.Atoi(tempSlice[3])
		if err != nil {
			return
		}
        y2, err := strconv.Atoi(tempSlice[4])
        lung(p, x1, y1, x2, y2)
	case "S", "?", "b", "B", "p", "P", "t":
		x, err := strconv.Atoi(tempSlice[1])
		if err != nil {
			return
		}
		y, err := strconv.Atoi(tempSlice[2])
		if err != nil {
			return
		}
		switch command {
		case "S":
			spegni(p, x, y)
		case "?":
			stato(p, x, y)
		case "b":
			blocco(p, x, y)
		case "B":
			bloccoOmog(p, x, y)
		case "p":
			propaga(p, x, y)
		case "P":
			propagaBlocco(p, x, y)
        case "t":
            pista(p, x, y, s)
		}
	case "r":
		regola(p, s)
	case "o":
		ordina(p)
	}
}
