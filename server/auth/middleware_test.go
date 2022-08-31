package auth

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/goccy/go-json"
	"github.com/golang-jwt/jwt"
	"github.com/unrolled/render"
	"rei.io/rei/internal/crypto"
)

func TestJWTAuthenticate(t *testing.T) {

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		render.New().JSON(w, 201, map[string]string{"status": "success"})
	})

	wrapper := JWTAuthenticate(handler)

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
	tokenString2, _, _ := crypto.GenerateJWT("testpassword123")

	w2 := httptest.NewRecorder()

	req2 := httptest.NewRequest(http.MethodPost, "/", nil)
	req2.AddCookie(&http.Cookie{
		Name:     "jwt",
		Value:    tokenString2,
		SameSite: http.SameSiteStrictMode,
		HttpOnly: true,
		Secure:   true,
	})
	req2.Header.Set("Content-Type", "application/json")

	wrapper.ServeHTTP(w2, req2)

	res2 := w2.Result()
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
}
