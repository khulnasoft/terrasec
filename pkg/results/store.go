

package results

// NewViolationStore returns a new violation store
func NewViolationStore() *ViolationStore {
	return &ViolationStore{
		Violations:        []*Violation{},
		SkippedViolations: []*Violation{},
		PassedRules:       []*PassedRule{},
	}
}

// AddResult Adds individual violations into the violation store
// when skip is true, violations are added to skipped violations
func (s *ViolationStore) AddResult(violation *Violation, isSkipped bool) {
	if isSkipped {
		s.SkippedViolations = append(s.SkippedViolations, violation)
	} else {
		s.Violations = append(s.Violations, violation)
	}
}

// GetResults Retrieves all violations from the violation store
// when skip is true, it returns only the skipped violations
func (s *ViolationStore) GetResults(isSkipped bool) []*Violation {
	if isSkipped {
		return s.SkippedViolations
	}
	return s.Violations
}

// AddPassedRule Adds individual passed rule into the violation store
func (s *ViolationStore) AddPassedRule(rule *PassedRule) {
	s.PassedRules = append(s.PassedRules, rule)
}

// GetPassedRules Retrieves all passed rules from the violation store
func (s *ViolationStore) GetPassedRules() []*PassedRule {
	return s.PassedRules
}
