package auth

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"encoding/json"

	"github.com/orlangure/gnomock"
	"github.com/orlangure/gnomock/preset/postgres"
	"rei.io/rei/internal/helpers"
)

func TestSignup(t *testing.T) {

	// First we spawn a mock postgres in docker
	p := postgres.Preset(
		postgres.WithUser("gnomock", "gnomick"),
		postgres.WithDatabase("mydb"),
	)

	container, _ := gnomock.Start(p)

	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s  dbname=%s sslmode=disable",
		container.Host, container.DefaultPort(),
		"gnomock", "gnomick", "mydb",
	)

	// First case: Should succeed
	body1 := []byte(`{"username":"arthur", "password":"blacklist123", "confirmPassword":"blacklist123", "secret":"thisismyfirstinvitecode"}`)
	req1 := httptest.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(body1))

	// Set the request to use that server
	ctx1 := context.WithValue(req1.Context(), helpers.ConnectionString{}, connStr)
	req1 = req1.WithContext(ctx1)

	req1.Header.Set("Content-Type", "application/json")

	w1 := httptest.NewRecorder()
	Signup(w1, req1)

	res1 := w1.Result()
	defer res1.Body.Close()
	data1, _ := io.ReadAll(res1.Body)

	var x1 interface{}
	json.Unmarshal(data1, &x1)

	got1 := x1

	if reflect.TypeOf(got1.(map[string]interface{})["Status"]) != reflect.TypeOf("") {
		t.Errorf("expected response to contain token but got %v", x1)
	}

	// Second case: Empty fields
	body2 := []byte(`{"username":"", "password":"", "confirmPassword":"", "secret":""}`)
	req2 := httptest.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(body2))

	// Set the request to use that server
	ctx2 := context.WithValue(req2.Context(), helpers.ConnectionString{}, connStr)
	req2 = req2.WithContext(ctx2)

	req2.Header.Set("Content-Type", "application/json")

	w2 := httptest.NewRecorder()
	Signup(w2, req2)

	res2 := w2.Result()
	defer res2.Body.Close()
	data2, _ := io.ReadAll(res2.Body)

	var x2 interface{}
	json.Unmarshal(data2, &x2)

	got2 := x2
	expected2 := map[string]interface{}{
		"Errors": map[string]interface{}{
			"password": "Password must not be empty",
			"secret":   "Sorry, you need the secret to access this website",
			"username": "Username must not be empty",
		},
	}

	if !reflect.DeepEqual(got2, expected2) {
		t.Errorf("expected response to be %v but got %v", expected2, got2)
	}

	// Second case: Short password & Unmatching passwords
	body3 := []byte(`{"username":"ashley", "password":"bjorne", "confirmPassword":"bjorna", "secret":"thisismyfirstinvitecode"}`)
	req3 := httptest.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(body3))

	// Set the request to use that server
	ctx3 := context.WithValue(req3.Context(), helpers.ConnectionString{}, connStr)
	req3 = req3.WithContext(ctx3)

	req3.Header.Set("Content-Type", "application/json")

	w3 := httptest.NewRecorder()
	Signup(w3, req3)

	res3 := w3.Result()
	defer res3.Body.Close()
	data3, _ := io.ReadAll(res3.Body)

	var x3 interface{}
	json.Unmarshal(data3, &x3)

	got3 := x3
	expected3 := map[string]interface{}{
		"Errors": map[string]interface{}{
			"password":        "Password must be longer than 8 digits",
			"confirmPassword": "Passwords are not matching",
		},
	}

	if !reflect.DeepEqual(got3, expected3) {
		t.Errorf("expected response to be %v but got %v", expected3, got3)
	}

	// Fourth case: Username already taken
	body4 := []byte(`{"username":"arthur", "password":"thisisalongpassword", "confirmPassword":"thisisalongpassword", "secret":"thisismyfirstinvitecode"}`)
	req4 := httptest.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(body4))

	// Set the request to use that server
	ctx4 := context.WithValue(req4.Context(), helpers.ConnectionString{}, connStr)
	req4 = req4.WithContext(ctx4)

	req4.Header.Set("Content-Type", "application/json")

	w4 := httptest.NewRecorder()
	Signup(w4, req4)

	res4 := w4.Result()
	defer res4.Body.Close()
	data4, _ := io.ReadAll(res4.Body)

	var x4 interface{}
	json.Unmarshal(data4, &x4)

	got4 := x4
	expected4 := map[string]interface{}{
		"Error": "Username already exists, try a different one",
	}

	if !reflect.DeepEqual(got4, expected4) {
		t.Errorf("expected response to be %v but got %v", expected4, got4)
	}
}

