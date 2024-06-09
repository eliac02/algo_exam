//Elia Cortesi 01911A

package main

import (
	"strconv"
	"strings"
)

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
	case "S", "?", "b", "B", "p", "P":
		if len(tempSlice) != 3 {
			return
		}
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
		}
	case "r":
		regola(p, s)
	case "o":
		ordina(p)
	}
}
