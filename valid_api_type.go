package lvalidator

import (
	"errors"
	"github.com/goodluckxu/go-lib/handle_interface"
	"reflect"
	"strconv"
	"strings"
)

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
			dataValue := reflect.ValueOf(data)
			if dataValue.Kind() == reflect.Ptr {
				dataValue = dataValue.Elem()
			}
			if dataValue.Kind() == reflect.Struct {
				return rs
			}
			if number, err := strconv.ParseFloat(validData.(string), 64); err == nil {
				newData := handle_interface.UpdateInterface(reflect.ValueOf(data).Elem().Interface(), []handle_interface.Rule{
					{FindField: validRule, UpdateValue: number},
				})
				dataValue.Set(reflect.ValueOf(newData))
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
				return nil
			}
		case string:
			dataValue := reflect.ValueOf(data)
			if dataValue.Kind() == reflect.Ptr {
				dataValue = dataValue.Elem()
			}
			if dataValue.Kind() == reflect.Struct {
				return rs
			}
			if dataInt, err := strconv.ParseInt(validData.(string), 10, 64); err == nil {
				newData := handle_interface.UpdateInterface(reflect.ValueOf(data).Elem().Interface(), []handle_interface.Rule{
					{FindField: validRule, UpdateValue: float64(dataInt)},
				})
				dataValue.Set(reflect.ValueOf(newData))
				return nil
			}
		}
		return rs
	})
}

// 验证布尔
func (v validApi) Bool(data interface{}, ruleKey string, ruleValue string) error {
	return Func.ValidData(data, ruleKey, func(validData interface{}, validNotes, validRule string) error {
		info := strings.ReplaceAll(Lang.Bool, "{ruleKey}", validNotes)
		rs := errors.New(info)
		if validData == nil {
			return nil
		}
		switch validData.(type) {
		case bool:
			return nil
		case float64:
			validDataFloat := validData.(float64)
			if validDataFloat == 0 || validDataFloat == 1 {
				return nil
			}
		}
		return rs
	})
}