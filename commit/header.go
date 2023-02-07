package commit

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
)

func Header() {
	var Header string
	selectHeader := &survey.Select{
		Message: "选择头部Type:",
		Options: []string{"feat", "fix", "docs", "style", "refactor", "test", "chore"},
		Default: "feat",
	}
	survey.AskOne(selectHeader, &Header)
	fmt.Printf("Type:%s", Header)

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
	fmt.Println(Header + Scope)

	var Subject string
}
