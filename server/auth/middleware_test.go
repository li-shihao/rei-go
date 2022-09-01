package auth

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/goccy/go-json"
	"github.com/golang-jwt/jwt"
	"github.com/orlangure/gnomock"
	"github.com/orlangure/gnomock/preset/postgres"
	"github.com/unrolled/render"
	"rei.io/rei/internal/helpers"
)

func TestAuthenticate(t *testing.T) {

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

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		render.New().JSON(w, 201, map[string]string{"status": "success"})
	})

	wrapper := Authenticate(handler)

	// First case: No cookie set
	req1 := httptest.NewRequest(http.MethodGet, "/", nil)
	req1.Header.Set("Content-Type", "application/json")

	w1 := httptest.NewRecorder()
	wrapper.ServeHTTP(w1, req1)

	res1 := w1.Result()
	defer res1.Body.Close()
	data1, _ := io.ReadAll(res1.Body)

	var x1 interface{}
	json.Unmarshal(data1, &x1)

	if _, ok := x1.(map[string]interface{})["Error"]; !ok {
		t.Errorf("expected response to be %v but got %v", false, x1)
	}

	// Second case: Cookie set
	body2 := []byte(`{"username":"arthur", "password":"blacklist123", "confirmPassword":"blacklist123", "secret":"thisismyfirstinvitecode"}`)
	req2 := httptest.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(body2))

	// Set the request to use that server
	req2.Header.Set("Content-Type", "application/json")

	w2a := httptest.NewRecorder()
	ctx2 := context.WithValue(req2.Context(), helpers.ConnectionString{}, connStr)
	req2 = req2.WithContext(ctx2)
	Signup(w2a, req2)

	cookie := w2a.Header()["Set-Cookie"][0]
	req2 = httptest.NewRequest(http.MethodPost, "/", nil)
	req2.Header.Add("Cookie", cookie)
	req2 = req2.WithContext(ctx2)
	w2b := httptest.NewRecorder()
	wrapper.ServeHTTP(w2b, req2)

	res2 := w2b.Result()

	defer res2.Body.Close()
	data2, _ := io.ReadAll(res2.Body)
	var x2 interface{}
	json.Unmarshal(data2, &x2)

	if v, ok := x2.(map[string]interface{})["status"]; !ok {
		t.Errorf("expected response to be %v but got %v", true, x2)
	} else if v.(string) != "success" {
		t.Errorf("expected response to be %v but got %v", true, x2)
	}

	// Third case: Bad token
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{"username": "123"})
	tokenString3, _ := token.SignedString("fakekey")

	w3 := httptest.NewRecorder()

	req3 := httptest.NewRequest(http.MethodPost, "/", nil)
	req3.AddCookie(&http.Cookie{
		Name:     "jwt",
		Value:    tokenString3,
		SameSite: http.SameSiteStrictMode,
		HttpOnly: true,
		Secure:   true,
	})
	req3.Header.Set("Content-Type", "application/json")

	wrapper.ServeHTTP(w3, req3)

	res3 := w3.Result()
	defer res3.Body.Close()
	data3, _ := io.ReadAll(res3.Body)

	var x3 interface{}
	json.Unmarshal(data3, &x3)

	if x3.(map[string]interface{})["Error"] != "bad token" {
		t.Errorf("expected response to be %v but got %v", false, x3)
	}

	//Fourth case: Logged in from somewhere else (Already signed up from case 2)
	body4 := []byte(`{"username":"arthur", "password":"blacklist123"}`)
	req4 := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body4))

	ctx4 := context.WithValue(req4.Context(), helpers.ConnectionString{}, connStr)
	req4 = req4.WithContext(ctx4)

	req4.Header.Set("Content-Type", "application/json")
	req4.RemoteAddr = "192.168.0.1"

	w4 := httptest.NewRecorder()
	Login(w4, req4)

	// Now test middlware and see if token gets rejected and tell us we have been logged off
	wc := httptest.NewRecorder()
	wrapper.ServeHTTP(wc, req2)

	res2 = wc.Result()
	defer res2.Body.Close()
	data2, _ = io.ReadAll(res2.Body)

	json.Unmarshal(data2, &x2)

	if x2.(map[string]interface{})["Error"] != "logged off from another client. Please login again" {
		t.Errorf("expected response to be %v but got %v", false, x2)
	}
}

func TestAdminOnly(t *testing.T) {

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

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		render.New().JSON(w, 201, map[string]string{"status": "success"})
	})

	wrapper := Authenticate(AdminOnly(handler))

	// First case: is Admin
	body1 := []byte(`{"username":"arthur", "password":"blacklist123", "confirmPassword":"blacklist123", "secret":"thisismyfirstinvitecode"}`)
	req1 := httptest.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(body1))

	// Set the request to use that server
	req1.Header.Set("Content-Type", "application/json")

	w1a := httptest.NewRecorder()
	ctx1 := context.WithValue(req1.Context(), helpers.ConnectionString{}, connStr)
	req1 = req1.WithContext(ctx1)
	Signup(w1a, req1)

	cookie1 := w1a.Header()["Set-Cookie"][0]
	req1 = httptest.NewRequest(http.MethodPost, "/", nil)
	req1.Header.Add("Cookie", cookie1)
	req1 = req1.WithContext(ctx1)
	w1b := httptest.NewRecorder()
	wrapper.ServeHTTP(w1b, req1)

	res1 := w1b.Result()

	defer res1.Body.Close()
	data1, _ := io.ReadAll(res1.Body)
	var x1 interface{}
	json.Unmarshal(data1, &x1)

	if _, ok := x1.(map[string]interface{})["status"]; !ok {
		t.Errorf("expected response to be %v but got %v", true, x1)
	}

	// Second case: is not Admin
	body2 := []byte(`{"username":"arthurr", "password":"blacklist123", "confirmPassword":"blacklist123", "secret":"thisismyfirstinvitecode"}`)
	req2 := httptest.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(body2))

	// Set the request to use that server
	req2.Header.Set("Content-Type", "application/json")

	w2a := httptest.NewRecorder()
	ctx2 := context.WithValue(req2.Context(), helpers.ConnectionString{}, connStr)
	req2 = req2.WithContext(ctx2)
	Signup(w2a, req2)

	cookie2 := w2a.Header()["Set-Cookie"][0]
	req2 = httptest.NewRequest(http.MethodPost, "/", nil)
	req2.Header.Add("Cookie", cookie2)
	req2 = req2.WithContext(ctx1)
	w2b := httptest.NewRecorder()
	wrapper.ServeHTTP(w2b, req2)

	res2 := w2b.Result()

	defer res2.Body.Close()
	data2, _ := io.ReadAll(res2.Body)
	var x2 interface{}
	json.Unmarshal(data2, &x2)

	if _, ok := x2.(map[string]interface{})["Error"]; !ok {
		t.Errorf("expected response to be %v but got %v", true, x2)
	}

}
