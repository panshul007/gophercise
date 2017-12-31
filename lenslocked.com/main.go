package main

import (
	"net/http"

	"fmt"
	"github.com/gorilla/mux"
	"gophercise/lenslocked.com/controllers"
	"gophercise/lenslocked.com/models"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "postgres"
	dbname = "lenslocked_dev"
)

func main() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, user, dbname)
	services, err := models.NewServices(psqlInfo)
	must(err)

	// TODO: fix this
	//defer us.Close()
	//us.AutoMigrate()
	//us.DestructiveReset()

	staticC := controllers.NewStatic()
	usersC := controllers.NewUsers(services.User)

	r := mux.NewRouter()
	r.Handle("/", staticC.Home).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")
	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")
	// Using Handle when an interface is passed which implements the ServeHTTP method
	r.Handle("/login", usersC.LoginView).Methods("GET")
	// Using HandleFunc when a method reference is passed
	r.HandleFunc("/login", usersC.Login).Methods("POST")
	r.HandleFunc("/cookietest", usersC.CookieTest).Methods("GET")
	fmt.Println("Starting the server at port: 3000...")
	http.ListenAndServe(":3000", r)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
