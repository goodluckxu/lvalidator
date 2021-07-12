# lvalidator
golang 中使用 *http.Request 验证数据

目前支持解析ValidJson,ValidXml

## 用法
验证规则可以使用英文状态的分号 ; 分隔，也可以定义成 []interface{} 类型
~~~go
var data interface{}
valid := lvalidator.New(c.Request)
err := valid.ValidJson(map[string]interface{}{
    "a": "sort:1;notes:测试;required;number;integer;gt:1;gte:3;lt:15;lte:10",
    "b": []interface{}{
        "sort:2",
        "notes:飞机",
        "date",
        func(value interface{}) error {
            return errors.New("验证失败")
        },
    },
    "c": "",
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
func(value interface{}) error {
    // 逻辑代码
    return nil
}
~~~

## 所有验证规则
[required](#required)必填
[string](#string)字符串
[number](#number)数字
[integer](#integer)整数
[gt](#gt)大于
[gte](#gte)大于等于
[lt](#lt)小于
[lte](#lte)小于等于
[date](#date)日期

#### <a id="required">required规则</a>
验证是否必填。null，字符串为""，数字类型为0，bool类型为false，数组为[]都不通过
~~~
required
~~~

#### <a id="string">string规则</a>
验证是否是字符串
~~~
string
~~~

#### <a id="number">number规则</a>
验证是否是数字
~~~
number
~~~

### <a id="integer">integer规则</a>
验证是否是整数
~~~
integer
~~~

### <a id="gt">gt规则</a>
验证是否大于某个数
~~~
gt:10
~~~

### <a id="gte">gte规则</a>
验证是否大于等于某个数
~~~
gte:10
~~~

### <a id="lt">lt规则</a>
验证是否小于某个数
~~~
lt:10
~~~

### <a id="lte">lte规则</a>
验证是否小于等于某个数
~~~
lte:10
~~~

#### <a id="date">date规则</a>
验证是否是日期格式 Y-m-d H:i:s类型
~~~
date
~~~