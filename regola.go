package tiles

import (
	"strconv"
	"strings"
)

func regola(p piano, r string) {
	rulesSplitted := strings.Split(r, " ")
	var newRule rule
	newRule.raw = r
	newRule.color = rulesSplitted[0]
	newRule.ruleset = make(map[string]int)
	for i := 1; i <= (len(rulesSplitted)-1)/2; i++ {
		num, _ := strconv.Atoi(rulesSplitted[i*2-1])
		newRule.ruleset[rulesSplitted[i*2]] = num
	}
    newRule.usage = 0
	p.rules = append(p.rules, newRule)
}
