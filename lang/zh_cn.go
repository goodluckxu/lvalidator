package lang

type ZhCn struct {
	Lang
}

func (z *ZhCn) Init() {
	z.Required = "{ruleKey}为必填项"
	z.String = "{ruleKey}必须是字符串"
	z.Number = "{ruleKey}必须是数字"
	z.Integer = "{ruleKey}必须是整数"
	z.Gt = "{ruleKey}必须大于{ruleValue}"
	z.Gte = "{ruleKey}必须大于等于{ruleValue}"
	z.Lt = "{ruleKey}必须小于{ruleValue}"
	z.Lte = "{ruleKey}必须小于等于{ruleValue}"
	z.Date = "{ruleKey}必须是日期格式"
}
