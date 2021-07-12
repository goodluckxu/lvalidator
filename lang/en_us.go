package lang

type EnUs struct {
	Lang
}

func (z *EnUs) Init() {
	z.Required = "Field {ruleKey} is required"
	z.Number = "Field {ruleKey} is a number"
	z.Integer = "Field {ruleKey} is a integer"
	z.Gt = "Field {ruleKey} must be greater than {ruleValue}"
	z.Gte = "Field {ruleKey} must be greater than or equal {ruleValue}"
}
