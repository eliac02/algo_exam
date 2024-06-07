package main

import (
	"strconv"
	"strings"
)

func regola(p piano, r string) {
	rulesSplitted := strings.Split(r, " ")
	var newRule rule
	newRule.raw = r
	newRule.color = rulesSplitted[1]
	newRule.ruleset = make(map[string]int)
	for i := 2; i <= (len(rulesSplitted) - 2); i += 2 {
		num, err := strconv.Atoi(rulesSplitted[i])
		if err != nil {
			return
		}
		newRule.ruleset[rulesSplitted[i+1]] = num
	}
    //usg, _ := strconv.Atoi(rulesSplitted[len(rulesSplitted)-1])
	//newRule.usage = usg
	newRule.usage = 0
    *p.rules = append(*p.rules, newRule)
}
