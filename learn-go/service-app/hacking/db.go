package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"log"
	"time"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "diwakar"
	password = "root"
	dbname   = "postgres"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("pgx", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	fmt.Println("connected")
	//Insert(db)
	//Insert2(ctx, db)
	//Update(ctx, db)
	//delete(ctx, db)
	//querySingleRecord(ctx, db)
	QueryMultipleRecords(ctx, db)
}

func Insert(db *sql.DB) {

	sqlStatement := `INSERT INTO users (age, email, first_name,last_name)
					VALUES ($1, $2, $3, $4)`

	//Exec is used if your sql statement would not return anything
	res, err := db.Exec(sqlStatement, 22, "abc@email.com", "abc", "abc")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(res.LastInsertId())

}

func Insert2(ctx context.Context, db *sql.DB) {
	sqlStatement := `INSERT INTO users (age, email, first_name,last_name)
					VALUES ($1, $2, $3, $4)
					RETURNING id,email
					`
	var (
		id    int
		email string
	)
	//QueryRowContext is used when sqlStatement returns one row back
	err := db.QueryRowContext(ctx, sqlStatement, 24, "abhishek@email.com", "abhishek", "arora").Scan(&id, &email)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(id, email)
}

func Update(ctx context.Context, db *sql.DB) {
	sqlStatement := `Update users 
                     Set last_name= $1
					 where id =$2;`

	res, err := db.ExecContext(ctx, sqlStatement, "xyz", 2)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(res.RowsAffected())

}

func delete(ctx context.Context, db *sql.DB) {

	sqlStatement := `Delete FROM users
                    where id =$1;
`

	res, err := db.ExecContext(ctx, sqlStatement, 2)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(res.RowsAffected())

}

func querySingleRecord(ctx context.Context, db *sql.DB) {
	sqlStatement := `Select id, email FROM users where id = $1;`

	var (
		id    int
		email string
	)
	err := db.QueryRowContext(ctx, sqlStatement, 1).Scan(&id, &email)

	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(id, email)
}

func QueryMultipleRecords(ctx context.Context, db *sql.DB) {
	q := `"Select id, email FROM users "`
	//exec the query
	//QueryContext // we expect the multiple rows back
	rows, err := db.QueryContext(ctx, q)

	if err != nil {
		log.Println(err)
		return
	}

	for rows.Next() {
		var (
			id    int
			email string
		)

		err = rows.Scan(&id, &email)
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Println(id, email)
	}

}
