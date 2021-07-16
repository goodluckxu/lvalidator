package lvalidator

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var Func Function

type Function struct {
}

// 将传入的规格map解析
func (f Function) parseRules(rules map[string]interface{}) ([]map[string]interface{}, error) {
	ruleList := []map[string]interface{}{}
	for key, rule := range rules {
		ruleMap := map[string]interface{}{}
		list := []interface{}{}
		if rule == nil {
			return nil, errors.New("规格类型不正确")
		}
		ruleType := reflect.TypeOf(rule).Kind()
		if ruleType == reflect.String {
			if rule.(string) != "" {
				for _, v := range strings.Split(rule.(string), "|") {
					list = append(list, v)
				}
			}
		} else if ruleType == reflect.Slice {
			for _, v := range rule.([]interface{}) {
				list = append(list, v)
			}
		} else {
			return nil, errors.New("规格类型不正确")
		}
		sort := ""
		notes := ""
		newList := []interface{}{}
		for _, v := range list {
			if v == nil {
				return nil, errors.New("规格类型不正确")
			}
			vType := reflect.TypeOf(v).Kind()
			if vType == reflect.String {
				vList := strings.Split(v.(string), ":")
				if vList[0] != "sort" && vList[0] != "notes" {
					newList = append(newList, v)
					continue
				}
				if len(vList) < 2 {
					continue
				}
				switch vList[0] {
				case "sort":
					sort = vList[1]
				case "notes":
					notes = vList[1]
				}
			} else if vType == reflect.Func {
				newList = append(newList, v)
			} else {
				return nil, errors.New("规格类型不正确")
			}
		}
		if len(newList) == 0 {
			continue
		}
		RuleNotes[key] = notes
		if notes == "" {
			RuleNotes[key] = key
		}
		ruleMap["key"] = key
		ruleMap["sort"] = sort
		ruleMap["notes"] = notes
		ruleMap["list"] = newList
		ruleList = append(ruleList, ruleMap)
	}
	f.sortSliceMapStringInterface(ruleList, "sort")
	return ruleList, nil
}

// 读取body内容
func (f Function) readBody(r *http.Request) []byte {
	var bodyBytes []byte // 我们需要的body内容
	// 从原有Request.Body读取
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return bodyBytes
	}
	// 新建缓冲区并替换原有Request.body
	r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	return bodyBytes
}

// 排序[]map[string]interface{}
func (f Function) sortSliceMapStringInterface(data []map[string]interface{}, args ...string) {
	if len(args) == 0 {
		return
	}
	lenData := len(data)
	for i := 0; i < lenData-1; i++ {
		for j := 0; j < lenData-1-i; j++ {
			for _, arg := range args {
				argList := strings.Split(arg, " ")
				field := argList[0]
				sortType := "asc"
				if len(argList) > 1 {
					sortType = argList[1]
				}
				strJ := f.formatNumber(data[j][field])
				strJAdd := f.formatNumber(data[j+1][field])
				compareStr := strings.Compare(strJ, strJAdd)
				if compareStr == 0 {
					continue
				}
				if sortType == "asc" && compareStr == 1 {
					temp := data[j]
					data[j] = data[j+1]
					data[j+1] = temp
				} else if sortType == "desc" && compareStr == -1 {
					temp := data[j]
					data[j] = data[j+1]
					data[j+1] = temp
				}
				break
			}
		}
	}
	return
}

// 将数字转成字符串
func (f Function) formatNumber(i interface{}) string {
	if i == nil {
		return ""
	}
	var int64I int64
	switch reflect.TypeOf(i).Kind().String() {
	case "int":
		int64I = int64(i.(int))
	case "int8":
		int64I = int64(i.(int8))
	case "int16":
		int64I = int64(i.(int16))
	case "int32":
		int64I = int64(i.(int32))
	case "int64":
		int64I = i.(int64)
	case "uint":
		int64I = int64(i.(uint))
	case "uint8":
		int64I = int64(i.(uint8))
	case "uint16":
		int64I = int64(i.(uint16))
	case "uint32":
		int64I = int64(i.(uint32))
	case "uint64":
		int64I = int64(i.(uint64))
	case "float32":
		int64I = int64(i.(float32))
	case "float64":
		int64I = int64(i.(float64))
	case "string":
		return i.(string)
	default:
		return ""
	}
	return strconv.FormatInt(int64I, 10)
}

