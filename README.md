# lvalidator
golang 中使用 *http.Request 验证数据

目前支持解析ValidJson,ValidXml

## 用法
验证规则可以使用 | 分隔，也可以定义成 []interface{} 类型
~~~go
var data interface{}
valid := lvalidator.New(c.Request)
err := valid.ValidJson(map[string]interface{}{
    "a": "sort:1|notes:测试|required|number|integer|gt:1|gte:3|lt:15|lte:10",
    "b": []interface{}{
        "sort:2",
        "notes:飞机",
        "date",
        func(validData  interface{}) error {
            return nil
        },
    },
    "c": "string|min:2|max:4",
}, &data)
~~~
或者
~~~go
var data interface{}
valid := lvalidator.New(c.Request)
err := valid.ValidJson([]lvalidator.RuleRow{
    {Key: "a", Rules: "required|string", Notes: "测试"},
    {Key: "b", Rules: []interface{}{
        "date",
        func(validData  interface{}) error {
            return nil
        },
    }, Notes: "你好"}
}, &data)
~~~

## 排序(只适用于map[string]interface{}类型的规格)
在规则中添加sort如下
~~~
sort:1
~~~

## 重定义字段说明(只适用于map[string]interface{}类型的规格)
在规则中添加notes如下
~~~
notes:注释
~~~

## 自定义验证
~~~go
// validData 需要验证的值
// validNotes 字段注释(非必填)
func(validData interface{},validNotes string) error {
    // 逻辑代码
    return nil
}
~~~

## 所有验证类型

