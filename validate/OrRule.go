package validate

// OrRule validation rule to combine rules with OR logical operation.
// When one of rules returns no errors, then this rule also returns no errors.
// When all rules return errors, then the rule returns all errors.
//	see IValidationRule
//	Example:
//		var schema = NewSchema()
//			.WithRule(NewOrRule(
//				NewValueComparisonRule("LT", 1),
//				NewValueComparisonRule("GT", 10),
//			));
//		schema.Validate();          // Result: no error
//		schema.Validate5);          // Result: 5 must be less than 1 or 5 must be more than 10
//		schema.Validate(20);        // Result: no error
type OrRule struct {
	rules []IValidationRule
}

// NewOrRule creates a new validation rule and ses its values
//	Parameters: rule IValidationRule a rule to be negated.
//	Returns: *OrRule
func NewOrRule(rules ...IValidationRule) *OrRule {
	return &OrRule{
		rules: rules,
	}
}

// Validate validates a given value against this rule.
//	Parameters:
//		- path string a dot notation path to th value.
//		- schema  ISchema a schema this rule is called from
//		- value any a value to be validated.
//	Returns: []*ValidationResult a list with validation results to add new results.
func (c *OrRule) Validate(path string, schema ISchema, value any) []*ValidationResult {
	if len(c.rules) == 0 {
		return nil
	}

	results := make([]*ValidationResult, 0)

	for _, rule := range c.rules {
		ruleResults := rule.Validate(path, schema, value)

		if len(ruleResults) == 0 {
			return nil
		}

		results = append(results, ruleResults...)
	}

	return results
}
