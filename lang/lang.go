package lang

type Lang struct {
	Required string // 必填
	String   string // 字符串
	Number   string // 数字
	Integer  string // 整数
	Gt       string // 大于
	Gte      string // 大于等于
	Lt       string // 小于
	Lte      string // 小于等于
	Eq       string // 等于
	Date     string // 日期 Y-m-d H:i:s
}
