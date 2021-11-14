package lib

import (
	"fmt"
	"os/exec"
)

type Method struct {
	Topic       string
	Name        string
	Description string
	Source      string
}

type Technique struct {
	Name        string
	Category    string
	Description string
	Source      string
}

func GetRandomFromCategory(category string) {
	// if windows
	// PATH=$PATH:$(where python3)
	// if linux
	// PATH=$PATH:$(which python3)
	cmd := exec.Command("python3", "cockroachdb/cliGetRandom.py", category)
	out, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return
	}

	fmt.Println(string(out))
}
