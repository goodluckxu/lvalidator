package lvalidator

import (
	"errors"
	"strconv"
	"strings"
)

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