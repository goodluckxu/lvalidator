package lvalidator

import (
	"errors"
	"github.com/goodluckxu/go-lib/handle_interface"
	"reflect"
	"regexp"
	"strconv"
	"strings"
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
	return Func.ValidData(data, ruleKey, func(validData interface{}, validNotes, validRule string) error {
		info := strings.ReplaceAll(Lang.Required, "{ruleKey}", validNotes)
		rs := errors.New(info)
		if validData == nil {
			return rs
		}
		switch validData.(type) {
		case string:
			if validData.(string) == "" {
				return rs
			}
		case float64:
			if validData.(float64) == 0 {
				return rs
			}
		case bool:
			if validData.(bool) == false {
				return rs
			}
		case []interface{}:
			if len(validData.([]interface{})) == 0 {
				return rs
			}
		}
		return nil
	})
}

// 验证数组
func (v validApi) Array(data interface{}, ruleKey string) error {
	return Func.ValidData(data, ruleKey, func(validData interface{}, validNotes, validRule string) error {
		info := strings.ReplaceAll(Lang.Array, "{ruleKey}", validNotes)
		rs := errors.New(info)
		if validData == nil {
			return nil
		}
		if _, bl := validData.([]interface{}); bl {
			return nil
		}
		return rs
	})
}

// 验证对象
func (v validApi) Map(data interface{}, ruleKey string) error {
	return Func.ValidData(data, ruleKey, func(validData interface{}, validNotes, validRule string) error {
		info := strings.ReplaceAll(Lang.Map, "{ruleKey}", validNotes)
		rs := errors.New(info)
		if validData == nil {
			return nil
		}
		if _, bl := validData.(map[string]interface{}); bl {
			return nil
		}
		return rs
	})
}

