package databasemysql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
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

func TestDatatypeSql(t *testing.T) {
	db := SetConnection()
	defer db.Close()

	ctx := context.Background()

	payloads := []Customer{
		{
			name:      "Agus",
			email:     "agus@gmail.com",
			balance:   2322,
			rating:    sql.NullFloat64{Float64: 23.131, Valid: true},
			birthdate: "2025-01-01",
			isMarried: 0,
		},
		{
			name:      "Ahmad",
			email:     "ahmad@gmail.com",
			balance:   521,
			rating:    sql.NullFloat64{Float64: 5.21, Valid: true},
			birthdate: "2025-01-01",
			isMarried: 0,
		},
	}

	err := Create(ctx, db, &payloads)

	if err != nil {
		panic(err)
	}
}

type Customer struct {
	id, name, email string
	balance         int
	rating          sql.NullFloat64
	birthdate       string
	isMarried       int
	createdAt       string
}

func TestScanNew(t *testing.T) {
	db := SetConnection()
	defer db.Close()

	ctx := context.Background()

	rows, err := db.QueryContext(ctx, "select id, name, email, balance, rating, birthdate, is_married, created_at from customers")

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var customers []Customer

	for rows.Next() {
		var cust Customer

		err := rows.Scan(&cust.id, &cust.name, &cust.email, &cust.balance, &cust.rating, &cust.birthdate, &cust.isMarried, &cust.createdAt)
		if err != nil {
			panic(err)
		}

		customers = append(customers, cust)
	}

	for _, v := range customers {
		fmt.Println(v)
	}
}

func Create(ctx context.Context, db *sql.DB, payloads *[]Customer) error {
	if len(*payloads) == 0 {
		return errors.New("Payloads cannot be empty")
	}

	for _, val := range *payloads {
		randomID := strconv.Itoa(rand.Intn(1000))
		query := "INSERT INTO customers (id, name, email, balance, rating, birthdate, is_married) VALUES (?, ?, ?, ?, ?, ?, ?)"

		_, err := db.ExecContext(ctx, query, randomID, val.name, val.email, val.balance, val.rating, val.birthdate, val.isMarried)

		if err != nil {
			return errors.New(err.Error())
		}
	}

	return nil
}

func TestPrepareStmt(t *testing.T) {
	db := SetConnection()

	defer db.Close()

	ctx := context.Background()

	stmt, err := db.PrepareContext(ctx, "INSERT INTO comments (email, comment) VALUES (?, ?)")

	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	for i := 1; i <= 10; i++ {
		res, err := stmt.ExecContext(ctx, "asep@gmail.com", "halo ini asep ("+strconv.Itoa(i)+")")

		if err != nil {
			panic(err)
		}

		lastID, err := res.LastInsertId()

		if err != nil {
			panic(err)
		}

		fmt.Println("Last ID:", lastID)
	}
}

func TestDBTransaction(t *testing.T) {
	db := SetConnection()

	defer db.Close()

	ctx := context.Background()

	tx, err := db.Begin()

	if err != nil {
		panic(err)
	}

	for i := 1; i <= 10; i++ {
		res, err := tx.ExecContext(ctx, "INSERT INTO comments (email, comment) VALUES (?, ?)", "asep@gmail.com", "halo ini asep ("+strconv.Itoa(i)+")")

		if err != nil {
			tx.Rollback()
			panic(err)
		}

		lastID, err := res.LastInsertId()

		if err != nil {
			tx.Rollback()
			panic(err)
		}

		fmt.Println("Last ID:", lastID)
	}

	tx.Commit()
}
