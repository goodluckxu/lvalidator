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
        func(value interface{}) error {
            return nil
        },
    },
    "c": "string|min:2|max:4",
}, &data)
~~~

## 排序
在规则中添加sort如下
~~~
sort:1
~~~

## 重定义字段说明
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

## 所有验证规则
[valid_condition_field](#through_condition_field)
[required](#required) |
[array](#array) |
[map](#map) |
[string](#string) |
[len](#len) |
[min](#min) |
[max](#max) |
[number](#number) |
[integer](#integer) |
[bool](#bool) |
[gt](#gt) |
[gte](#gte) |
[lt](#lt) |
[lte](#lte) |
[date](#date) |
[date_gt](#date_gt) |
[date_gte](#date_gte) |
[date_lt](#date_lt) |
[date_lte](#date_lte) |
[eq_field](#eq_field) |
[email](#email) |
[phone](#phone) |
[in](#in)

#### <a id="valid_condition_field">valid_condition_field规则</a>
如果条件满足则验证后面的规则，field为传入字段，condition为条件

condition规则为 =,>,>=,<,<=,in，多个条件用&分隔，格式为= 1&< 5&in 12;15
~~~
through_condition_field:field,condition
~~~

#### <a id="required">required规则</a>
验证是否必填。null，字符串为""，数字类型为0，bool类型为false，数组为[]都不通过
~~~
required
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