// 验证字符串
func (v validApi) String(data interface{}, ruleKey string) error {
	return Func.ValidData(data, ruleKey, func(validData interface{}, validNotes, validRule string) error {
		info := strings.ReplaceAll(Lang.String, "{ruleKey}", validNotes)
		rs := errors.New(info)
		if validData == nil {
			return nil
		}
		if _, bl := validData.(string); bl {
			return nil
		}
		return rs
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
			if len(validData.(string)) == int(ruleValueFloat64) {
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
			if len(validData.(string)) >= int(ruleValueFloat64) {
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
			if len(validData.(string)) <= int(ruleValueFloat64) {
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

// 验证数字
func (v validApi) Number(data interface{}, ruleKey string) error {
	return Func.ValidData(data, ruleKey, func(validData interface{}, validNotes, validRule string) error {
		info := strings.ReplaceAll(Lang.Number, "{ruleKey}", validNotes)
		rs := errors.New(info)
		if validData == nil {
			return nil
		}
		switch validData.(type) {
		case float64:
			return nil
		case string:
			if number, err := strconv.ParseFloat(validData.(string), 64); err == nil {
				newData := handle_interface.UpdateInterface(reflect.ValueOf(data).Elem().Interface(), []handle_interface.Rule{
					{FindField: validRule, UpdateValue: number},
				})
				*data.(*interface{}) = newData
				return nil
			}
		}
		return rs
	})
}

// 验证整数
func (v validApi) Integer(data interface{}, ruleKey string) error {
	return Func.ValidData(data, ruleKey, func(validData interface{}, validNotes, validRule string) error {
		info := strings.ReplaceAll(Lang.Integer, "{ruleKey}", validNotes)
		rs := errors.New(info)
		if validData == nil {
			return nil
		}
		switch validData.(type) {
		case float64:
			validDataInt := int(validData.(float64))
			if validData.(float64) == float64(validDataInt) {
				newData := handle_interface.UpdateInterface(reflect.ValueOf(data).Elem().Interface(), []handle_interface.Rule{
					{FindField: validRule, UpdateValue: validDataInt},
				})
				*data.(*interface{}) = newData
				return nil
			}
		case string:
			dataInt, err := strconv.ParseInt(validData.(string), 10, 64)
			if err == nil {
				newData := handle_interface.UpdateInterface(reflect.ValueOf(data).Elem().Interface(), []handle_interface.Rule{
					{FindField: validRule, UpdateValue: int(dataInt)},
				})
				*data.(*interface{}) = newData
				return nil
			}
		}
		return rs
	})
}

// 验证大于
func (v validApi) Gt(data interface{}, ruleKey string, ruleValue string) error {
	return Func.ValidData(data, ruleKey, func(validData interface{}, validNotes, validRule string) error {
		ruleValueFloat64, _ := strconv.ParseFloat(ruleValue, 64)
		info := strings.ReplaceAll(Lang.Gt, "{ruleKey}", validNotes)
		info = strings.ReplaceAll(info, "{ruleValue}", ruleValue)
		rs := errors.New(info)
		if validData == nil {
			return nil
		}
		switch validData.(type) {
		case float64:
			if validData.(float64) > ruleValueFloat64 {
				return nil
			}
		case string:
			dataFloat64, err := strconv.ParseFloat(validData.(string), 64)
			if err != nil {
				return rs
			}
			if dataFloat64 > ruleValueFloat64 {
				return nil
			}
		}
		return rs
	})
}

// 验证大于等于
func (v validApi) Gte(data interface{}, ruleKey string, ruleValue string) error {
	return Func.ValidData(data, ruleKey, func(validData interface{}, validNotes, validRule string) error {
		ruleValueFloat64, _ := strconv.ParseFloat(ruleValue, 64)
		info := strings.ReplaceAll(Lang.Gte, "{ruleKey}", validNotes)
		info = strings.ReplaceAll(info, "{ruleValue}", ruleValue)
		rs := errors.New(info)
		if validData == nil {
			return nil
		}
		switch validData.(type) {
		case float64:
			if validData.(float64) >= ruleValueFloat64 {
				return nil
			}
		case string:
			dataFloat64, err := strconv.ParseFloat(validData.(string), 64)
			if err != nil {
				return rs
			}
			if dataFloat64 >= ruleValueFloat64 {
				return nil
			}
		}
		return rs
	})
}

// 验证小于
func (v validApi) Lt(data interface{}, ruleKey string, ruleValue string) error {
	return Func.ValidData(data, ruleKey, func(validData interface{}, validNotes, validRule string) error {
		ruleValueFloat64, _ := strconv.ParseFloat(ruleValue, 64)
		info := strings.ReplaceAll(Lang.Lt, "{ruleKey}", validNotes)
		info = strings.ReplaceAll(info, "{ruleValue}", ruleValue)
		rs := errors.New(info)
		if validData == nil {
			return nil
		}
		switch validData.(type) {
		case float64:
			if validData.(float64) < ruleValueFloat64 {
				return nil
			}
		case string:
			dataFloat64, err := strconv.ParseFloat(validData.(string), 64)
			if err != nil {
				return rs
			}
			if dataFloat64 < ruleValueFloat64 {
				return nil
			}
		}
		return rs
	})
}

// 验证小于等于
func (v validApi) Lte(data interface{}, ruleKey string, ruleValue string) error {
	return Func.ValidData(data, ruleKey, func(validData interface{}, validNotes, validRule string) error {
		ruleValueFloat64, _ := strconv.ParseFloat(ruleValue, 64)
		info := strings.ReplaceAll(Lang.Lte, "{ruleKey}", validNotes)
		info = strings.ReplaceAll(info, "{ruleValue}", ruleValue)
		rs := errors.New(info)
		if validData == nil {
			return nil
		}
		switch validData.(type) {
		case float64:
			if validData.(float64) <= ruleValueFloat64 {
				return nil
			}
		case string:
			dataFloat64, err := strconv.ParseFloat(validData.(string), 64)
			if err != nil {
				return rs
			}
			if dataFloat64 <= ruleValueFloat64 {
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

// 日期大于
func (v validApi) DateGt(data interface{}, ruleKey string, ruleValue string) error {
	return Func.ValidData(data, ruleKey, func(validData interface{}, validNotes, validRule string) error {
		ruleValueTime, err := Func.TimeParse(ruleValue)
		if err != nil {
			info := strings.ReplaceAll(Lang.Error, "{rule}", "date_gt")
			info = strings.ReplaceAll(info, "{ruleValue}", ruleValue)
			info = strings.ReplaceAll(info, "{error}", err.Error())
			return errors.New(info)
		}
		info := strings.ReplaceAll(Lang.DateGt, "{ruleKey}", validNotes)
		info = strings.ReplaceAll(info, "{ruleValue}", ruleValue)
		rs := errors.New(info)
		if validData == nil {
			return nil
		}
		validDataString, bl := validData.(string)
		if !bl {
			return rs
		}
		validDataTime, err := Func.TimeParse(validDataString)
		if err != nil {
			return rs
		}
		if validDataTime.Unix() > ruleValueTime.Unix() {
			return nil
		}
		return rs
	})
}

// 日期大于等于
func (v validApi) DateGte(data interface{}, ruleKey string, ruleValue string) error {
	return Func.ValidData(data, ruleKey, func(validData interface{}, validNotes, validRule string) error {
		ruleValueTime, err := Func.TimeParse(ruleValue)
		if err != nil {
			info := strings.ReplaceAll(Lang.Error, "{rule}", "date_gt")
			info = strings.ReplaceAll(info, "{ruleValue}", ruleValue)
			info = strings.ReplaceAll(info, "{error}", err.Error())
			return errors.New(info)
		}
		info := strings.ReplaceAll(Lang.DateGte, "{ruleKey}", validNotes)
		info = strings.ReplaceAll(info, "{ruleValue}", ruleValue)
		rs := errors.New(info)
		if validData == nil {
			return nil
		}
		validDataString, bl := validData.(string)
		if !bl {
			return rs
		}
		validDataTime, err := Func.TimeParse(validDataString)
		if err != nil {
			return rs
		}
		if validDataTime.Unix() >= ruleValueTime.Unix() {
			return nil
		}
		return rs
	})
}

// 日期小于
func (v validApi) DateLt(data interface{}, ruleKey string, ruleValue string) error {
	return Func.ValidData(data, ruleKey, func(validData interface{}, validNotes, validRule string) error {
		ruleValueTime, err := Func.TimeParse(ruleValue)
		if err != nil {
			info := strings.ReplaceAll(Lang.Error, "{rule}", "date_gt")
			info = strings.ReplaceAll(info, "{ruleValue}", ruleValue)
			info = strings.ReplaceAll(info, "{error}", err.Error())
			return errors.New(info)
		}
		info := strings.ReplaceAll(Lang.DateLt, "{ruleKey}", validNotes)
		info = strings.ReplaceAll(info, "{ruleValue}", ruleValue)
		rs := errors.New(info)
		if validData == nil {
			return nil
		}
		validDataString, bl := validData.(string)
		if !bl {
			return rs
		}
		validDataTime, err := Func.TimeParse(validDataString)
		if err != nil {
			return rs
		}
		if validDataTime.Unix() < ruleValueTime.Unix() {
			return nil
		}
		return rs
	})
}

// 日期小于等于
func (v validApi) DateLte(data interface{}, ruleKey string, ruleValue string) error {
	return Func.ValidData(data, ruleKey, func(validData interface{}, validNotes, validRule string) error {
		ruleValueTime, err := Func.TimeParse(ruleValue)
		if err != nil {
			info := strings.ReplaceAll(Lang.Error, "{rule}", "date_gt")
			info = strings.ReplaceAll(info, "{ruleValue}", ruleValue)
			info = strings.ReplaceAll(info, "{error}", err.Error())
			return errors.New(info)
		}
		info := strings.ReplaceAll(Lang.DateLte, "{ruleKey}", validNotes)
		info = strings.ReplaceAll(info, "{ruleValue}", ruleValue)
		rs := errors.New(info)
		if validData == nil {
			return nil
		}
		validDataString, bl := validData.(string)
		if !bl {
			return rs
		}
		validDataTime, err := Func.TimeParse(validDataString)
		if err != nil {
			return rs
		}
		if validDataTime.Unix() <= ruleValueTime.Unix() {
			return nil
		}
		return rs
	})
}

// 等于字段
func (v validApi) EqField(data interface{}, ruleKey string, ruleValue string) error {
	return Func.ValidDoubleData(data, ruleKey, ruleValue, func(validData, compareData interface{}, validNotes, compareNotes string) error {
		info := strings.ReplaceAll(Lang.EqField, "{ruleKey}", validNotes)
		info = strings.ReplaceAll(info, "{ruleValue}", compareNotes)
		rs := errors.New(info)
		if validData == nil || compareData == nil {
			if validData == compareData {
				return nil
			}
			return rs
		}
		switch validData.(type) {
		case string:
			if compareDataString, bl := compareData.(string); bl {
				if validData.(string) == compareDataString {
					return nil
				}
			}
		case float64:
			if compareDataFloat64, bl := compareData.(float64); bl {
				if validData.(float64) == compareDataFloat64 {
					return nil
				}
			}
		case bool:
			if compareDataBool, bl := compareData.(bool); bl {
				if validData.(bool) == compareDataBool {
					return nil
				}
			}
		default:
			if reflect.DeepEqual(validData, compareData) {
				return nil
			}
		}
		return rs
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
