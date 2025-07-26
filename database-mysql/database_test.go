package databasemysql

import (
	"context"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestExecContext(t *testing.T) {
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
