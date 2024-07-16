package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/bxcodec/faker/v3"
	_ "github.com/lib/pq"
	"pg.search/handlers"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "testdb"
)

func main() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to database: %q", err)
	}
	fmt.Println("Successfully connected to the database")

	// createTable(db)
	// totalFiles := 100000
	// insertValues(db, totalFiles)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from postgres search")
	})

	http.HandleFunc("/search", handlers.SearchHandler(db))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func createTable(db *sql.DB) {
	createTableSQL := `CREATE TABLE IF NOT EXISTS files (
        id SERIAL PRIMARY KEY,
        filename TEXT NOT NULL,
        filepath TEXT NOT NULL
    );`
	_, err := db.Exec(createTableSQL)
	if err != nil {
		log.Fatalf("Error creating table: %q", err)
	}
}

func insertValues(db *sql.DB, totalFiles int) {
	fileNames := []string{}
	for i := 0; i < totalFiles; i++ {
		filename := faker.Word() + ".txt"
		fileNames = append(fileNames, filename)
		filepath := "/" + faker.Word() + "/" + faker.Word() + "/" + filename
		_, err := db.Exec("INSERT INTO files (filename, filepath) VALUES ($1, $2)", filename, filepath)
		if err != nil {
			log.Fatalf("Error inserting data: %q", err)
		}
	}

	jsonData, err := json.MarshalIndent(fileNames, "", "  ")
	if err != nil {
		fmt.Println("error marshaling")
	}
	err = os.WriteFile("scripts/file_names.json", jsonData, 0644)
	if err != nil {
		fmt.Println("error writing in json file")
	}
	fmt.Println("Random values inserted successfully")
}
