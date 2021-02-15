package main

// FOR RUNNING
// go mod vendor
//go build
// .\Login_Auth.exe

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//CREATE JWT SECRET KEY THAT ARE USED TO SIGN OUR JWT TOKEN
var jwtKey = []byte("secret_key")

//HERE WE ALSO USED DATABASE
var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

//VIA THIS WE CAN PASS DATA THROUGH API
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//CLAIMS ARE USING TO CREATE  PAYLOAD OF OUR JWT
//INSIDE THE PAYLOAD WE HAVE PASSING THE USER NAME AND WHEN THE PARTICULAR TOKEN IS EXPIRING
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func Login(w http.ResponseWriter, r *http.Request) {
	var credentials Credentials
	// ACCESS DATA FROM REQUEST BODY
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// CHECK EXPECTED USER PRESENT IN OUR MAP( OR DATABASE )
	expectedPassword, ok := users[credentials.Username]

	if !ok || expectedPassword != credentials.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// THIS IS EXPIRATION TIME OF OUR TOKEN
	expirationTime := time.Now().Add(time.Minute * 5)

	// CREATE CLAIMS FOR OUR TOKEN
	claims := &Claims{
		Username: credentials.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// CREATE TOKEN
	// HERE WE HS 256 ALGORITHM FOR SIGNING IN
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// GET THE TOKEN STRING FROM THE TOKEN
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// SET ABOVE THING INTO OUR COOKIES
	http.SetCookie(w,
		&http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expirationTime,
		})

}

func Home(w http.ResponseWriter, r *http.Request) {
	// GET THE COOKIE FROM REQUEST
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// GET THE VALUE OF THE TOKEN
	tokenStr := cookie.Value

	// CREATE THE  REFERENCE OF CLAIMS
	claims := &Claims{}

	// PARSE WITH CLAIMS AND GET THE TOKEN
	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	//  IF ALL THINGS ARE CORRECT THEN PRINT THEM
	w.Write([]byte(fmt.Sprintf("Hello, %s", claims.Username)))

}

func Refresh(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokenStr := cookie.Value

	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// CHECK THE TIME INTERVAL IF THE TIME IS LESS THAN 30 SECOND THEN YOU CAN REFRESH THE TOKEN OTHERWISE NOT
	// if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	expirationTime := time.Now().Add(time.Minute * 5)

	claims.ExpiresAt = expirationTime.Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w,
		&http.Cookie{
			Name:    "refresh_token",
			Value:   tokenString,
			Expires: expirationTime,
		})

}
