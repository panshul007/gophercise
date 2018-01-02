package main

import (
	"net/http"

	"fmt"

	"github.com/gorilla/mux"
	"gophercise/lenslocked.com/controllers"
	"gophercise/lenslocked.com/middleware"
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

	defer services.Close()
	services.AutoMigrate()
	//us.DestructiveReset()

	r := mux.NewRouter()
	staticC := controllers.NewStatic()
	usersC := controllers.NewUsers(services.User)
	galleriesC := controllers.NewGalleries(services.Gallery, services.Image, r)

	userMw := middleware.User{UserService: services.User}
	requireUserMw := middleware.RequireUser{User: userMw}

	r.Handle("/", staticC.Home).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")
	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")
	// Using Handle when an interface is passed which implements the ServeHTTP method
	r.Handle("/login", usersC.LoginView).Methods("GET")
	// Using HandleFunc when a method reference is passed
	r.HandleFunc("/login", usersC.Login).Methods("POST")
	//r.HandleFunc("/cookietest", usersC.CookieTest).Methods("GET")

	// Gallery routes
	r.Handle("/galleries/new", requireUserMw.Apply(galleriesC.New)).Methods("GET")
	r.Handle("/galleries", requireUserMw.ApplyFn(galleriesC.Index)).Methods("GET")
	r.HandleFunc("/galleries", requireUserMw.ApplyFn(galleriesC.Create)).Methods("POST")
	r.HandleFunc("/galleries/{id:[0-9]+}/edit", requireUserMw.ApplyFn(galleriesC.Edit)).Methods("GET").Name(controllers.EditGallery)
	r.HandleFunc("/galleries/{id:[0-9]+}/update", requireUserMw.ApplyFn(galleriesC.Update)).Methods("POST")
	r.HandleFunc("/galleries/{id:[0-9]+}/images", requireUserMw.ApplyFn(galleriesC.ImageUpload)).Methods("POST")
	r.HandleFunc("/galleries/{id:[0-9]+}/delete", requireUserMw.ApplyFn(galleriesC.Delete)).Methods("POST")
	r.HandleFunc("/galleries/{id:[0-9]+}", galleriesC.Show).Methods("GET").Name(controllers.ShowGallery)

	fmt.Println("Starting the server at port: 3000...")

	// To apply the user middleware to all requests received.
	http.ListenAndServe(":3000", userMw.Apply(r))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
