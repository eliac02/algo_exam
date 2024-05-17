package tiles

import (
    "strings"
    "strconv"
)

func regola(p piano, r string) {
    listRules := strings.Split(r, " ")
    var newRule rule
    newRule.raw = r
    newRule.color = listRules[0]
    newRule.ruleset = make(map[string]int)
    for i := 1; i <= (len(listRules)-1)/2; i++ {
        num, _ := strconv.Atoi(listRules[i*2-1])
        newRule.ruleset[listRules[i*2]] = num
    }
    p.rules = append(p.rules, newRule)
}
