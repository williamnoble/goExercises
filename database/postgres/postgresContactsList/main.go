package main

import (
	"encoding/json"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/types"
	"github.com/pkg/errors"
)

// postgresql://postgres:@DatabaseUser246@localhost:32768

// ContactFavorites defines a contact user favorites objects
type ContactFavorites struct {
	Colors []string `json:"colors"`

	// Contact model structure, note FavoritesJSON is a JSON Field which stores the JSONified favorites,
	// wheras the Favorities field is what we use to store the unMarshalled Data
}
type Contact struct {
	Name, Address, Phone string
	ID                   int
	FavoritesJSON        types.JSONText    `db:"favorites"`
	Favorites            *ContactFavorites `db:"-"`

	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}

var (
	connectionString = flag.String("conn", getenvWithDefault("DATABASE_URL", ""), "PostgreSQL connection string")
	listenAddr       = flag.String("addr", ":8080", "HTTP address to listen on")
	db               *sqlx.DB
	tmpl             = template.New("")
)

func getenvWithDefault(name, defaultValue string) string {
	val := os.Getenv(name)
	if val == "" {
		val = defaultValue
	}
	return val
}

func handler(w http.ResponseWriter, r *http.Request) {
	contacts, err := fetchContacts()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	tmpl.ExecuteTemplate(w, "index.html", struct{ Contacts []*Contact }{contacts})
}
func fetchContacts() ([]*Contact, error) {
	contacts := []*Contact{}
	err := db.Select(&contacts, "select * from contacts")
	if err != nil {
		return nil, errors.Wrap(err, "unable to fetch contacts")
	}

	for _, contact := range contacts {
		err := json.Unmarshal(contact.FavoritesJSON, &contact.Favorites)

		if err != nil {
			return nil, errors.Wrap(err, "Unable to parse JSON favorites")
		}
	}
	return contacts, nil

}

func main() {

	tmpl.Funcs(template.FuncMap{"StringsJoin": strings.Join})
	flag.Parse()
	var err error
	//_, err = tmpl.ParseGlob(filepath.Join(".", "templates", "*.html"))
	_, err = tmpl.ParseGlob(filepath.Join("./templates/index.html"))
	if err != nil {
		log.Fatalf("unable to parse template files: %v\n", err)
	}

	http.HandleFunc("/", handler)

	if *connectionString == "" {
		log.Fatalln("Please pass connecting string using the -conn option")
	}
	db, err = sqlx.Connect("pgx", *connectionString)
	if err != nil {
		log.Fatalf("Unable to establish a connection: %v\n", err)

	}
	log.Printf("Connected successfully to DB\n")
	log.Printf("listening on %s\n", *listenAddr)
	http.ListenAndServe(*listenAddr, nil)

}
