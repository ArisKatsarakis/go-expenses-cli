package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"

	"github.com/ArisKatsarakis/go-expenses-cli/internal/utils"
	_ "github.com/go-sql-driver/mysql"
)

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		clearScreen()
		fmt.Println("+++++++++++INCOME/EXPENSE+++++++++++")
		utils.RenderGlobalCmds()
		char, _, err := reader.ReadRune()
		if err != nil {
			fmt.Print("error reading \n")
		}
		utils.ExecuteGlobalCmd(char)
	}
}
