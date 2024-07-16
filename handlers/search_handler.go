package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type File struct {
	ID       int    `json:"id"`
	Filename string `json:"filename"`
	Filepath string `json:"filepath"`
}

func SearchHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		query := r.URL.Query().Get("query")
		if query == "" {
			http.Error(w, "Query parameter is required", http.StatusBadRequest)
			return
		}

		rows, err := db.Query("SELECT id, filename, filepath FROM files WHERE filename LIKE $1", "%"+query+"%")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var files []File
		for rows.Next() {
			var file File
			err := rows.Scan(&file.ID, &file.Filename, &file.Filepath)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			files = append(files, file)
		}
		if err = rows.Err(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		totalSearchTime := float64(time.Since(startTime).Milliseconds())
		fmt.Printf("Time taken for search %s is %f \n", query, totalSearchTime)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(files)
	}
}
