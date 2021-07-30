package lvalidator

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

type Valid struct {
	request  *http.Request
	validApi validApi
}

func New(r *http.Request) *Valid {
	valid := new(Valid)
	valid.request = r
	return valid
}

func (v *Valid) ValidJson(rules map[string]interface{}, data interface{}) error {
	ruleList, err := Func.parseRules(rules)
	if err != nil {
		return err
	}
	body := Func.readBody(v.request)
	if err := json.Unmarshal(body, data); err != nil {
		return err
	}
	for _, rule := range ruleList {
		err := v.validRule(data, rule)
		if err != nil {
			return err
		}
	}
	return nil
}

func (v *Valid) ValidXml(rules map[string]interface{}, data interface{}) error {
	ruleList, err := Func.parseRules(rules)
	if err != nil {
		return err
	}
	body := Func.readBody(v.request)
	if err := xml.Unmarshal(body, data); err != nil {
		return err
	}
	newData := reflect.ValueOf(data).Elem().Interface()
	for _, rule := range ruleList {
		err := v.validRule(newData, rule)
		if err != nil {
			return err
		}
	}
	return nil
}

// 验证规则
func (v *Valid) validRule(data interface{}, rule map[string]interface{}) error {
	ruleKey := rule["key"].(string)
	ruleList := rule["list"].([]interface{})
	for _, val := range ruleList {
		switch val.(type) {
		case string:
			vList := strings.Split(val.(string), ":")
			vVal := strings.Join(vList[1:], ":")
			if err := v.validStringRule(data, ruleKey, vList[0], vVal); err != nil {
				if vList[0] == "through_condition_field" && err.Error() == "" {
					continue
				}
				return err
			}
			if vList[0] == "through_condition_field" {
				return nil
			}
		case func(interface{}) error:
			if err := Func.ValidData(data, ruleKey, func(validData interface{}, validNotes, validRule string) error {
				if err := val.(func(interface{}) error)(validData); err != nil {
					return err
				}
				return nil
			}); err != nil {
				return err
			}
		case func(interface{}, string) error:
			if err := Func.ValidData(data, ruleKey, func(validData interface{}, validNotes, validRule string) error {
				if err := val.(func(interface{}, string) error)(validData, validNotes); err != nil {
					return err
				}
				return nil
			}); err != nil {
				return err
			}
		}
	}
	return nil
}

// 验证字符串
func (v *Valid) validStringRule(data interface{}, ruleKey string, rule string, ruleValue string) error {
	fn := v.getValidApiFunc(rule)
	if fn.IsValid() {
		rValues := []reflect.Value{}
		if data == nil {
			nilArg := reflect.Zero(reflect.TypeOf((*error)(nil)).Elem())
			rValues = append(rValues, nilArg)
		} else {
			rValues = append(rValues, reflect.ValueOf(data))
		}
		rValues = append(rValues, reflect.ValueOf(ruleKey))
		if fn.Type().NumIn() >= 3 {
			rValues = append(rValues, reflect.ValueOf(ruleValue))
		}
		rs := fn.Call(rValues)
		errInterface := rs[0].Interface()
		if errInterface == nil {
			return nil
		}
		if err := errInterface.(error); err != nil {
			return err
		}
	}
	return errors.New(fmt.Sprintf("不存在的规则: %s", rule))
}

// 获取所有的验证方法
func (v *Valid) getValidApiFunc(rule string) reflect.Value {
	if !ValidApiMap[rule].IsValid() {
		validApiValue := reflect.ValueOf(validApi{})
		numMethod := validApiValue.NumMethod()
		apiMap := map[string]reflect.Value{}
		for i := 0; i < numMethod; i++ {
			methodName := validApiValue.Type().Method(i).Name
			apiMap[Func.humpToUnderline(methodName)] = validApiValue.Method(i)
		}
		ValidApiMap = apiMap
	}
	return ValidApiMap[rule]
}
