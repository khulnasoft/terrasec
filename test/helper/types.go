

package helper

import (
	"github.com/khulnasoft/terrasec/pkg/results"
)

type violations []*results.Violation
type passedRules []*results.PassedRule

// sort for violations
func (v violations) Len() int {
	return len(v)
}

func (v violations) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

func (v violations) Less(i, j int) bool {
	if v[i].File < v[j].File {
		return true
	}
	if v[i].File > v[j].File {
		return false
	}

	if v[i].ResourceType < v[j].ResourceType {
		return true
	}

	if v[i].ResourceType > v[j].ResourceType {
		return false
	}

	if v[i].RuleName < v[j].RuleName {
		return true
	}

	if v[i].RuleName > v[j].RuleName {
		return false
	}

	if v[i].ResourceName < v[j].ResourceName {
		return true
	}

	if v[i].ResourceName > v[j].ResourceName {
		return false
	}

	if v[i].LineNumber < v[j].LineNumber {
		return true
	}

	return v[i].LineNumber > v[j].LineNumber
}

// sort for passed rules
func (p passedRules) Len() int {
	return len(p)
}

func (p passedRules) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p passedRules) Less(i, j int) bool {

	if p[i].RuleName < p[j].RuleName {
		return true
	}

	if p[i].RuleName > p[j].RuleName {
		return false
	}

	if p[i].RuleID < p[j].RuleID {
		return true
	}

	return p[i].RuleID > p[j].RuleID
}
