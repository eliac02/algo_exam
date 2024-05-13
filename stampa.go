package tiles

import (
    "fmt"
)

func stampa(p piano) {
    fmt.Println("(")
    for i:=0; i<len(p.rules); i++ {
        fmt.Print(p.rules[i].color)
        for key, value := range p.rules[i].ruleset {
            fmt.Print(value, key)
        }
        fmt.Println()
    }
    fmt.Printf(")\n")
}
