package lvalidator

import (
	"errors"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type validApi struct {
}

// 格式固定
// 规则调用时将驼峰转成_调用，比如RequiredIf变为required_if
// 传值 data 需要验证的数据
// 传值 ruleKey 验证的字段
// 传值 ruleValue 需要验证的规则(非必传)
// 返值 error 错误信息

// 验证必填
func (v validApi) Required(data interface{}, ruleKey string) error {
	rs := errors.New(strings.ReplaceAll(Lang.Required, "{ruleKey}", ruleKey))
	if data == nil {
		return rs
	}
	dataValue := reflect.ValueOf(data)
	if dataValue.Kind() == reflect.String && data.(string) == "" {
		return rs
	} else if dataValue.Kind() == reflect.Float64 && data.(float64) == 0 {
		return rs
	} else if dataValue.Kind() == reflect.Bool && data.(bool) == false {
		return rs
	} else if dataValue.Kind() == reflect.Slice && len(data.([]interface{})) == 0 {
		return rs
	}
	return nil
}

// 验证字符串
func (v validApi) String(data interface{}, ruleKey string) error {
	rs := errors.New(strings.ReplaceAll(Lang.String, "{ruleKey}", ruleKey))
	if data == nil {
		return rs
	}
	dataValue := reflect.ValueOf(data)
	if dataValue.Kind() == reflect.String {
		return nil
	}
	return rs
}

// 验证数字
func (v validApi) Number(data interface{}, ruleKey string) error {
	rs := errors.New(strings.ReplaceAll(Lang.Number, "{ruleKey}", ruleKey))
	if data == nil {
		return rs
	}
	dataValue := reflect.ValueOf(data)
	if dataValue.Kind() == reflect.Float64 {
		return nil
	} else if dataValue.Kind() == reflect.String {
		_, err := strconv.ParseFloat(data.(string), 64)
		if err != nil {
			return rs
		}
	} else {
		return rs
	}
	return nil
}

// 验证整数
func (v validApi) Integer(data interface{}, ruleKey string) error {
	rs := errors.New(strings.ReplaceAll(Lang.Integer, "{ruleKey}", ruleKey))
	if data == nil {
		return rs
	}
	dataValue := reflect.ValueOf(data)
	dataString := ""
	if dataValue.Kind() == reflect.Float64 {
		dataString = strconv.FormatFloat(data.(float64), 'f', -1, 64)
	} else if dataValue.Kind() == reflect.String {
		dataString = data.(string)
	} else {
		return rs
	}
	reg := regexp.MustCompile(`^\d*$`)
	if !reg.MatchString(dataString) {
		return rs
	}
	return nil
}

// 验证大于
func (v validApi) Gt(data interface{}, ruleKey string, ruleValue string) error {
	ruleValueFloat64, _ := strconv.ParseFloat(ruleValue, 64)
	info := strings.ReplaceAll(Lang.Gt, "{ruleKey}", ruleKey)
	info = strings.ReplaceAll(info, "{ruleValue}", ruleValue)
	rs := errors.New(info)
	if data == nil {
		return rs
	}
	dataValue := reflect.ValueOf(data)
	if dataValue.Kind() == reflect.Float64 && data.(float64) <= ruleValueFloat64 {
		return rs
	} else if dataValue.Kind() == reflect.String {
		dataFloat64, err := strconv.ParseFloat(data.(string), 64)
		if err != nil {
			return rs
		}
		if dataFloat64 <= ruleValueFloat64 {
			return rs
		}
	} else {
		return rs
	}
	return nil
}

// 验证大于等于
func (v validApi) Gte(data interface{}, ruleKey string, ruleValue string) error {
	ruleValueFloat64, _ := strconv.ParseFloat(ruleValue, 64)
	info := strings.ReplaceAll(Lang.Gte, "{ruleKey}", ruleKey)
	info = strings.ReplaceAll(info, "{ruleValue}", ruleValue)
	rs := errors.New(info)
	if data == nil {
		return rs
	}
	dataValue := reflect.ValueOf(data)
	if dataValue.Kind() == reflect.Float64 && data.(float64) < ruleValueFloat64 {
		return rs
	} else if dataValue.Kind() == reflect.String {
		dataFloat64, err := strconv.ParseFloat(data.(string), 64)
		if err != nil {
			return rs
		}
		if dataFloat64 < ruleValueFloat64 {
			return rs
		}
	} else {
		return rs
	}
	return nil
}

// 验证小于
func (v validApi) Lt(data interface{}, ruleKey string, ruleValue string) error {
	ruleValueFloat64, _ := strconv.ParseFloat(ruleValue, 64)
	info := strings.ReplaceAll(Lang.Lt, "{ruleKey}", ruleKey)
	info = strings.ReplaceAll(info, "{ruleValue}", ruleValue)
	rs := errors.New(info)
	if data == nil {
		return rs
	}
	dataValue := reflect.ValueOf(data)
	if dataValue.Kind() == reflect.Float64 && data.(float64) >= ruleValueFloat64 {
		return rs
	} else if dataValue.Kind() == reflect.String {
		dataFloat64, err := strconv.ParseFloat(data.(string), 64)
		if err != nil {
			return rs
		}
		if dataFloat64 >= ruleValueFloat64 {
			return rs
		}
	} else {
		return rs
	}
	return nil
}

// 验证小于等于
func (v validApi) Lte(data interface{}, ruleKey string, ruleValue string) error {
	ruleValueFloat64, _ := strconv.ParseFloat(ruleValue, 64)
	info := strings.ReplaceAll(Lang.Lte, "{ruleKey}", ruleKey)
	info = strings.ReplaceAll(info, "{ruleValue}", ruleValue)
	rs := errors.New(info)
	if data == nil {
		return rs
	}
	dataValue := reflect.ValueOf(data)
	if dataValue.Kind() == reflect.Float64 && data.(float64) > ruleValueFloat64 {
		return rs
	} else if dataValue.Kind() == reflect.String {
		dataFloat64, err := strconv.ParseFloat(data.(string), 64)
		if err != nil {
			return rs
		}
		if dataFloat64 > ruleValueFloat64 {
			return rs
		}
	} else {
		return rs
	}
	return nil
}

// 验证日期
func (v validApi) Date(data interface{}, ruleKey string) error {
	info := strings.ReplaceAll(Lang.Date, "{ruleKey}", ruleKey)
	rs := errors.New(info)
	dataValue := reflect.ValueOf(data)
	if dataValue.Kind() != reflect.String {
		return rs
	}
	if len(data.(string)) > 19 {
		return rs
	}
	formatAtByte := []byte("0000-00-00 00:00:00")
	copy(formatAtByte, []byte(data.(string)))
	_, err := time.ParseInLocation("2006-01-02 15:04:05", string(formatAtByte), time.Local)
	if err != nil {
		return rs
	}
	return nil
}
