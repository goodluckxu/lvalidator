package lvalidator

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// 验证字段条件满足通过
func (v validApi) ValidConditionField(data interface{}, ruleKey string, ruleValue string) error {
	rs := errors.New(fmt.Sprintf("valid_condition_field:%s不是一个正确的规则", ruleValue))
	dataByte, _ := json.Marshal(data)
	var newData interface{}
	_ = json.Unmarshal(dataByte, &newData)
	ruleValueList := strings.Split(ruleValue, ",")
	if len(ruleValueList) != 2 {
		return rs
	}
	whereStringList := strings.Split(ruleValueList[1], "&")
	starInfo := Func.GetTwoFieldStarInfo(ruleKey, ruleValueList[0])
	return Func.handleValidData(newData, ruleKey, ruleKey, func(validData interface{}, rule string) error {
		return Func.handleValidData(newData, ruleValueList[0], ruleValueList[0], func(compareData interface{}, compareRule string) error {
			if !Func.IsTwoFieldCompare(rule, compareRule, starInfo) {
				return nil
			}
			isValid := true
			for _, whereString := range whereStringList {
				whereList := strings.Split(whereString, " ")
				if len(whereList) != 2 {
					return rs
				}
				switch whereList[0] {
				case "=":
					if strings.Compare(Func.formatNumber(compareData), whereList[1]) != 0 {
						isValid = false
					}
				case ">":
					compareFloat64, compareErr := strconv.ParseFloat(Func.formatNumber(compareData), 64)
					whereFloat64, whereErr := strconv.ParseFloat(Func.formatNumber(whereList[1]), 64)
					if compareErr != nil || whereErr != nil {
						return rs
					}
					if compareFloat64 <= whereFloat64 {
						isValid = false
					}
				case ">=":
					compareFloat64, compareErr := strconv.ParseFloat(Func.formatNumber(compareData), 64)
					whereFloat64, whereErr := strconv.ParseFloat(Func.formatNumber(whereList[1]), 64)
					if compareErr != nil || whereErr != nil {
						return rs
					}
					if compareFloat64 < whereFloat64 {
						isValid = false
					}
				case "<":
					compareFloat64, compareErr := strconv.ParseFloat(Func.formatNumber(compareData), 64)
					whereFloat64, whereErr := strconv.ParseFloat(Func.formatNumber(whereList[1]), 64)
					if compareErr != nil || whereErr != nil {
						return rs
					}
					if compareFloat64 >= whereFloat64 {
						isValid = false
					}
				case "<=":
					compareFloat64, compareErr := strconv.ParseFloat(Func.formatNumber(compareData), 64)
					whereFloat64, whereErr := strconv.ParseFloat(Func.formatNumber(whereList[1]), 64)
					if compareErr != nil || whereErr != nil {
						return rs
					}
					if compareFloat64 > whereFloat64 {
						isValid = false
					}
				case "in":
					stringList := strings.Split(whereList[1], ";")
					if bl, _ := Func.InArray(Func.formatNumber(compareData), stringList); !bl {
						isValid = false
					}
				default:
					return rs
				}
			}
			if isValid {
				return errors.New("")
			}
			// nil时需要验证，errors空字符串是通过
			return nil
		})
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
		if Func.IsEqualData(validData, compareData) {
			return nil
		}
		return rs
	})
}

// 大于字段
func (v validApi) GtField(data interface{}, ruleKey string, ruleValue string) error {
	return Func.ValidDoubleData(data, ruleKey, ruleValue, func(validData, compareData interface{}, validNotes, compareNotes string) error {
		info := strings.ReplaceAll(Lang.GtField, "{ruleKey}", validNotes)
		info = strings.ReplaceAll(info, "{ruleValue}", compareNotes)
		rs := errors.New(info)
		if validData == nil {
			return nil
		}
		switch validData.(type) {
		case float64:
			compareDataFloat64, _ := compareData.(float64)
			if validData.(float64) > compareDataFloat64 {
				return nil
			}
		case string:
			if dataFloat64, err := strconv.ParseFloat(validData.(string), 64); err == nil {
				compareDataFloat64, _ := compareData.(float64)
				if dataFloat64 > compareDataFloat64 {
					return nil
				}
			}
		}
		return rs
	})
}

// 大于等于字段
func (v validApi) GteField(data interface{}, ruleKey string, ruleValue string) error {
	return Func.ValidDoubleData(data, ruleKey, ruleValue, func(validData, compareData interface{}, validNotes, compareNotes string) error {
		info := strings.ReplaceAll(Lang.GteField, "{ruleKey}", validNotes)
		info = strings.ReplaceAll(info, "{ruleValue}", compareNotes)
		rs := errors.New(info)
		if validData == nil {
			return nil
		}
		switch validData.(type) {
		case float64:
			compareDataFloat64, _ := compareData.(float64)
			if validData.(float64) >= compareDataFloat64 {
				return nil
			}
		case string:
			if dataFloat64, err := strconv.ParseFloat(validData.(string), 64); err == nil {
				compareDataFloat64, _ := compareData.(float64)
				if dataFloat64 >= compareDataFloat64 {
					return nil
				}
			}
		}
		return rs
	})
}

