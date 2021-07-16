package lang

type ZhCn struct {
	Lang
}

func (z *ZhCn) Init() {
	z.Error = "验证值错误。规则：{rule}:{ruleValue}，错误：{error}"
	z.Required = "{ruleKey}为必填项"
	z.Array = "{ruleKey}必须是数组"
	z.Map = "{ruleKey}必须是对象"
	z.String = "{ruleKey}必须是字符串"
	z.Len = "{ruleKey}长度必须是{ruleValue}"
	z.Min = "{ruleKey}最小长度为{ruleValue}"
	z.Max = "{ruleKey}最大长度为{ruleValue}"
	z.Number = "{ruleKey}必须是数字"
	z.Integer = "{ruleKey}必须是整数"
	z.Gt = "{ruleKey}必须大于{ruleValue}"
	z.Gte = "{ruleKey}必须大于等于{ruleValue}"
	z.Lt = "{ruleKey}必须小于{ruleValue}"
	z.Lte = "{ruleKey}必须小于等于{ruleValue}"
	z.Date = "{ruleKey}必须是日期格式"
	z.DateGt = "{ruleKey}必须大于{ruleValue}"
	z.DateGte = "{ruleKey}必须大于等于{ruleValue}"
	z.DateLt = "{ruleKey}必须小于{ruleValue}"
	z.DateLte = "{ruleKey}必须小于等于{ruleValue}"
	z.EqField = "{ruleKey}必须等于{ruleValue}"
	z.Email = "{ruleKey}必须是邮箱"
	z.Phone = "{ruleKey}必须是手机号"
}
