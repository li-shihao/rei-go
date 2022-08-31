package api

import (
	"net/http"

	"github.com/unrolled/render"
	"rei.io/rei/internal/database"
)

// For homepage display
func TotalTransactionCount(w http.ResponseWriter, r *http.Request) {

	// Set response type header for json
	w.Header().Set("Content-Type", "application/json")

	// Unwrap context to obtain correct db connection string
	ctx := r.Context()

	// Assert type string
	connStr := ctx.Value(key{}).(string)

	// Initialise db
	db := new(database.EntClient)
	db.Init("postgres", connStr)

	count, err := db.QueryTotalTransactionCount()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	// Rendering json repsonse
	render.New().JSON(w, 200, map[string]int{"Total transaction count": *count})
}