// 小于字段
func (v validApi) LtField(data interface{}, ruleKey string, ruleValue string) error {
	return Func.ValidDoubleData(data, ruleKey, ruleValue, func(validData, compareData interface{}, validNotes, compareNotes string) error {
		info := strings.ReplaceAll(Lang.LtField, "{ruleKey}", validNotes)
		info = strings.ReplaceAll(info, "{ruleValue}", compareNotes)
		rs := errors.New(info)
		if validData == nil {
			return nil
		}
		switch validData.(type) {
		case float64:
			compareDataFloat64, _ := compareData.(float64)
			if validData.(float64) < compareDataFloat64 {
				return nil
			}
		case string:
			if dataFloat64, err := strconv.ParseFloat(validData.(string), 64); err == nil {
				compareDataFloat64, _ := compareData.(float64)
				if dataFloat64 < compareDataFloat64 {
					return nil
				}
			}
		}
		return rs
	})
}

// 小于等于字段
func (v validApi) LteField(data interface{}, ruleKey string, ruleValue string) error {
	return Func.ValidDoubleData(data, ruleKey, ruleValue, func(validData, compareData interface{}, validNotes, compareNotes string) error {
		info := strings.ReplaceAll(Lang.LteField, "{ruleKey}", validNotes)
		info = strings.ReplaceAll(info, "{ruleValue}", compareNotes)
		rs := errors.New(info)
		if validData == nil {
			return nil
		}
		switch validData.(type) {
		case float64:
			compareDataFloat64, _ := compareData.(float64)
			if validData.(float64) <= compareDataFloat64 {
				return nil
			}
		case string:
			if dataFloat64, err := strconv.ParseFloat(validData.(string), 64); err == nil {
				compareDataFloat64, _ := compareData.(float64)
				if dataFloat64 <= compareDataFloat64 {
					return nil
				}
			}
		}
		return rs
	})
}

// 日期大于字段
func (v validApi) DateGtField(data interface{}, ruleKey string, ruleValue string) error {
	return Func.ValidDoubleData(data, ruleKey, ruleValue, func(validData, compareData interface{}, validNotes, compareNotes string) error {
		compareDataString, bl := compareData.(string)
		ruleValueTime, err := Func.TimeParse(compareDataString)
		if !bl || err != nil {
			return errors.New(fmt.Sprintf("date_gt_field:%s字段不是日期格式", ruleValue))
		}
		info := strings.ReplaceAll(Lang.DateGtField, "{ruleKey}", validNotes)
		info = strings.ReplaceAll(info, "{ruleValue}", compareNotes)
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

// 日期大于等于字段
func (v validApi) DateGteField(data interface{}, ruleKey string, ruleValue string) error {
	return Func.ValidDoubleData(data, ruleKey, ruleValue, func(validData, compareData interface{}, validNotes, compareNotes string) error {
		compareDataString, bl := compareData.(string)
		ruleValueTime, err := Func.TimeParse(compareDataString)
		if !bl || err != nil {
			return errors.New(fmt.Sprintf("date_gte_field:%s字段不是日期格式", ruleValue))
		}
		info := strings.ReplaceAll(Lang.DateGteField, "{ruleKey}", validNotes)
		info = strings.ReplaceAll(info, "{ruleValue}", compareNotes)
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

// 日期小于字段
func (v validApi) DateLtField(data interface{}, ruleKey string, ruleValue string) error {
	return Func.ValidDoubleData(data, ruleKey, ruleValue, func(validData, compareData interface{}, validNotes, compareNotes string) error {
		compareDataString, bl := compareData.(string)
		ruleValueTime, err := Func.TimeParse(compareDataString)
		if !bl || err != nil {
			return errors.New(fmt.Sprintf("date_lt_field:%s字段不是日期格式", ruleValue))
		}
		info := strings.ReplaceAll(Lang.DateLtField, "{ruleKey}", validNotes)
		info = strings.ReplaceAll(info, "{ruleValue}", compareNotes)
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

// 日期小于等于字段
func (v validApi) DateLteField(data interface{}, ruleKey string, ruleValue string) error {
	return Func.ValidDoubleData(data, ruleKey, ruleValue, func(validData, compareData interface{}, validNotes, compareNotes string) error {
		compareDataString, bl := compareData.(string)
		ruleValueTime, err := Func.TimeParse(compareDataString)
		if !bl || err != nil {
			return errors.New(fmt.Sprintf("date_lte_field:%s字段不是日期格式", ruleValue))
		}
		info := strings.ReplaceAll(Lang.DateLteField, "{ruleKey}", validNotes)
		info = strings.ReplaceAll(info, "{ruleValue}", compareNotes)
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
