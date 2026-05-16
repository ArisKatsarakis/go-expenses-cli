package model

import (
	"os"

	"github.com/ArisKatsarakis/go-expenses-cli/internal/services/incomes"
)

type CommandExec struct {
	ExecuteName  string
	Name         string
	ExecutedFunc func()
}

var GlobalCmds = []CommandExec{
	{
		ExecuteName: "q",
		Name:        "Quit Application",
		ExecutedFunc: func() {
			os.Exit(0)
		},
	},
	{
		ExecuteName:  "i",
		Name:         "income add",
		ExecutedFunc: incomes.AddIncome,
	},
	{
		ExecuteName:  "l",
		Name:         "incomes list",
		ExecutedFunc: incomes.ListIncomes,
	},
}
