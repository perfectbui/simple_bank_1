package main

import (
	db "LearningTransaction/simplebank_1/db/sqlc"
	"LearningTransaction/simplebank_1/utils"
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "myPassword"
	dbname   = "postgres"
)

var testQueries *db.Queries
var testDB *sql.DB
var z = 0

func CreateRandomAccount12() db.Accounts {
	arg := db.CreateAccountParams{
		Owner:    utils.RandomOwner(),
		Balance:  utils.RandomMoney(),
		Currency: utils.RandomCurrency(),
	}
	account, _ := testQueries.CreateAccount(context.Background(), arg)
	return account
}

func main() {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	testDB, err = sql.Open(dbDriver, psqlInfo)
	if err != nil {
		fmt.Println(err)
		return
	}
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}
	testQueries = db.New(testDB)
	store := db.NewStore(testDB)
	account1 := CreateRandomAccount12()
	account2 := CreateRandomAccount12()
	amount := int64(10)
	testcaseQty := 5
	for i := 0; i < testcaseQty; i++ {
		store.TransferTx(context.Background(), db.CreateTransferParams{FromAccountID: account1.ID, ToAccountID: account2.ID, Amount: amount})
	}

}
