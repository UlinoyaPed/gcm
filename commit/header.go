package commit

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
)

func Header() string {
	var Type string
	selectType := &survey.Select{
		Message: "选择头部Type:",
		Options: []string{"feat", "fix", "docs", "style", "refactor", "test", "chore"},
		Default: "feat",
	}
	survey.AskOne(selectType, &Type)

	var Scope string
	selectScope := &survey.Select{
		Message: "选择影响范围:",
		Options: []string{"a", "Other"},
	}
	survey.AskOne(selectScope, &Scope)
	if Scope == "Other" {
		inputScope := &survey.Input{
			Message: "输入Scpoe:",
		}
		survey.AskOne(inputScope, &Scope)
	}

	var Subject string
	inputSubject := &survey.Input{
		Message: "输入简短描述:",
	}
	survey.AskOne(inputSubject, &Subject)

	Header := fmt.Sprintf("%s(%s): %s", Type, Scope, Subject)
	fmt.Println(Header)
	return Header
}
