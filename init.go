package lvalidator

import (
	"lang"
	"reflect"
)

var (
	Lang        lang.ZhCn
	RuleNotes   map[string]string
	ValidApiMap map[string]reflect.Value
)

func init() {
	Lang.Init()
	RuleNotes = map[string]string{}
}
