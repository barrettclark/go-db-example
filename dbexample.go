package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	if database, _ := PGConnect(); database != nil {
		defer database.Close()
		fruits := GetFruit(database)
		fmt.Printf("%s\n", fruits)
	}
}

func PGConnect() (*sql.DB, error) {
	database, err := sql.Open("postgres", "dbname=railsconf2015 user=developer password=password sslmode=verify-full")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return database, nil
}

type Fruit struct {
	Id   int
	Name string
}

func GetFruit(database *sql.DB) []*Fruit {
	fruits := make([]*Fruit, 0)
	rows, err := database.Query("SELECT * FROM fruits")
	if err != nil {
		log.Fatalf("[x] Error when getting the list of fruit. Reason %s", err.Error())
	}
	for rows.Next() {
		fruit := toFruit(rows)
		fruits = append(fruits, fruit)
	}
	return fruits
}

func toFruit(rows *sql.Rows) *Fruit {
	var id int
	var name string
	rows.Scan(&id, &name)

	return &Fruit{
		Id:   id,
		Name: name,
	}
}
