package main

func getAdiacenti(x, y int) [8]piastrella {
    adjs := [8]piastrella{
        {x, y+1}, {x-1, y+1}, {x-1, y}, {x-1, y-1}, {x, y-1}, {x+1, y-1}, {x+1, y}, {x+1, y+1},
    }
    return adjs
}

func ruleOk(reg rule, colorCount map[string]int) bool {
    for color, count := range reg.ruleset {
        if colorCount[color] < count {
            return false
        }
    }
    return true
}

func esploraBlocco(p piano, start piastrella, seen map[piastrella]bool) map[piastrella]*properties {
    root := p.Find(start)
    block := make(map[piastrella]*properties)
    pila := []piastrella{root}
    seen[root] = true

    for len(pila) > 0 {
        current := pila[len(pila)-1]
        pila = pila[:len(pila)-1]
        block[current] = p.tiles[current]

        adiacenti := getAdiacenti(current.x, current.y)
        for _, adj := range adiacenti {
            if _, exists := p.tiles[adj]; exists && !seen[adj] {
                if p.Find(adj) == root {
                    seen[adj] = true
                    pila = append(pila, adj)
                }
            }
        }
    }
return block
}