func TestLogin(t *testing.T) {

	// First we spawn a mock postgres in docker
	p := postgres.Preset(
		postgres.WithUser("gnomock", "gnomick"),
		postgres.WithDatabase("mydb"),
	)

	container, _ := gnomock.Start(p)

	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s  dbname=%s sslmode=disable",
		container.Host, container.DefaultPort(),
		"gnomock", "gnomick", "mydb",
	)

	// First case: Empty fields
	body1 := []byte(`{"username":"", "password":""}`)
	req1 := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body1))

	// Set the request to use that server
	ctx1 := context.WithValue(req1.Context(), helpers.ConnectionString{}, connStr)
	req1 = req1.WithContext(ctx1)

	req1.Header.Set("Content-Type", "application/json")

	w1 := httptest.NewRecorder()
	Login(w1, req1)

	res1 := w1.Result()
	defer res1.Body.Close()
	data1, _ := io.ReadAll(res1.Body)

	var x1 interface{}
	json.Unmarshal(data1, &x1)

	got1 := x1
	expected1 := map[string]interface{}{
		"Errors": map[string]interface{}{
			"password": "Password must not be empty",
			"username": "Username must not be empty",
		},
	}

	if !reflect.DeepEqual(got1, expected1) {
		t.Errorf("expected response to be %v but got %v", expected1, got1)
	}

	// Second case: Register first then see if can login

	// Registration phase
	signupBody := []byte(`{"username":"arthur", "password":"blacklist123", "confirmPassword":"blacklist123", "secret":"thisismyfirstinvitecode"}`)
	signupRequest := httptest.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(signupBody))

	// Set the request to use that server
	ctxSignup := context.WithValue(signupRequest.Context(), helpers.ConnectionString{}, connStr)
	signupRequest = signupRequest.WithContext(ctxSignup)

	signupRequest.Header.Set("Content-Type", "application/json")

	wSignup := httptest.NewRecorder()
	Signup(wSignup, signupRequest)

	// Signin phase
	body2 := []byte(`{"username":"arthur", "password":"blacklist123"}`)
	req2 := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body2))

	// Set the request to use that server
	ctx2 := context.WithValue(req2.Context(), helpers.ConnectionString{}, connStr)
	req2 = req2.WithContext(ctx2)

	req2.Header.Set("Content-Type", "application/json")

	w2 := httptest.NewRecorder()
	Login(w2, req2)

	res2 := w2.Result()
	defer res2.Body.Close()
	data2, _ := io.ReadAll(res2.Body)

	var x2 interface{}
	json.Unmarshal(data2, &x2)

	got2 := x2
	if reflect.TypeOf(got2.(map[string]interface{})["Status"]) != reflect.TypeOf("") {
		t.Errorf("expected response to contain token but got %v", x2)
	}

	// Third case: Wrong credentials
	body3 := []byte(`{"username":"arthur", "password":"wrongpassword"}`)
	req3 := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body3))

	// Set the request to use that server
	ctx3 := context.WithValue(req3.Context(), helpers.ConnectionString{}, connStr)
	req3 = req3.WithContext(ctx3)

	req3.Header.Set("Content-Type", "application/json")

	w3 := httptest.NewRecorder()
	Login(w3, req3)

	res3 := w3.Result()
	defer res3.Body.Close()
	data3, _ := io.ReadAll(res3.Body)

	var x3 interface{}
	json.Unmarshal(data3, &x3)

	got3 := x3
	expected3 := map[string]interface{}{
		"Error": "Incorrect credentials, try again",
	}

	if !reflect.DeepEqual(got3, expected3) {
		t.Errorf("expected response to be %v but got %v", expected3, got3)
	}
}
