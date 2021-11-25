package lvalidator

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"regexp"
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
		switch rule.(type) {
		case string:
			if rule.(string) != "" {
				for _, v := range strings.Split(rule.(string), "|") {
					list = append(list, v)
				}
			}
		case []interface{}:
			for _, v := range rule.([]interface{}) {
				list = append(list, v)
			}
		default:
			return nil, errors.New("规格类型不正确")
		}
		sort := 0
		notes := ""
		newList := []interface{}{}
		for _, v := range list {
			if v == nil {
				return nil, errors.New("规格类型不正确")
			}
			switch v.(type) {
			case string:
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
					sort, _ = strconv.Atoi(vList[1])
				case "notes":
					notes = vList[1]
				}
			case func(interface{}) error, func(interface{}, string) error:
				newList = append(newList, v)
			default:
				return nil, errors.New("规格类型不正确")
			}
		}
		RuleNotes[key] = notes
		ruleMap["key"] = key
		ruleMap["sort"] = sort
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
// sortSliceMapStringInterface 排序[]map[string]interface{}
func (f Function) sortSliceMapStringInterface(data []map[string]interface{}, args ...string) (rs []map[string]interface{}) {
	rs = []map[string]interface{}{}
	if len(args) == 0 {
		return
	}
	for _, v := range data {
		rs = append(rs, v)
	}
	lenData := len(rs)
	for i := 0; i < lenData-1; i++ {
		for j := 0; j < lenData-1-i; j++ {
			for _, arg := range args {
				argList := strings.Split(arg, " ")
				field := argList[0]
				sortType := "asc"
				if len(argList) > 1 {
					sortType = argList[1]
				}
				jY := rs[j][field]
				jYAdd := rs[j+1][field]
				var floatJ float64
				var floatJAdd float64
				isFloat := false
				switch jY.(type) {
				// case uint, uint8, uint16, uint32, uint64, int, int8, int16, int32, int64, float32, float64:
				case uint:
					floatJ = float64(jY.(uint))
					floatJAdd = float64(jYAdd.(uint))
					isFloat = true
				case uint8:
					floatJ = float64(jY.(uint8))
					floatJAdd = float64(jYAdd.(uint8))
					isFloat = true
				case uint16:
					floatJ = float64(jY.(uint16))
					floatJAdd = float64(jYAdd.(uint16))
					isFloat = true
				case uint32:
					floatJ = float64(jY.(uint32))
					floatJAdd = float64(jYAdd.(uint32))
					isFloat = true
				case uint64:
					floatJ = float64(jY.(uint64))
					floatJAdd = float64(jYAdd.(uint64))
					isFloat = true
				case int:
					floatJ = float64(jY.(int))
					floatJAdd = float64(jYAdd.(int))
					isFloat = true
				case int8:
					floatJ = float64(jY.(int8))
					floatJAdd = float64(jYAdd.(int8))
					isFloat = true
				case int16:
					floatJ = float64(jY.(int16))
					floatJAdd = float64(jYAdd.(int16))
					isFloat = true
				case int32:
					floatJ = float64(jY.(int32))
					floatJAdd = float64(jYAdd.(int32))
					isFloat = true
				case int64:
					floatJ = float64(jY.(int64))
					floatJAdd = float64(jYAdd.(int64))
					isFloat = true
				case float32:
					floatJ = float64(jY.(float32))
					floatJAdd = float64(jYAdd.(float32))
					isFloat = true
				case float64:
					floatJ = jY.(float64)
					floatJAdd = jYAdd.(float64)
					isFloat = true
				}
				compareStr := 0
				if isFloat {
					if floatJ > floatJAdd {
						compareStr = 1
					} else if floatJ < floatJAdd {
						compareStr = -1
					}
				} else {
					strJ := f.formatNumber(jY)
					strJAdd := f.formatNumber(jYAdd)
					compareStr = strings.Compare(strJ, strJAdd)
				}
				if compareStr == 0 {
					continue
				}
				if sortType == "asc" && compareStr == 1 {
					temp := rs[j]
					rs[j] = rs[j+1]
					rs[j+1] = temp
				} else if sortType == "desc" && compareStr == -1 {
					temp := rs[j]
					rs[j] = rs[j+1]
					rs[j+1] = temp
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
	switch i.(type) {
	case int:
		int64I = int64(i.(int))
	case int8:
		int64I = int64(i.(int8))
	case int16:
		int64I = int64(i.(int16))
	case int32:
		int64I = int64(i.(int32))
	case int64:
		int64I = i.(int64)
	case uint:
		int64I = int64(i.(uint))
	case uint8:
		int64I = int64(i.(uint8))
	case uint16:
		int64I = int64(i.(uint16))
	case uint32:
		int64I = int64(i.(uint32))
	case uint64:
		int64I = int64(i.(uint64))
	case float32:
		int64I = int64(i.(float32))
	case float64:
		int64I = int64(i.(float64))
	case string:
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

	arrayValue := reflect.ValueOf(array)
	if arrayValue.Kind() != reflect.Slice {
		return
	}
	for i := 0; i < arrayValue.Len(); i++ {
		if reflect.DeepEqual(val, arrayValue.Index(i).Interface()) {
			index = i
			exists = true
			return
		}
	}
	return
}

// 验证数据
func (f Function) ValidData(
	data interface{},
	ruleKey string,
	fn func(validData interface{}, validNotes, validRule string) error,
) error {
	dataByte, _ := json.Marshal(reflect.ValueOf(data).Elem().Interface())
	var newData interface{}
	_ = json.Unmarshal(dataByte, &newData)
	return f.handleValidData(&newData, ruleKey, ruleKey, func(validData interface{}, rule string) error {
		return fn(validData, f.GetNotes(ruleKey, rule), rule)
	})
}

// 验证两个字段数据
func (f Function) ValidDoubleData(
	data interface{},
	ruleKey string,
	compareKey string,
	fn func(validData, compareData interface{}, validNotes, compareNotes string) error,
) error {
	dataByte, _ := json.Marshal(data)
	var newData interface{}
	_ = json.Unmarshal(dataByte, &newData)
	starInfo := f.GetTwoFieldStarInfo(ruleKey, compareKey)
	return f.handleValidData(newData, ruleKey, ruleKey, func(validData interface{}, rule string) error {
		return f.handleValidData(newData, compareKey, compareKey, func(compareData interface{}, compareRule string) error {
			if !f.IsTwoFieldCompare(rule, compareRule, starInfo) {
				return nil
			}
			validNotes := f.GetNotes(ruleKey, rule)
			compareNotes := f.GetNotes(compareKey, compareRule)
			return fn(validData, compareData, validNotes, compareNotes)
		})
	})
}

// 处理验证数据
func (f Function) handleValidData(
	newData interface{},
	inputRule string,
	ruleKey string,
	fn func(validData interface{}, rule string) error,
) error {
	if newData == nil {
		return fn(nil, ruleKey)
	}
	dataValue := reflect.ValueOf(newData)
	if dataValue.Kind() == reflect.Ptr {
		dataValue = dataValue.Elem()
	}
	data := dataValue.Interface()
	if data == nil {
		return fn(nil, ruleKey)
	}
	inputRuleList := strings.Split(inputRule, ".")
	nowRule := inputRuleList[0]
	if len(inputRuleList) > 1 {
		otherRule := strings.Join(inputRuleList[1:], ".")
		switch data.(type) {
		case []interface{}:
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
		case map[string]interface{}:
			dataMap := data.(map[string]interface{})
			if err := f.handleValidData(dataMap[nowRule], otherRule, ruleKey, fn); err != nil {
				return err
			}
		}
	} else {
		switch data.(type) {
		case []interface{}:
			dataList := data.([]interface{})
			index, err := f.ParseInt(nowRule)
			if err == nil {
				if index > len(dataList)-1 {
					return fn(nil, ruleKey)
				}
				if err := fn(dataList[index], ruleKey); err != nil {
					return err
				}
			} else {
				for key, val := range dataList {
					indexStr := f.formatNumber(key)
					dLen := len(ruleKey) - len(inputRule)
					newRuleKey := ruleKey[0:dLen] + indexStr
					if err := fn(val, newRuleKey); err != nil {
						return err
					}
				}
			}
		case map[string]interface{}:
			dataMap := data.(map[string]interface{})
			if err := fn(dataMap[nowRule], ruleKey); err != nil {
				return err
			}
		}
	}
	return nil
}

// 解析int类型
func (f Function) ParseInt(value string) (int, error) {
	valueInt64, err := strconv.ParseInt(value, 10, 64)
	return int(valueInt64), err
}

// 获取两个字段最小的*数量
func (f Function) GetTwoFieldStarInfo(ruleOne, ruleTwo string) map[string]interface{} {
	ruleOneList := strings.Split(ruleOne, ".")
	ruleTwoList := strings.Split(ruleTwo, ".")
	starOneLen := 0
	starOneIndexList := []int{}
	for i := 0; i < len(ruleOneList); i++ {
		if ruleOneList[i] == "*" {
			starOneLen++
			starOneIndexList = append(starOneIndexList, i)
		}
	}
	startTwoLen := 0
	starTwoIndexList := []int{}
	for i := 0; i < len(ruleTwoList); i++ {
		if ruleTwoList[i] == "*" {
			startTwoLen++
			starTwoIndexList = append(starTwoIndexList, i)
		}
	}
	starLen := starOneLen
	if startTwoLen < starLen {
		starLen = startTwoLen
	}
	return map[string]interface{}{
		"one": starOneIndexList[0:starLen],
		"two": starTwoIndexList[0:starLen],
	}
}

// 判断两个字段是否可比较
func (f Function) IsTwoFieldCompare(ruleOne, ruleTwo string, starInfo map[string]interface{}) bool {
	starOneIndexList := starInfo["one"].([]int)
	starTwoIndexList := starInfo["two"].([]int)
	ruleOneList := strings.Split(ruleOne, ".")
	ruleTwoList := strings.Split(ruleTwo, ".")
	oneCompare := []string{}
	for i := 0; i < len(ruleOneList); i++ {
		if bl, _ := f.InArray(i, starOneIndexList); bl {
			oneCompare = append(oneCompare, ruleOneList[i])
		}
	}
	twoCompare := []string{}
	for i := 0; i < len(ruleTwoList); i++ {
		if bl, _ := f.InArray(i, starTwoIndexList); bl {
			twoCompare = append(twoCompare, ruleTwoList[i])
		}
	}
	return reflect.DeepEqual(oneCompare, twoCompare)
}

// 是否是integer类型
func (f Function) IsInteger(value string) bool {
	reg := regexp.MustCompile(`^\d*$`)
	return reg.MatchString(value)
}

// 获取注释
func (f Function) GetNotes(ruleKey, rule string) string {
	notes := RuleNotes[ruleKey]
	if notes == "" {
		notes = rule
	}
	return notes
}

/**
 * 验证时间
 */
func (f Function) ValidDate(date string, format string) (err error) {
	err = fmt.Errorf("日期格式和值不匹配，格式：%s, 值：%s", format, date)
	format = strings.ReplaceAll(format, "Y", "YYYY")
	format = strings.ReplaceAll(format, "m", "mm")
	format = strings.ReplaceAll(format, "d", "dd")
	format = strings.ReplaceAll(format, "H", "HH")
	format = strings.ReplaceAll(format, "i", "ii")
	format = strings.ReplaceAll(format, "s", "ss")
	if len(date) != len(format) {
		return
	}
	if err := f.validSingleDate("YYYY", &date, &format); err != nil {
		return err
	}
	if err := f.validSingleDate("mm", &date, &format); err != nil {
		return err
	}
	if err := f.validSingleDate("dd", &date, &format); err != nil {
		return err
	}
	if err := f.validSingleDate("HH", &date, &format); err != nil {
		return err
	}
	if err := f.validSingleDate("ii", &date, &format); err != nil {
		return err
	}
	if err := f.validSingleDate("ss", &date, &format); err != nil {
		return err
	}
	if date != format {
		return
	}
	return nil
}

func (f Function) validSingleDate(single string, date, format *string) error {
	lenSingle := len(single)
	newDate := *date
	newFormat := *format
	for {
		index := strings.Index(newFormat, single)
		if index == -1 {
			break
		}
		validDate := newDate[index : index+lenSingle]
		switch single {
		case "YYYY":
			if !regexp.MustCompile(`\d{4}`).MatchString(validDate) {
				return fmt.Errorf("年格式和值不匹配，格式：Y, 值：%s", validDate)
			}
		case "mm":
			validInt64, _ := strconv.ParseInt(validDate, 10, 64)
			if validInt64 < 1 || validInt64 > 12 {
				return fmt.Errorf("月格式和值不匹配，格式：m, 值：%s", validDate)
			}
		case "dd":
			validInt64, _ := strconv.ParseInt(validDate, 10, 64)
			if validInt64 < 1 || validInt64 > 31 {
				return fmt.Errorf("日格式和值不匹配，格式：d, 值：%s", validDate)
			}
		case "HH":
			validInt64, _ := strconv.ParseInt(validDate, 10, 64)
			if validInt64 < 0 || validInt64 > 23 {
				return fmt.Errorf("时格式和值不匹配，格式：H, 值：%s", validDate)
			}
		case "ii":
			validInt64, _ := strconv.ParseInt(validDate, 10, 64)
			if validInt64 < 0 || validInt64 > 59 {
				return fmt.Errorf("分格式和值不匹配，格式：i, 值：%s", validDate)
			}
		case "ss":
			validInt64, _ := strconv.ParseInt(validDate, 10, 64)
			if validInt64 < 0 || validInt64 > 59 {
				return fmt.Errorf("秒格式和值不匹配，格式：s, 值：%s", validDate)
			}
		}
		newFormat = newFormat[0:index] + newFormat[index+lenSingle:]
		newDate = newDate[0:index] + newDate[index+lenSingle:]
	}
	*date = newDate
	*format = newFormat
	return nil
}

// 比较两个数是否相等
func (f Function) IsEqualData(dataOne, dataTwo interface{}) bool {
	switch dataOne.(type) {
	case string:
		if compareDataString, bl := dataTwo.(string); bl {
			if dataOne.(string) == compareDataString {
				return true
			}
		}
	case float64:
		if compareDataFloat64, bl := dataTwo.(float64); bl {
			if dataOne.(float64) == compareDataFloat64 {
				return true
			}
		}
	case bool:
		if compareDataBool, bl := dataTwo.(bool); bl {
			if dataOne.(bool) == compareDataBool {
				return true
			}
		}
	default:
		if reflect.DeepEqual(dataOne, dataTwo) {
			return true
		}
	}
	return false
}
