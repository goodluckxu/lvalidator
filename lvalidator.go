package lvalidator

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"go-lib/handle_interface"
	"net/http"
	"reflect"
	"strings"
)

type Valid struct {
	request     *http.Request
	validApi    validApi
	validApiMap interface{}
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
	newData := reflect.ValueOf(data).Elem().Interface()
	for _, rule := range ruleList {
		err := v.validRule(newData, rule)
		if err != nil {
			return err
		}
	}
	return nil
}

func (v *Valid) ValidXml(data interface{}) error {
	body := Func.readBody(v.request)
	if err := xml.Unmarshal(body, data); err != nil {
		return err
	}
	return nil
}

// 验证规则
func (v *Valid) validRule(data interface{}, rule map[string]interface{}) error {
	ruleKey := rule["key"].(string)
	notes := rule["notes"].(string)
	validData := handle_interface.GetInterface(data, ruleKey)
	ruleList := rule["list"].([]interface{})
	for _, val := range ruleList {
		vType := reflect.TypeOf(val).Kind()
		if vType == reflect.String {
			vList := strings.Split(val.(string), ":")
			vVal := ""
			if len(vList) > 1 {
				vVal = vList[1]
			}
			if notes != "" {
				ruleKey = notes
			}
			if err := v.validStringRule(validData, ruleKey, vList[0], vVal); err != nil {
				return err
			}
		} else if vType == reflect.Func {
			if err := val.(func(value interface{}) error)(validData); err != nil {
				return err
			}
		} else {
			return errors.New("不存在的规则：键为" + ruleKey)
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
		if fn.Type().NumIn() == 3 {
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
	return nil
}

// 获取所有的验证方法
func (v *Valid) getValidApiFunc(rule string) reflect.Value {
	if v.validApiMap == nil {
		validApiValue := reflect.ValueOf(validApi{})
		numMethod := validApiValue.NumMethod()
		apiMap := map[string]reflect.Value{}
		for i := 0; i < numMethod; i++ {
			methodName := validApiValue.Type().Method(i).Name
			apiMap[Func.humpToUnderline(methodName)] = validApiValue.Method(i)
		}
		v.validApiMap = apiMap
	}
	return v.validApiMap.(map[string]reflect.Value)[rule]
}
