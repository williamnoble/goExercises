package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/williamnoble/goExercises/http/sql-gorm-api/api/app/config"
	"github.com/williamnoble/goExercises/http/sql-gorm-api/api/app/handler"
	"github.com/williamnoble/goExercises/http/sql-gorm-api/api/app/model"
)

type Application struct {
	Router *mux.Router
	DB     *gorm.DB
}

func (a *Application) Initialize(config *config.Config) {
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True",
		config.DB.Username,
		config.DB.Password,
		config.DB.Host,
		config.DB.Port,
		config.DB.Name,
		config.DB.Charset)

	fmt.Println(dbURI, "+", config.DB.Dialect)

	db, err := gorm.Open("mysql", dbURI)
	if err != nil {
		log.Fatal("Could not connect to the db")
	}

	a.DB = model.DBMigrate(db)
	a.Router = mux.NewRouter()
	a.setRouters()
}
func (a *Application) setRouters() {
	// Routing for handling the projects
	a.Get("/employees", a.GetAllEmployees)
	a.Post("/employees", a.CreateEmployee)
	a.Get("/employees/{title}", a.GetEmployee)
	a.Put("/employees/{title}", a.UpdateEmployee)
	a.Delete("/employees/{title}", a.DeleteEmployee)
	a.Put("/employees/{title}/disable", a.DisableEmployee)
	a.Put("/employees/{title}/enable", a.EnableEmployee)
}

func (a *Application) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

func (a *Application) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

func (a *Application) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

func (a *Application) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

func (a *Application) GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	handler.GetAllEmployees(a.DB, w, r)
}

func (a *Application) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	handler.CreateEmployee(a.DB, w, r)
}

func (a *Application) GetEmployee(w http.ResponseWriter, r *http.Request) {
	handler.GetEmployee(a.DB, w, r)
}

func (a *Application) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	handler.UpdateEmployee(a.DB, w, r)
}

func (a *Application) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	handler.DeleteEmployee(a.DB, w, r)
}

func (a *Application) DisableEmployee(w http.ResponseWriter, r *http.Request) {
	handler.DisableEmployee(a.DB, w, r)
}

func (a *Application) EnableEmployee(w http.ResponseWriter, r *http.Request) {
	handler.EnableEmployee(a.DB, w, r)
}

// Run the app on it's router
func (a *Application) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}
