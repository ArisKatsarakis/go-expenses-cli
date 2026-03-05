package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ArisKatsarakis/go-expenses-cli/tutorial"
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
	db, err := sql.Open("mysql", databaseConnectionURL)
	if err != nil {
		return err
	}

	queries := tutorial.New(db)

	authors, err := queries.ListAuthors(ctx)
	if err != nil {
		return err
	}
	fmt.Println(authors)

	return nil
}

func main() {
	fmt.Println("Testing the main")
	err := runningQuery()
	if err != nil {
		fmt.Println(err.Error())
	}
}
