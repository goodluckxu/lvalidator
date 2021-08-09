package lang

type Lang struct {
	Error      string // 验证错误
	Required   string // 必填
	Array      string // 数组
	Map        string // 对象
	String     string // 字符串
	Len        string // 长度
	Min        string // 最小长度
	Max        string // 最大长度
	Number     string // 数字
	Integer    string // 整数
	Bool       string // 布尔类型
	Gt         string // 大于
	Gte        string // 大于等于
	Lt         string // 小于
	Lte        string // 小于等于
	Date       string // 日期 Y-m-d H:i:s
	DateFormat string // 日期自定义格式 Y-m-d H:i:s
	DateGt     string // 日期大于值
	DateGte    string // 日期大于等于值
	DateLt     string // 日期小于值
	DateLte    string // 日期小于等于值
	EqField    string // 等于字段
	Email      string // 邮箱
	Phone      string // 手机
	In         string // 是否在数组里面
	Unique     string // 数组内的值唯一
	Regexp     string // 正则验证
}
