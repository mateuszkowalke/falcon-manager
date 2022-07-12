package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"time"

	_ "github.com/lib/pq"
)

type falcon struct {
	id                int
	name              string
	ringNo            string
	speciesId         int
	sex               string
	birthDate         time.Time
	source            string
	aviaryId          int
	fatherId          int
	motherId          int
	widthYoung        int
	lengthYoung       int
	weightYoung       int
	widthOld          int
	lengthOld         int
	weightOld         int
	notes             string
	createdAt         time.Time
	updatedAt         time.Time
	breedingProjectId int
}

type Species struct {
	Id   int
	Name string
}

func main() {
	db := connect()
	defer db.Close()

	mux := http.NewServeMux()
	// static files serving
	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("/create-falcon", func(w http.ResponseWriter, r *http.Request) {
		species, err := getSpecies(db)
		if err != nil {
			log.Printf("error retrieving the list of species: %s\n", err)
		}
		serveTemplate("falcon-form.html", w, species)
	})
	mux.HandleFunc("/falcon", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			falconData := getFalconDataFromRequest(r)
			err := createFalcon(db, &falconData)
			if err != nil {
				log.Printf("Could not create a falcon %v+", falconData)
			}
			serveTemplate("falcon.html", w, nil)
			return
		}
		serveTemplate("falcon.html", w, nil)
	})
	mux.HandleFunc("/falcons", func(w http.ResponseWriter, r *http.Request) {
		serveTemplate("falcons.html", w, nil)
	})
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		serveTemplate("index.html", w, nil)
	})

	http.ListenAndServe(":3000", mux)
}

func serveTemplate(name string, w http.ResponseWriter, data any) {
	lp := filepath.Join("pages", "layout.html")
	fp := filepath.Join("pages", name)

	tmpl, err := template.ParseFiles(lp, fp)
	if err != nil {
		log.Printf("Something went wrong parsing template files for \"%s\":\n%s\n", name, err)
	}

	err = tmpl.ExecuteTemplate(w, "layout", data)
	if err != nil {
		log.Printf("Something went wrong executing template files for \"%s\":\n%s\n", name, err)
	}
}

func getFalconDataFromRequest(r *http.Request) falcon {
	return falcon{
		ringNo: r.FormValue("ring_no"),
		name:   r.FormValue("name"),
	}
}

func createFalcon(db *sql.DB, falcon *falcon) error {
	db.Exec("INSERT ()")
	return nil
}

func getSpecies(db *sql.DB) ([]Species, error) {
	rows, err := db.Query("SELECT id, name FROM species;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var species []Species
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			return nil, err
		}
		species = append(species, Species{id, name})
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return species, nil
}

func connect() *sql.DB {
	db, err := sql.Open("postgres", "postgresql://testuser:testpass@pgdb:5432/falcon_manager?sslmode=disable")
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	return db
}
