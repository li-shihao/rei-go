package api

import (
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/unrolled/render"
	"rei.io/rei/internal/database"
	"rei.io/rei/internal/helpers"
)

// For homepage display
func TotalTransactionCount(w http.ResponseWriter, r *http.Request) {

	// Set response type header for json
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-CSRF-Token", csrf.Token(r))

	// Unwrap context to obtain correct db connection string
	ctx := r.Context()

	// Assert type string
	connStr := ctx.Value(helpers.ConnectionString{}).(string)

	// Initialise db
	db := new(database.EntClient)
	db.Init("postgres", connStr)

	count, err := db.QueryTotalTransactionCount()
	if err != nil {
		render.New().JSON(w, 500, map[string]string{"Error": "Something went wrong."})
		return
	}

	// Rendering json repsonse
	render.New().JSON(w, 200, map[string]int{"Total transaction count": *count})
}