// 将驼峰转成下划线
func (f Function) humpToUnderline(value string) string {
	lenValue := len(value)
	newValue := ""
	for i := 0; i < lenValue; i++ {
		newBy := string(value[i])
		if i == 0 {
			if value[i] >= 65 && value[i] <= 90 {
				newBy = strings.ToLower(newBy)
			}
		} else {
			if value[i] >= 65 && value[i] <= 90 {
				newBy = strings.ToLower("_" + newBy)
			}
		}
		newValue += newBy
	}
	return newValue
}

// TimeParse 字符串转时间
func (f Function) TimeParse(date string) (time.Time, error) {
	formatAtByte := []byte("0000-00-00 00:00:00")
	copy(formatAtByte, []byte(date))
	return time.ParseInLocation("2006-01-02 15:04:05", string(formatAtByte), time.Local)
}

// InArray 判断某一个值是否含在切片之中
func (f Function) InArray(val interface{}, array interface{}) (exists bool, index int) {
	exists = false
	index = -1

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				index = i
				exists = true
				return
			}
		}
	}
	return
}

func (f Function) ValidData(
	data interface{},
	ruleKey string,
	fn func(validData interface{}, rule string) error,
) error {
	dataByte, _ := json.Marshal(data)
	var newData interface{}
	_ = json.Unmarshal(dataByte, &newData)
	return f.handleValidData(newData, ruleKey, ruleKey, fn)
}

func (f Function) handleValidData(
	data interface{},
	inputRule string,
	ruleKey string,
	fn func(validData interface{}, rule string) error,
) error {
	if data == nil {
		return fn(nil, ruleKey)
	}
	inputRuleList := strings.Split(inputRule, ".")
	dataType := reflect.ValueOf(data).Kind()
	nowRule := inputRuleList[0]
	if len(inputRuleList) > 1 {
		otherRule := strings.Join(inputRuleList[1:], ".")
		if dataType == reflect.Slice || dataType == reflect.Array {
			dataList := data.([]interface{})
			index, err := f.ParseInt(nowRule)
			if err == nil {
				if index > len(dataList)-1 {
					return fn(nil, ruleKey)
				}
				if err := f.handleValidData(dataList[index], otherRule, ruleKey, fn); err != nil {
					return err
				}
			} else {
				for key, val := range dataList {
					indexStr := f.formatNumber(key)
					dLen := len(ruleKey) - len(inputRule)
					newRuleKey := ruleKey[0:dLen] + indexStr + "." + otherRule
					if err := f.handleValidData(val, otherRule, newRuleKey, fn); err != nil {
						return err
					}
				}
			}
		} else if dataType == reflect.Map {
			dataMap := data.(map[string]interface{})
			if err := f.handleValidData(dataMap[nowRule], otherRule, ruleKey, fn); err != nil {
				return err
			}
		}
	} else {
		if dataType == reflect.Slice || dataType == reflect.Array {
			dataList := data.([]interface{})
			index, err := f.ParseInt(nowRule)
			if err == nil {
				if index > len(dataList)-1 {
					return fn(nil, ruleKey)
				}
				if err := fn(dataList[index], ruleKey); err != nil {
					return err
				}
			}
		} else if dataType == reflect.Map {
			dataMap := data.(map[string]interface{})
			if err := fn(dataMap[nowRule], ruleKey); err != nil {
				return err
			}
		}
	}
	return nil
}

func (f Function) ParseInt(value string) (int, error) {
	valueInt64, err := strconv.ParseInt(value, 10, 64)
	return int(valueInt64), err
}
