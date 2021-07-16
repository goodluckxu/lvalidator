package lvalidator

import (
	"errors"
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
// 传值 notes 注释(非必传)
// 返值 error 错误信息

// 验证必填
func (v validApi) Required(data interface{}, ruleKey string) error {
	return Func.ValidData(data, ruleKey, func(validData interface{}, rule string) error {
		notes := Func.GetNotes(ruleKey, rule)
		info := strings.ReplaceAll(Lang.Required, "{ruleKey}", notes)
		rs := errors.New(info)
		if validData == nil {
			return rs
		}
		dataValue := reflect.ValueOf(validData)
		if dataValue.Kind() == reflect.String && validData.(string) == "" {
			return rs
		} else if dataValue.Kind() == reflect.Float64 && validData.(float64) == 0 {
			return rs
		} else if dataValue.Kind() == reflect.Bool && validData.(bool) == false {
			return rs
		} else if dataValue.Kind() == reflect.Slice && len(validData.([]interface{})) == 0 {
			return rs
		}
		return nil
	})
}

// 验证数组
func (v validApi) Array(data interface{}, ruleKey string) error {
	return Func.ValidData(data, ruleKey, func(validData interface{}, rule string) error {
		notes := Func.GetNotes(ruleKey, rule)
		info := strings.ReplaceAll(Lang.Array, "{ruleKey}", notes)
		rs := errors.New(info)
		if validData == nil {
			return nil
		}
		dataValue := reflect.ValueOf(validData)
		if dataValue.Kind() == reflect.Slice {
			return nil
		}
		return rs
	})
}

// 验证对象
func (v validApi) Map(data interface{}, ruleKey string) error {
	return Func.ValidData(data, ruleKey, func(validData interface{}, rule string) error {
		notes := Func.GetNotes(ruleKey, rule)
		info := strings.ReplaceAll(Lang.Map, "{ruleKey}", notes)
		rs := errors.New(info)
		if validData == nil {
			return nil
		}
		dataValue := reflect.ValueOf(validData)
		if dataValue.Kind() == reflect.Map {
			return nil
		}
		return rs
	})
}

// 验证字符串
func (v validApi) String(data interface{}, ruleKey string) error {
	return Func.ValidData(data, ruleKey, func(validData interface{}, rule string) error {
		notes := Func.GetNotes(ruleKey, rule)
		info := strings.ReplaceAll(Lang.String, "{ruleKey}", notes)
		rs := errors.New(info)
		if validData == nil {
			return nil
		}
		dataValue := reflect.ValueOf(validData)
		if dataValue.Kind() == reflect.String {
			return nil
		}
		return rs
	})
}

// 验证长度相等
func (v validApi) Len(data interface{}, ruleKey string, ruleValue string) error {
	return Func.ValidData(data, ruleKey, func(validData interface{}, rule string) error {
		ruleValueFloat64, _ := strconv.ParseFloat(ruleValue, 64)
		notes := Func.GetNotes(ruleKey, rule)
		info := strings.ReplaceAll(Lang.Len, "{ruleKey}", notes)
		info = strings.ReplaceAll(info, "{ruleValue}", ruleValue)
		rs := errors.New(info)
		if validData == nil {
			return nil
		}
		dataValue := reflect.ValueOf(validData)
		if dataValue.Kind() == reflect.String && len(validData.(string)) == int(ruleValueFloat64) {
			return nil
		} else if dataValue.Kind() == reflect.Slice && len(validData.([]interface{})) == int(ruleValueFloat64) {
			return nil
		}
		return rs
	})
}

// 验证长度相等
func (v validApi) Min(data interface{}, ruleKey string, ruleValue string) error {
	return Func.ValidData(data, ruleKey, func(validData interface{}, rule string) error {
		ruleValueFloat64, _ := strconv.ParseFloat(ruleValue, 64)
		notes := Func.GetNotes(ruleKey, rule)
		info := strings.ReplaceAll(Lang.Min, "{ruleKey}", notes)
		info = strings.ReplaceAll(info, "{ruleValue}", ruleValue)
		rs := errors.New(info)
		if validData == nil {
			return nil
		}
		dataValue := reflect.ValueOf(validData)
		if dataValue.Kind() == reflect.String && len(validData.(string)) >= int(ruleValueFloat64) {
			return nil
		} else if dataValue.Kind() == reflect.Slice && len(validData.([]interface{})) >= int(ruleValueFloat64) {
			return nil
		}
		return rs
	})
}

// 验证长度相等
func (v validApi) Max(data interface{}, ruleKey string, ruleValue string) error {
	return Func.ValidData(data, ruleKey, func(validData interface{}, rule string) error {
		ruleValueFloat64, _ := strconv.ParseFloat(ruleValue, 64)
		notes := Func.GetNotes(ruleKey, rule)
		info := strings.ReplaceAll(Lang.Max, "{ruleKey}", notes)
		info = strings.ReplaceAll(info, "{ruleValue}", ruleValue)
		rs := errors.New(info)
		if validData == nil {
			return nil
		}
		dataValue := reflect.ValueOf(validData)
		if dataValue.Kind() == reflect.String && len(validData.(string)) <= int(ruleValueFloat64) {
			return nil
		} else if dataValue.Kind() == reflect.Slice && len(validData.([]interface{})) <= int(ruleValueFloat64) {
			return nil
		}
		return rs
	})
}

// 验证数字
func (v validApi) Number(data interface{}, ruleKey string) error {
	return Func.ValidData(data, ruleKey, func(validData interface{}, rule string) error {
		notes := Func.GetNotes(ruleKey, rule)
		info := strings.ReplaceAll(Lang.Number, "{ruleKey}", notes)
		rs := errors.New(info)
		if validData == nil {
			return nil
		}
		dataValue := reflect.ValueOf(validData)
		if dataValue.Kind() == reflect.Float64 {
			return nil
		} else if dataValue.Kind() == reflect.String {
			_, err := strconv.ParseFloat(validData.(string), 64)
			if err == nil {
				return nil
			}
		}
		return rs
	})
}

// 验证整数
func (v validApi) Integer(data interface{}, ruleKey string) error {
	return Func.ValidData(data, ruleKey, func(validData interface{}, rule string) error {
		notes := Func.GetNotes(ruleKey, rule)
		info := strings.ReplaceAll(Lang.Integer, "{ruleKey}", notes)
		rs := errors.New(info)
		if validData == nil {
			return nil
		}
		dataValue := reflect.ValueOf(validData)
		dataString := ""
		if dataValue.Kind() == reflect.Float64 {
			dataString = strconv.FormatFloat(validData.(float64), 'f', -1, 64)
		} else if dataValue.Kind() == reflect.String {
			dataString = validData.(string)
		} else {
			return rs
		}
		reg := regexp.MustCompile(`^\d*$`)
		if !reg.MatchString(dataString) {
			return rs
		}
		return nil
	})
}

// 验证大于
func (v validApi) Gt(data interface{}, ruleKey string, ruleValue string) error {
	return Func.ValidData(data, ruleKey, func(validData interface{}, rule string) error {
		ruleValueFloat64, _ := strconv.ParseFloat(ruleValue, 64)
		notes := Func.GetNotes(ruleKey, rule)
		info := strings.ReplaceAll(Lang.Gt, "{ruleKey}", notes)
		info = strings.ReplaceAll(info, "{ruleValue}", ruleValue)
		rs := errors.New(info)
		if validData == nil {
			return nil
		}
		dataValue := reflect.ValueOf(validData)
		if dataValue.Kind() == reflect.Float64 && validData.(float64) > ruleValueFloat64 {
			return nil
		} else if dataValue.Kind() == reflect.String {
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
	return Func.ValidData(data, ruleKey, func(validData interface{}, rule string) error {
		ruleValueFloat64, _ := strconv.ParseFloat(ruleValue, 64)
		notes := Func.GetNotes(ruleKey, rule)
		info := strings.ReplaceAll(Lang.Gte, "{ruleKey}", notes)
		info = strings.ReplaceAll(info, "{ruleValue}", ruleValue)
		rs := errors.New(info)
		if validData == nil {
			return nil
		}
		dataValue := reflect.ValueOf(validData)
		if dataValue.Kind() == reflect.Float64 && validData.(float64) >= ruleValueFloat64 {
			return nil
		} else if dataValue.Kind() == reflect.String {
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
	return Func.ValidData(data, ruleKey, func(validData interface{}, rule string) error {
		ruleValueFloat64, _ := strconv.ParseFloat(ruleValue, 64)
		notes := Func.GetNotes(ruleKey, rule)
		info := strings.ReplaceAll(Lang.Lt, "{ruleKey}", notes)
		info = strings.ReplaceAll(info, "{ruleValue}", ruleValue)
		rs := errors.New(info)
		if validData == nil {
			return nil
		}
		dataValue := reflect.ValueOf(validData)
		if dataValue.Kind() == reflect.Float64 && validData.(float64) < ruleValueFloat64 {
			return nil
		} else if dataValue.Kind() == reflect.String {
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
	return Func.ValidData(data, ruleKey, func(validData interface{}, rule string) error {
		ruleValueFloat64, _ := strconv.ParseFloat(ruleValue, 64)
		notes := Func.GetNotes(ruleKey, rule)
		info := strings.ReplaceAll(Lang.Lte, "{ruleKey}", notes)
		info = strings.ReplaceAll(info, "{ruleValue}", ruleValue)
		rs := errors.New(info)
		if validData == nil {
			return nil
		}
		dataValue := reflect.ValueOf(validData)
		if dataValue.Kind() == reflect.Float64 && validData.(float64) <= ruleValueFloat64 {
			return nil
		} else if dataValue.Kind() == reflect.String {
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
	return Func.ValidData(data, ruleKey, func(validData interface{}, rule string) error {
		notes := Func.GetNotes(ruleKey, rule)
		info := strings.ReplaceAll(Lang.Date, "{ruleKey}", notes)
		rs := errors.New(info)
		if validData == nil {
			return nil
		}
		dataValue := reflect.ValueOf(validData)
		if dataValue.Kind() != reflect.String {
			return rs
		}
		if len(validData.(string)) > 19 {
			return rs
		}
		_, err := Func.TimeParse(validData.(string))
		if err != nil {
			return rs
		}
		return nil
	})
}

// 日期大于
func (v validApi) DateGt(data interface{}, ruleKey string, ruleValue string) error {
	return Func.ValidData(data, ruleKey, func(validData interface{}, rule string) error {
		ruleValueTime, err := Func.TimeParse(ruleValue)
		if err != nil {
			info := strings.ReplaceAll(Lang.Error, "{rule}", "date_gt")
			info = strings.ReplaceAll(info, "{ruleValue}", ruleValue)
			info = strings.ReplaceAll(info, "{error}", err.Error())
			return errors.New(info)
		}
		notes := Func.GetNotes(ruleKey, rule)
		info := strings.ReplaceAll(Lang.DateGt, "{ruleKey}", notes)
		info = strings.ReplaceAll(info, "{ruleValue}", ruleValue)
		rs := errors.New(info)
		if validData == nil {
			return nil
		}
		dataValue := reflect.ValueOf(validData)
		if dataValue.Kind() != reflect.String {
			return rs
		}
		validDataTime, err := Func.TimeParse(validData.(string))
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
	return Func.ValidData(data, ruleKey, func(validData interface{}, rule string) error {
		ruleValueTime, err := Func.TimeParse(ruleValue)
		if err != nil {
			info := strings.ReplaceAll(Lang.Error, "{rule}", "date_gt")
			info = strings.ReplaceAll(info, "{ruleValue}", ruleValue)
			info = strings.ReplaceAll(info, "{error}", err.Error())
			return errors.New(info)
		}
		notes := Func.GetNotes(ruleKey, rule)
		info := strings.ReplaceAll(Lang.DateGte, "{ruleKey}", notes)
		info = strings.ReplaceAll(info, "{ruleValue}", ruleValue)
		rs := errors.New(info)
		if validData == nil {
			return nil
		}
		dataValue := reflect.ValueOf(validData)
		if dataValue.Kind() != reflect.String {
			return rs
		}
		validDataTime, err := Func.TimeParse(validData.(string))
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
	return Func.ValidData(data, ruleKey, func(validData interface{}, rule string) error {
		ruleValueTime, err := Func.TimeParse(ruleValue)
		if err != nil {
			info := strings.ReplaceAll(Lang.Error, "{rule}", "date_gt")
			info = strings.ReplaceAll(info, "{ruleValue}", ruleValue)
			info = strings.ReplaceAll(info, "{error}", err.Error())
			return errors.New(info)
		}
		notes := Func.GetNotes(ruleKey, rule)
		info := strings.ReplaceAll(Lang.DateLt, "{ruleKey}", notes)
		info = strings.ReplaceAll(info, "{ruleValue}", ruleValue)
		rs := errors.New(info)
		if validData == nil {
			return nil
		}
		dataValue := reflect.ValueOf(validData)
		if dataValue.Kind() != reflect.String {
			return rs
		}
		validDataTime, err := Func.TimeParse(validData.(string))
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
	return Func.ValidData(data, ruleKey, func(validData interface{}, rule string) error {
		ruleValueTime, err := Func.TimeParse(ruleValue)
		if err != nil {
			info := strings.ReplaceAll(Lang.Error, "{rule}", "date_gt")
			info = strings.ReplaceAll(info, "{ruleValue}", ruleValue)
			info = strings.ReplaceAll(info, "{error}", err.Error())
			return errors.New(info)
		}
		notes := Func.GetNotes(ruleKey, rule)
		info := strings.ReplaceAll(Lang.DateLte, "{ruleKey}", notes)
		info = strings.ReplaceAll(info, "{ruleValue}", ruleValue)
		rs := errors.New(info)
		if validData == nil {
			return nil
		}
		dataValue := reflect.ValueOf(validData)
		if dataValue.Kind() != reflect.String {
			return rs
		}
		validDataTime, err := Func.TimeParse(validData.(string))
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
	starInfo := Func.GetTwoFieldStarInfo(ruleKey, ruleValue)
	return Func.ValidData(data, ruleKey, func(validData interface{}, rule string) error {
		return Func.ValidData(data, ruleValue, func(ruleData interface{}, ruleField string) error {
			if !Func.IsTwoFieldCompare(rule, ruleField, starInfo) {
				return nil
			}
			notesKey := Func.GetNotes(ruleKey, rule)
			info := strings.ReplaceAll(Lang.EqField, "{ruleKey}", notesKey)
			notesValue := Func.GetNotes(ruleValue, ruleField)
			info = strings.ReplaceAll(info, "{ruleValue}", notesValue)
			rs := errors.New(info)
			if validData == nil || ruleData == nil {
				if validData == ruleData {
					return nil
				}
			}
			validValue := reflect.ValueOf(validData)
			ruleValue := reflect.ValueOf(ruleData)
			if validValue.Kind() != ruleValue.Kind() {
				return rs
			}
			if validValue.Kind() == reflect.String ||
				validValue.Kind() == reflect.Float64 ||
				validValue.Kind() == reflect.Bool {
				if validData == ruleData {
					return nil
				}
			} else {
				if reflect.DeepEqual(validData, ruleData) {
					return nil
				}
			}
			return rs
		})
	})
}

// 验证邮箱
func (v validApi) Email(data interface{}, ruleKey string) error {
	return Func.ValidData(data, ruleKey, func(validData interface{}, rule string) error {
		notes := Func.GetNotes(ruleKey, rule)
		info := strings.ReplaceAll(Lang.Email, "{ruleKey}", notes)
		rs := errors.New(info)
		if validData == nil {
			return nil
		}
		dataValue := reflect.ValueOf(validData)
		if dataValue.Kind() != reflect.String {
			return rs
		}
		reg := regexp.MustCompile(`^[A-Za-z0-9\\u4e00-\\u9fa5]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$`)
		if reg.MatchString(validData.(string)) {
			return nil
		}
		return rs
	})
}

// 验证手机
func (v validApi) Phone(data interface{}, ruleKey string) error {
	return Func.ValidData(data, ruleKey, func(validData interface{}, rule string) error {
		notes := Func.GetNotes(ruleKey, rule)
		info := strings.ReplaceAll(Lang.Phone, "{ruleKey}", notes)
		rs := errors.New(info)
		if validData == nil {
			return nil
		}
		dataValue := reflect.ValueOf(validData)
		if dataValue.Kind() != reflect.String {
			return rs
		}
		reg := regexp.MustCompile(`^1[0-9]{10}$`)
		if reg.MatchString(validData.(string)) {
			return nil
		}
		return rs
	})
}
