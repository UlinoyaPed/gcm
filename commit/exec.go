package commit

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

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
