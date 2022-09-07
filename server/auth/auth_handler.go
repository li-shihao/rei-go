package auth

import (
	"io"
	"log"
	"net"
	"net/http"

	"encoding/json"

	"github.com/unrolled/render"
	"rei.io/rei/internal/crypto"
	"rei.io/rei/internal/database"
	"rei.io/rei/internal/helpers"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	// Set response type header for json
	w.Header().Set("Content-Type", "application/json")

	/*
		Section 1: Basic parsing & Checking
	*/

	// Response object
	errors := map[string]string{}

	// Signup form format
	type request struct {
		Username        string `json:"username"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirmPassword"`
		Secret          string `json:"secret"`
	}

	// Placeholder for unmarshalling json
	var incomeRequest request

	// Unmarshals json request
	x, err := io.ReadAll(r.Body)
	if err != nil {
		render.New().JSON(w, 500, map[string]string{"Error": "Something went wrong"})
	}
	json.Unmarshal(x, &incomeRequest)

	// Check if any field are bad
	if incomeRequest.Username == "" {
		errors["username"] = "Username must not be empty"
	}
	if len(incomeRequest.Password) <= 8 {
		errors["password"] = "Password must be longer than 8 digits"
	}
	if incomeRequest.Password == "" {
		errors["password"] = "Password must not be empty"
	}
	if incomeRequest.ConfirmPassword != incomeRequest.Password {
		errors["confirmPassword"] = "Passwords are not matching"
	}
	if incomeRequest.Secret != "thisismyfirstinvitecode" {
		errors["secret"] = "Sorry, you need the secret to access this website"
	}

	if len(errors) != 0 {
		render.New().JSON(w, 400, map[string]interface{}{"Errors": errors})
		return
	}

	/*
		Section 2: Check if user already exists
	*/

	// Unwrap context to obtain correct db connection string
	ctx := r.Context()

	// Assert type string
	connStr := ctx.Value(helpers.ConnectionString{}).(string)

	// Initialise db
	db := new(database.EntClient)
	db.Init("postgres", connStr)

	exist, err := db.QueryUserExist(incomeRequest.Username)
	if err != nil {
		render.New().JSON(w, 500, map[string]string{"Error": "Something went wrong"})
		return
	}
	if *exist {
		render.New().JSON(w, 400, map[string]string{"Error": "Username already exists, try a different one"})
		return
	}

	/*
		Section 3: Create new user
	*/

	// Generate jwt using username
	_, err = db.CreateUser(incomeRequest.Username, incomeRequest.Password)

	if err != nil {
		render.New().JSON(w, 500, map[string]string{"Error": "Something went wrong"})
		return
	}

	tokenString, err := crypto.GenerateJWT(incomeRequest.Username)

	if err != nil {
		render.New().JSON(w, 500, map[string]string{"Error": "Something went wrong"})
		return
	}

	// Set cookie on user
	http.SetCookie(w, &http.Cookie{
		Name:  "jwt",
		Path:  "/",
		Value: tokenString,
		//SameSite: http.SameSiteStrictMode,
		//HttpOnly: true,
		// Secure: true,
	})

	host, _, _ := net.SplitHostPort(r.RemoteAddr)

	// Create session in db
	db.CreateSession(incomeRequest.Username, host)

	// Rendering json repsonse
	render.New().JSON(w, 201, map[string]string{"Status": "Signup succeeded"})
}

func Login(w http.ResponseWriter, r *http.Request) {

	// Set response type header for json
	w.Header().Set("Content-Type", "application/json")

	/*
		Section 1: Basic parsing & Checking
	*/

	// Response object
	errors := map[string]string{}

	// Signup form format
	type request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// Placeholder for unmarshalling json
	var incomeRequest request

	// Unmarshals json request
	x, err := io.ReadAll(r.Body)
	if err != nil {
		render.New().JSON(w, 500, map[string]string{"Error": "Something went wrong"})
	}
	json.Unmarshal(x, &incomeRequest)

	// Check if any field are bad
	if incomeRequest.Username == "" {
		errors["username"] = "Username must not be empty"
	}
	if incomeRequest.Password == "" {
		errors["password"] = "Password must not be empty"
	}

	if len(errors) != 0 {
		render.New().JSON(w, 400, map[string]interface{}{"Errors": errors})
		return
	}

	/*
		Section 2: Check for authentication
	*/

	// Unwrap context to obtain correct db connection string
	ctx := r.Context()

	// Assert type string
	connStr := ctx.Value(helpers.ConnectionString{}).(string)

	// Initialise db
	db := new(database.EntClient)
	db.Init("postgres", connStr)

	correct, err := db.QueryUserCredentials(incomeRequest.Username, incomeRequest.Password)
	if correct == nil || !*correct {
		log.Println(incomeRequest.Username)
		log.Println(incomeRequest.Password)
		log.Println(correct)
		render.New().JSON(w, 400, map[string]string{"Error": "Incorrect credentials, try again"})
		return
	}
	if err != nil {
		render.New().JSON(w, 500, map[string]string{"Error": "Something went wrong"})
		return
	}

	// Find out if user has already logged in somewhere else
	if loggedIn, _, err := db.QuerySession(incomeRequest.Username); err != nil || loggedIn == nil {

		// If db query goes wrong
		render.New().JSON(w, 500, map[string]string{"Error": "Something went wrong"})
		return

		// If already logged in, delete previous session
	} else if *loggedIn {
		db.DeleteSession(incomeRequest.Username)
	}

	tokenString, err := crypto.GenerateJWT(incomeRequest.Username)

	if err != nil {
		render.New().JSON(w, 500, map[string]string{"Error": "Something went wrong"})
		return
	}

	// Set cookie on user
	http.SetCookie(w, &http.Cookie{
		Name:  "jwt",
		Path:  "/",
		Value: tokenString,
		//SameSite: http.SameSiteStrictMode,
		//HttpOnly: true,
		//Secure: true,
	})

	host, _, _ := net.SplitHostPort(r.RemoteAddr)

	// Create session in db
	db.CreateSession(incomeRequest.Username, host)

	// Rendering json repsonse
	render.New().JSON(w, 201, map[string]string{"Status": "Login succeeded"})
}

func Logout(w http.ResponseWriter, r *http.Request) {

	// Unwrap context to obtain correct db connection string
	ctx := r.Context()

	// Assert type string
	connStr := ctx.Value(helpers.ConnectionString{}).(string)
	user := ctx.Value(helpers.UsernameClaim{}).(string)

	// Initialise db
	db := new(database.EntClient)
	db.Init("postgres", connStr)

	if loggedIn, _, err := db.QuerySession(user); err != nil || loggedIn == nil || !*loggedIn {

		// If db query goes wrong
		render.New().JSON(w, 500, map[string]string{"Error": "Something went wrong"})
		return

		// If already logged in, delete previous session
	} else if *loggedIn {
		db.DeleteSession(user)
	}

	// Clear cookie on user
	http.SetCookie(w, &http.Cookie{
		Name:  "jwt",
		Path:  "/",
		Value: "",
		//SameSite: http.SameSiteStrictMode,
		//HttpOnly: true,
		//Secure: true,
	})

	// Rendering json repsonse
	render.New().JSON(w, 201, map[string]string{"Status": "Logout successful"})
}

func Confirm(w http.ResponseWriter, r *http.Request) {
	render.New().JSON(w, 201, map[string]string{"Status": "Authenticated"})
}
