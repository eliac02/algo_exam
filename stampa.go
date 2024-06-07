package main

import (
    "fmt"
)

func stampa(p piano) {
    fmt.Println("(")
    for _, rule := range *p.rules {
        fmt.Printf("%s: ", rule.color)
        for key, value := range rule.ruleset {
            fmt.Printf("%d %s ", value, key, )
        }
        fmt.Println()
    }
    fmt.Printf(")\n")
}
