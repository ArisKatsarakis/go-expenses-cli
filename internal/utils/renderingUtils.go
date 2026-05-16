package utils

import (
	"fmt"

	"github.com/ArisKatsarakis/go-expenses-cli/internal/model"
)

func RenderGlobalCmds() {
	cmds := model.GlobalCmds
	for _, c := range cmds {
		fmt.Print(c.Name, " : ", c.ExecuteName, "\n")
	}
}

func ExecuteGlobalCmd(cmd rune) {
	cmds := model.GlobalCmds
	for _, c := range cmds {
		if c.ExecuteName == string(cmd) {
			c.ExecutedFunc()
		}
	}

}
