package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	database, err := sql.Open("postgres", "dbname=railsconf2015 user=developer password=password sslmode=verify-full")
	defer database.Close()
	if err != nil {
		log.Fatal(err)
	}
	rows, err := database.Query("SELECT * FROM fruits")
	if err != nil {
		log.Fatalf("[x] Error when getting the list of fruit. Reason %s", err.Error())
	}
	var id int
	var name string
	for rows.Next() {
		rows.Scan(&id, &name)
		fruit := &Fruit{id, name}
		fmt.Printf("%+v\n", fruit)
	}
}

type Fruit struct {
	Id   int
	Name string
}

/*
func toFruit(rows *sql.Rows) *Fruit {
	var id int
	var name string
	rows.Scan(&id, &name)

	return &Fruit{
		Id:   id,
		Name: name,
	}
}
*/
