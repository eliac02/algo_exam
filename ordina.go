package tiles

func ordina(p piano) []rule {
    p.rules = mergeSort(p.rules)
    return p.rules
}
