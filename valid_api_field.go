package lvalidator

import (
	"errors"
	"strings"
)

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
