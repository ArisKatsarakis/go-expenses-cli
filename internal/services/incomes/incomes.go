package incomes

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/ArisKatsarakis/go-expenses-cli/db"
	"github.com/joho/godotenv"
)

func GetEnvVar(key string) string {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Errror: ", err.Error())
	}

	return os.Getenv(key)

}
func connectDb() (*sql.DB, error) {

	databaseConnectionURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		GetEnvVar("username"),
		GetEnvVar("password"),
		GetEnvVar("host"),
		GetEnvVar("port"),
		GetEnvVar("dbName"),
	)

	database, err := sql.Open("mysql", databaseConnectionURL)

	return database, err
}

func AddIncome() {
	var amount int64
	var name string
	var exec string
	fmt.Println("Inserting income: ")
	fmt.Println("Please give ammount: ")
	fmt.Scanf("%d", &amount)
	fmt.Printf("Inserted amount is : %d \n Press c for clear\n", amount)
	fmt.Println("Please give name of income ")
	fmt.Scanf("%s", &name)
	var params db.InsertIncomeParams
	params.Money = amount
	params.Name = name
	params.Timestamp = time.Now()
	database, err := connectDb()
	if err != nil {
		fmt.Println(err.Error())
	}
	connection := db.New(database)
	result, err := connection.InsertIncome(context.Background(), params)

	for {
		fmt.Println("Press c for clearing")
		fmt.Scanf("%s", &exec)
		if exec == "c" {
			break
		}

		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(result)
	}

}

func ListIncomes() {
	var exec string
	fmt.Println("Searching incomes")
	database, err := connectDb()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	connection := db.New(database)

	incomes, err := connection.AllIncomes(context.Background())
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for _, income := range incomes {
		fmt.Println("name: ", income.Name, " amount: ", income.Money)
	}
	for {
		fmt.Println("Press c for clearing")
		fmt.Scanf("%s", &exec)
		if exec == "c" {
			break
		}

	}
}
