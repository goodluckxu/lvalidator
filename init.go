package lvalidator

import "lvalidator/lang"

var (
	Lang      lang.ZhCn
	RuleNotes map[string]string
)

func init() {
	Lang.Init()
	RuleNotes = map[string]string{}
}