### 通用验证
[valid_condition](#valid_condition) |
[required](#required) |
[nullable](#nullable) |
[len](#len) |
[min](#min) |
[max](#max) |
[date](#date) |
[date_format](#date_format) |
[email](#email) |
[phone](#phone) |
[in](#in) |
[unique](#unique) |
[regexp](#regexp)
### 类型验证
[array](#array) |
[map](#map) |
[string](#string) |
[number](#number) |
[integer](#integer) |
[bool](#bool)
### 比较验证
[gt](#gt) |
[gte](#gte) |
[lt](#lt) |
[lte](#lte) |
[date_gt](#date_gt) |
[date_gte](#date_gte) |
[date_lt](#date_lt) |
[date_lte](#date_lte)
### 字段验证
[valid_condition_field](#valid_condition_field) |
[eq_field](#eq_field) |
[gt_field](#gt_field) |
[gte_field](#gte_field) |
[lt_field](#lt_field) |
[lte_field](#lte_field) |
[date_gt_field](#date_gt_field) |
[date_gte_field](#date_gte_field) |
[date_lt_field](#date_lt_field) |
[date_lte_field](#date_lte_field)

## 所有验证规则

#### <a id="valid_condition_field">valid_condition_field规则</a>
如果条件满足则验证后面的规则，field为传入字段，condition为条件

condition规则为 =,>,>=,<,<=,in，多个条件用&分隔，格式为= 1&< 5&in 12;15
~~~
valid_condition_field:field,condition
~~~

#### <a id="valid_condition">valid_condition规则</a>
如果条件满足则验证后面的规则，result为结果，有true,false,1,0四种，其在true,1验证后面数据
~~~
valid_condition:result
~~~

#### <a id="required">required规则</a>
验证是否必填。null，字符串为""，数字类型为0，bool类型为false，数组为[]，map为{}都不通过
参数为string,number,bool,array,map，代表string为空不验证，number为0不验证，bool为false不验证,array为[]不验证,map为{}不验证
~~~
required:number
~~~

#### <a id="nullable">nullable规则</a>
数据为空则不验证后面的规则，字符串为""，数字类型为0，bool类型为false，数组为[]，map为{}都不验证后的单规则
~~~
nullable
~~~

#### <a id="array">array规则</a>
验证是否是数组
~~~
array
~~~

#### <a id="map">map规则</a>
验证是否是对象
~~~
map
~~~

#### <a id="string">string规则</a>
验证是否是字符串
~~~
string
~~~

#### <a id="len">len规则</a>
验证是长度等于某个数。可获取字符串和数组长度
~~~
len:5
~~~

#### <a id="min">min规则</a>
验证是长度大于等于某个数。可获取字符串和数组长度
~~~
min:5
~~~

#### <a id="max">max规则</a>
验证是长度小于等于某个数。可获取字符串和数组长度
~~~
max:5
~~~

#### <a id="number">number规则</a>
验证是否是数字。可验证数字和字符串的数字
~~~
number
~~~

#### <a id="integer">integer规则</a>
验证是否是整数。可验证数字和字符串的数字
~~~
integer
~~~

#### <a id="bool">bool规则</a>
验证是否是布尔类型。可验证整数和布尔类型
~~~
bool
~~~

#### <a id="gt">gt规则</a>
验证是否大于某个数。可验证数字和字符串的数字
~~~
gt:10
~~~

#### <a id="gte">gte规则</a>
验证是否大于等于某个数。可验证数字和字符串的数字
~~~
gte:10
~~~

#### <a id="lt">lt规则</a>
验证是否小于某个数。可验证数字和字符串的数字
~~~
lt:10
~~~

#### <a id="lte">lte规则</a>
验证是否小于等于某个数。可验证数字和字符串的数字
~~~
lte:10
~~~

#### <a id="date">date规则</a>
验证是否是日期格式 Y-m-d H:i:s类型
~~~
date
~~~

#### <a id="date_format">date_format规则</a>
自定义验证是否是日期格式 Y-m-d H:i:s类型，自定义类型 Y年，m月，d日，H时，i分，s秒
~~~
date_format:Y-m-d H:i:s
~~~

#### <a id="date_gt">date_gt规则</a>
验证日期是否大于某个值
~~~
date_gt:2002-02-05
~~~

#### <a id="date_gte">date_gte规则</a>
验证日期是否大于等于某个值
~~~
date_gte:2002-02-05
~~~

#### <a id="date_lt">date_lt规则</a>
验证日期是否小于某个值
~~~
date_lt:2002-02-05
~~~

#### <a id="date_lte">date_lte规则</a>
验证日期是否小于等于某个值
~~~
date_lte:2002-02-05
~~~

#### <a id="eq_field">eq_field规则</a>
验证两个字段是否相同，field为传的其他字段
~~~
eq_field:field
~~~

#### <a id="email">email规则</a>
验证是否是邮箱
~~~
email
~~~

#### <a id="phone">phone规则</a>
验证是否是手机号
~~~
phone
~~~

#### <a id="in">in规则</a>
验证是否在数组里面
~~~
in:1,2,3
~~~

#### <a id="unique">unique规则</a>
验证数组内的值唯一
~~~
unique
~~~

#### <a id="regexp">regexp规则</a>
验证正则表达式
~~~
regexp:^\d$
~~~

#### <a id="gt_field">gt_field规则</a>
验证是否大于某个字段。可验证数字和字符串的数字
~~~
gt_field:field
~~~

#### <a id="gte_field">gte_field规则</a>
验证是否大于等于某个字段。可验证数字和字符串的数字
~~~
gte_field:field
~~~

#### <a id="lt_field">lt_field规则</a>
验证是否小于某个字段。可验证数字和字符串的数字
~~~
lt_field:field
~~~

#### <a id="lte_field">lte_field规则</a>
验证是否小于等于某个字段。可验证数字和字符串的数字
~~~
lte_field:field
~~~

#### <a id="date_gt_field">date_gt_field规则</a>
验证日期是否大于某个字段
~~~
date_gt_field:field
~~~

#### <a id="date_gte_field">date_gte_field规则</a>
验证日期是否大于等于某个字段
~~~
date_gte_field:field
~~~

#### <a id="date_lt_field">date_lt_field规则</a>
验证日期是否小于某个字段
~~~
date_lt_field:field
~~~

#### <a id="date_lte_field">date_lte_field规则</a>
验证日期是否小于等于某个字段
~~~
date_lte_field:field
~~~
