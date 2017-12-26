package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

//const (
//	host = "localhost"
//	port = 5432
//	user = "postgres"
//	dbname = "lenslocked_dev"
//)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, user, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//_, err = db.Exec(`
	//INSERT INTO users(name, email)
	//VALUES($1, $2) RETURNING id`, "Jon1 Calhoun1", "jon1@calhoun.io")

	var id int
	var name, email string
	//row := db.QueryRow(`
	//INSERT INTO users(name, email)
	//VALUES($1, $2) RETURNING id`,
	//	"Jon1 Calhoun1", "jon1@calhoun.io")

	row := db.QueryRow("Select id, name, email from users where id=$1", 1)
	err = row.Scan(&id, &name, &email)
	if err != nil {
		panic(err)
	}

	fmt.Printf("id: %d - name: %s - email: %s\n", id, name, email)
}
