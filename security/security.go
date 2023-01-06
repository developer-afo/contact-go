package security

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/valyala/fasthttp"
)

const (
	// JWT_KEY : const
	JWT_KEY = "3085e4c8a97e3fe2bc4d6a5b95d525b3"
)

// CreateToken : function
func CreateToken(userID string) (string, error) {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	// This is the information which frontend can use
	// The backend can also decode the token and get admin etc.
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = userID
	claims["iat"] = time.Now().Unix()
	claims["ver"] = 1 //version
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	// Generate encoded token and send it as response.
	// The signing string should be secret (a generated UUID works too)
	_token, err := token.SignedString([]byte(JWT_KEY))
	if err != nil {
		return "", err
	}

	return _token, nil
}

// ExtractUserID : function
func ExtractUserID(r *fasthttp.Request) (uid uuid.UUID, err error) {

	token, err := ExtractTokenObject(r)

	claims := token.Claims.(jwt.MapClaims)

	if claims["sub"] == nil {
		return uuid.Nil, errors.New("invalid token: user id not found")
	}
	userID := claims["sub"].(string)
	return uuid.Parse(userID)
}

// TokenValid : function
func TokenValid(r *fasthttp.Request) error {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(JWT_KEY), nil
	})
	if err != nil {
		return err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		Pretty(claims)
	}
	return nil
}

// ExtractToken : function
func ExtractToken(r *fasthttp.Request) string {

	keys := r.URI().QueryArgs()
	token := string(keys.Peek("token"))
	if token != "" {
		return token
	}

	bearerToken := string(r.Header.Peek("Authorization"))
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

// ExtractTokenObject : function
func ExtractTokenObject(r *fasthttp.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(JWT_KEY), nil
	})
	if err != nil {
		return nil, err
	}
	return token, err
}

// Pretty display the claims licely in the terminal
func Pretty(data interface{}) {
	b, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(string(b))
}

func GetSecret() string {
	return JWT_KEY
}
