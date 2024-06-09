//Elia Cortesi 01911A

package main

import (
    "fmt"
)

// stampa prints the list of the rules of the system
//
// @param p The system tiles-rules
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
