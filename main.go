package main

import (
	"bufio"
	"context"
	"database/sql"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/ArisKatsarakis/go-expenses-cli/db"
	_ "github.com/go-sql-driver/mysql"
)

func runningQuery() error {
	host := "linux197.papaki.gr"
	port := "3306"
	username := "cli_root"
	password := "katsar93@@"
	dbName := "accountingdb_cli"
	ctx := context.Background()
	databaseConnectionURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		username,
		password, // e.g. "super-secret-2026"
		host,
		port,
		dbName,
	)
	database, err := sql.Open("mysql", databaseConnectionURL)
	if err != nil {
		return err
	}

	queries := db.New(database)

	result, err := queries.InsertIncome(ctx, db.InsertIncomeParams{
		Name:      "Income_name",
		Money:     100,
		Timestamp: time.Now(),
	})
	if err != nil {
		return err
	}
	fmt.Println(result)

	results, err := queries.GetIncome(ctx, 1)
	if err != nil {
		return err
	}
	fmt.Println(results)

	return nil
}
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
		char, _, err := reader.ReadRune()
		if err != nil {
			fmt.Print("error reading \n")
		}
		switch char {
		case 'q':
			fmt.Println("Exit!")
			return
		}

	}
}
