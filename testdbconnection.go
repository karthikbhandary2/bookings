package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	//connect
	conn, err := sql.Open("pgx", "host=localhost port=5432 dbname=bookings user=Karthik password=****")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	log.Println("Connected to database")
	//ping
	err = conn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("pinged the database")

	err = getAllRows(conn)
	if err != nil {
		log.Fatal(err)
	}
	//insert 
	query := `insert into bookings (first_name, last_name) values ($1, $2)`
	_, err = conn.Exec(query, "Karthik", "Bhandary")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a row")

	err = getAllRows(conn)
	if err != nil {
		log.Fatal(err)
	}
	//update
	stmt := `update bookings set first_name = $1 where last_name = $2`
	_, err = conn.Exec(stmt, "Ujjwal", "Uchiha")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Updated a row")

	err = getAllRows(conn)
	if err != nil {
		log.Fatal(err)
	}
	//get one row
	var firstName, lastName string
	var id int
	query = `select id, first_name, last_name from bookings where id = $1`
	row := conn.QueryRow(query, 1)
	err = row.Scan(&id, &firstName, &lastName)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Row ", id, firstName, lastName)

	//delete
	query = `delete from bookings where id = $1`
	_, err = conn.Exec(query, 1)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Deleted a row")
	//get all rows
	err = getAllRows(conn)
	if err != nil {
		log.Fatal(err)
	}

}

func getAllRows(conn *sql.DB) error {
	rows, err := conn.Query("select * from bookings")
	if err != nil {
		log.Println(err)
	}

	defer rows.Close()

	var firstName, lastName string
	var id int

	for rows.Next() {
		err := rows.Scan(&id, &firstName, &lastName)
		if err != nil {
			log.Println(err)
			return err
		}

		fmt.Println(id, firstName, lastName)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Println("---------------------")
	return nil
}