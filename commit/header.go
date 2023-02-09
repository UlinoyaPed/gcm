package commit

import (
	"fmt"
	"os/exec"

	"github.com/AlecAivazis/survey/v2"

	"bytes"
	"log"
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
	inputSubject := &survey.Input{
		Message: "输入简短描述:",
	}
	survey.AskOne(inputSubject, &Subject)
	fmt.Println(Subject)

	FullCommitMessage := fmt.Sprintf("%s(%s): %s", Header, Scope, Subject)

	// cmd := exec.Command("git commit", "-m", FullCommitMessage)
	// out, _ := cmd.CombinedOutput()
	// fmt.Println(out)

	ExecCommand("git commit", "-m"+FullCommitMessage)
}

func ExecCommand(name string, args ...string) {
	cmd := exec.Command(name, args...) // 拼接参数与命令

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	var err error

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err = cmd.Run(); err != nil {
		log.Println(err)
	}

	fmt.Print(stdout.String())
	fmt.Print(stderr.String())
}
