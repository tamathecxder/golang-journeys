package databasemysql

import (
	"context"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestExecSql(t *testing.T) {
	db := SetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "insert into customers (id, name) values (2323, 'Mamat')"
	_, err := db.ExecContext(ctx, query)

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully created customer!")
}

func TestQuerySql(t *testing.T) {
	db := SetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "select id, name from customers"
	rows, err := db.QueryContext(ctx, query)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id, name string

		err := rows.Scan(&id, &name)

		if err != nil {
			panic(err)
		}

		fmt.Println("ID: ", id)
		fmt.Println("Name: ", name)
	}
}
