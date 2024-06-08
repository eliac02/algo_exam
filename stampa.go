package main

import (
    "fmt"
)

func stampa(p piano) {
    fmt.Println("(")
    for _, rule := range *p.rules {
        fmt.Printf("%s: ", rule.color)
        for _, rs := range rule.ruleset {
            fmt.Printf("%d %s ", rs.count, rs.color, )
        }
        fmt.Println()
    }
    fmt.Printf(")\n")
}
