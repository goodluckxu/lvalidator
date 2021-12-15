package lvalidator

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
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
func (v validApi) Required(data interface{}, ruleKey string, ruleValue string) error {
	return Func.ValidData(data, ruleKey, func(validData interface{}, validNotes, validRule string) error {
		info := strings.ReplaceAll(Lang.Required, "{ruleKey}", validNotes)
		rs := errors.New(info)
		if validData == nil {
			return rs
		}
		switch validData.(type) {
		case string:
			if validData.(string) == "" && ruleValue != "string" {
				return rs
			}
		case float64:
			if validData.(float64) == 0 && ruleValue != "number" {
				return rs
			}
		case bool:
			if validData.(bool) == false && ruleValue != "bool" {
				return rs
			}
		case []interface{}:
			if len(validData.([]interface{})) == 0 && ruleValue != "array" {
				return rs
			}
		case map[string]interface{}:
			if len(validData.(map[string]interface{})) == 0 && ruleValue != "map" {
				return rs
			}
		}
		return nil
	})
}

// 验证为空不验证
func (v validApi) Nullable(data interface{}, ruleKey string, ruleValue string) error {
	return Func.ValidData(data, ruleKey, func(validData interface{}, validNotes, validRule string) error {
		isNull := false
		if validData == nil {
			isNull = true
		} else {
			switch validData.(type) {
			case string:
				if validData.(string) == "" {
					isNull = true
				}
			case float64:
				if validData.(float64) == 0 {
					isNull = true
				}
			case bool:
				if validData.(bool) == false {
					isNull = true
				}
			case []interface{}:
				if len(validData.([]interface{})) == 0 {
					isNull = true
				}
			case map[string]interface{}:
				if len(validData.(map[string]interface{})) == 0 {
					isNull = true
				}
			}
		}
		if isNull {
			return nil
		}
		return errors.New("")
	})
}

// 验证长度相等
func (v validApi) Len(data interface{}, ruleKey string, ruleValue string) error {
	return Func.ValidData(data, ruleKey, func(validData interface{}, validNotes, validRule string) error {
		ruleValueFloat64, _ := strconv.ParseFloat(ruleValue, 64)
		info := strings.ReplaceAll(Lang.Len, "{ruleKey}", validNotes)
		info = strings.ReplaceAll(info, "{ruleValue}", ruleValue)
		rs := errors.New(info)
		if validData == nil {
			return nil
		}
		switch validData.(type) {
		case string:
			if utf8.RuneCountInString(validData.(string)) == int(ruleValueFloat64) {
				return nil
			}
		case []interface{}:
			if len(validData.([]interface{})) == int(ruleValueFloat64) {
				return nil
			}
		}
		return rs
	})
}

// 验证最小长度
func (v validApi) Min(data interface{}, ruleKey string, ruleValue string) error {
	return Func.ValidData(data, ruleKey, func(validData interface{}, validNotes, validRule string) error {
		ruleValueFloat64, _ := strconv.ParseFloat(ruleValue, 64)
		info := strings.ReplaceAll(Lang.Min, "{ruleKey}", validNotes)
		info = strings.ReplaceAll(info, "{ruleValue}", ruleValue)
		rs := errors.New(info)
		if validData == nil {
			return nil
		}
		switch validData.(type) {
		case string:
			if utf8.RuneCountInString(validData.(string)) >= int(ruleValueFloat64) {
				return nil
			}
		case []interface{}:
			if len(validData.([]interface{})) >= int(ruleValueFloat64) {
				return nil
			}
		}
		return rs
	})
}

// 验证最大长度
func (v validApi) Max(data interface{}, ruleKey string, ruleValue string) error {
	return Func.ValidData(data, ruleKey, func(validData interface{}, validNotes, validRule string) error {
		ruleValueFloat64, _ := strconv.ParseFloat(ruleValue, 64)
		info := strings.ReplaceAll(Lang.Max, "{ruleKey}", validNotes)
		info = strings.ReplaceAll(info, "{ruleValue}", ruleValue)
		rs := errors.New(info)
		if validData == nil {
			return nil
		}
		switch validData.(type) {
		case string:
			if utf8.RuneCountInString(validData.(string)) <= int(ruleValueFloat64) {
				return nil
			}
		case []interface{}:
			if len(validData.([]interface{})) <= int(ruleValueFloat64) {
				return nil
			}
		}
		return rs
	})
}

// 验证日期
func (v validApi) Date(data interface{}, ruleKey string) error {
	return Func.ValidData(data, ruleKey, func(validData interface{}, validNotes, validRule string) error {
		info := strings.ReplaceAll(Lang.Date, "{ruleKey}", validNotes)
		rs := errors.New(info)
		if validData == nil {
			return nil
		}
		validDataString, bl := validData.(string)
		if !bl {
			return rs
		}
		if len(validDataString) > 19 {
			return rs
		}
		if _, err := Func.TimeParse(validDataString); err != nil {
			return rs
		}
		return nil
	})
}

