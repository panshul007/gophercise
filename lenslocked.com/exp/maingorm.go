package main

import (
	//"database/sql"
	"fmt"

	"bufio"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
	"strings"
)

//const (
//	host   = "localhost"
//	port   = 5432
//	user   = "postgres"
//	dbname = "lenslocked_dev"
//)

type User1 struct {
	gorm.Model
	Name  string
	Email string `gorm:"not null;unique_index"`
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, user, dbname)
	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.DB().Ping()
	if err != nil {
		panic(err)
	}

	// Logging all SQL statements.
	db.LogMode(true)

	// Dropping the table to start fresh.
	// db.DropTableIfExists(&User{})

	db.AutoMigrate(&User{})

	name, email := getInfo()
	u := User{
		Name:  name,
		Email: email,
	}

	if err = db.Create(&u).Error; err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", u)
}

func getInfo() (name, email string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("What is ur name?")
	name, _ = reader.ReadString('\n')
	fmt.Println("What is ur email?")
	email, _ = reader.ReadString('\n')
	name = strings.TrimSpace(name)
	email = strings.TrimSpace(email)
	return name, email
}
