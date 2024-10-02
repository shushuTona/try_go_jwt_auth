package main

import (
	"fmt"

	"github.com/golang-jwt/jwt"
)

func GenerateToken() {
	token := jwt.New(jwt.SigningMethodHS256)

	fmt.Println("===== GenerateToken =====")
	fmt.Printf("token: %#v\n", token)
	fmt.Printf("token.Raw: %#v\n", token.Raw)
	fmt.Printf("token.Method: %#v\n", token.Method)
	fmt.Printf("token.Header: %#v\n", token.Header)
	fmt.Printf("token.Claims: %#v\n", token.Claims)
	fmt.Printf("token.Claims.Valid(): %#v\n", token.Claims.Valid())
	fmt.Printf("token.Signature: %#v\n", token.Signature)
	fmt.Printf("token.Valid: %#v\n", token.Valid)

	testToken, err := token.SignedString([]byte("HOGE"))
	if err != nil {
		panic(err)
	}

	fmt.Println(testToken)
	// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.e30.fWg5MAbC90YGsrdLlQdBgKaY3jt6vYTKqzVSj1EQ6Jk
}

func GenerateTokenWithStandardClaims() {
	claims := jwt.StandardClaims{}
	claims.Subject = "{\"id\": 100}"
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	fmt.Println("===== GenerateTokenWithStandardClaims =====")
	fmt.Printf("token.Raw: %#v\n", token.Raw)
	fmt.Printf("token.Method: %#v\n", token.Method)
	fmt.Printf("token.Header: %#v\n", token.Header)
	fmt.Printf("token.Claims: %#v\n", token.Claims)
	fmt.Printf("token.Claims.Valid(): %#v\n", token.Claims.Valid())
	fmt.Printf("token.Signature: %#v\n", token.Signature)
	fmt.Printf("token.Valid: %#v\n", token.Valid)

	testToken, err := token.SignedString([]byte("HOGE"))
	if err != nil {
		panic(err)
	}

	fmt.Println(testToken)
	// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJ7XCJpZFwiOiAxMDB9In0.WUNDMVe9-BekeaQisp4lb8o6jNhg9UQiJECSFlJqonU
}

func GenerateTokenWithMapClaims() {
	claims := jwt.MapClaims{}
	claims["sub"] = "{\"id\": 100}"
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	fmt.Println("===== GenerateTokenWithMapClaims =====")
	fmt.Printf("token: %#v\n", token)
	fmt.Printf("token.Raw: %#v\n", token.Raw)
	fmt.Printf("token.Method: %#v\n", token.Method)
	fmt.Printf("token.Header: %#v\n", token.Header)
	fmt.Printf("token.Claims: %#v\n", token.Claims)
	fmt.Printf("token.Claims.Valid(): %#v\n", token.Claims.Valid())
	fmt.Printf("token.Signature: %#v\n", token.Signature)
	fmt.Printf("token.Valid: %#v\n", token.Valid)

	testToken, err := token.SignedString([]byte("HOGE"))
	if err != nil {
		panic(err)
	}

	fmt.Println(testToken)
	// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJ7XCJpZFwiOiAxMDB9In0.WUNDMVe9-BekeaQisp4lb8o6jNhg9UQiJECSFlJqonU
}