// 验证日期格式化
func (v validApi) DateFormat(data interface{}, ruleKey string, ruleValue string) error {
	return Func.ValidData(data, ruleKey, func(validData interface{}, validNotes, validRule string) error {
		info := strings.ReplaceAll(Lang.DateFormat, "{ruleKey}", validNotes)
		info = strings.ReplaceAll(info, "{ruleValue}", ruleValue)
		rs := errors.New(info)
		if validData == nil {
			return nil
		}
		validDataString, bl := validData.(string)
		if !bl {
			return rs
		}
		if ruleValue == "" {
			ruleValue = "Y-m-d H:i:s"
		}
		if err := Func.ValidDate(validDataString, ruleValue); err != nil {
			return rs
		}
		return nil
	})
}

// 验证邮箱
func (v validApi) Email(data interface{}, ruleKey string) error {
	return Func.ValidData(data, ruleKey, func(validData interface{}, validNotes, validRule string) error {
		info := strings.ReplaceAll(Lang.Email, "{ruleKey}", validNotes)
		rs := errors.New(info)
		if validData == nil {
			return nil
		}
		validDataString, bl := validData.(string)
		if !bl {
			return rs
		}
		reg := regexp.MustCompile(`^[A-Za-z0-9\\u4e00-\\u9fa5]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$`)
		if reg.MatchString(validDataString) {
			return nil
		}
		return rs
	})
}

// 验证手机
func (v validApi) Phone(data interface{}, ruleKey string) error {
	return Func.ValidData(data, ruleKey, func(validData interface{}, validNotes, validRule string) error {
		info := strings.ReplaceAll(Lang.Phone, "{ruleKey}", validNotes)
		rs := errors.New(info)
		if validData == nil {
			return nil
		}
		validDataString, bl := validData.(string)
		if !bl {
			return rs
		}
		reg := regexp.MustCompile(`^1[0-9]{10}$`)
		if reg.MatchString(validDataString) {
			return nil
		}
		return rs
	})
}

// 验证是否在数组里面
func (v validApi) In(data interface{}, ruleKey string, ruleValue string) error {
	return Func.ValidData(data, ruleKey, func(validData interface{}, validNotes, validRule string) error {
		info := strings.ReplaceAll(Lang.In, "{ruleKey}", validNotes)
		info = strings.ReplaceAll(info, "{ruleValue}", ruleValue)
		rs := errors.New(info)
		if validData == nil {
			return nil
		}
		list := strings.Split(ruleValue, ",")
		if bl, _ := Func.InArray(Func.formatNumber(validData), list); !bl {
			return rs
		}
		return nil
	})
}

// 验证数组内的值唯一
func (v validApi) Unique(data interface{}, ruleKey string) error {
	starIndex := -1
	ruleKeyList := strings.Split(ruleKey, ".")
	for key, rule := range ruleKeyList {
		if rule == "*" {
			starIndex = key
		}
	}
	if starIndex+1 == len(ruleKeyList) {
		starIndex--
	}
	if starIndex == -1 {
		starIndex++
	}
	var compareRule string
	var compareData interface{}
	return Func.ValidData(data, ruleKey, func(validData interface{}, validNotes, validRule string) error {
		info := strings.ReplaceAll(Lang.Unique, "{ruleKey}", validNotes)
		rs := errors.New(info)
		validRuleList := strings.Split(validRule, ".")
		beforeRule := strings.Join(validRuleList[0:starIndex], ".")
		if compareRule == "" {
			compareRule = beforeRule
		} else {
			if !Func.IsEqualData(compareRule, beforeRule) {
				compareRule = beforeRule
				compareData = nil
			}
		}
		if compareData == nil {
			compareData = validData
		} else {
			if Func.IsEqualData(compareData, validData) {
				return rs
			}
		}
		return nil
	})
}

// 正则表达式
func (v validApi) Regexp(data interface{}, ruleKey string, ruleValue string) error {
	return Func.ValidData(data, ruleKey, func(validData interface{}, validNotes, validRule string) error {
		info := strings.ReplaceAll(Lang.Regexp, "{ruleKey}", validNotes)
		rs := errors.New(info)
		if validData == nil {
			return nil
		}
		validDataString, bl := validData.(string)
		if !bl {
			return rs
		}
		reg := regexp.MustCompile(ruleValue)
		if reg.MatchString(validDataString) {
			return nil
		}
		return rs
	})
}
