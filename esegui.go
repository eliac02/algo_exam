package tiles

import (
	"strconv"
	"strings"
)

func esegui(p piano, s string) {
	tempSlice := strings.Split(s, " ")
    x, _ := strconv.Atoi(tempSlice[1])
    y, _ := strconv.Atoi(tempSlice[2])
    switch tempSlice[0] {
	case "C":
        i, _ := strconv.Atoi(tempSlice[4])
        colora(p, x, y, tempSlice[3], i)
	case "S":
        spegni(p, x, y)
	case "r":
        regola(p, s)
	case "?":
        _, _ = stato(p, x, y)
	case "s":
        stampa(p)
	case "b":
        blocco(p, x, y)
	case "B":
        bloccoOmog(p, x, y)
	case "p":
        propaga(p, x, y)
	case "P":
        propagaBlocco(p, x, y)
	case "o":
        ordina(p)
	case "q":
        return
	}
}
