package main

import (
	"fmt"
)

func main() {
	// GenerateToken()
	// GenerateTokenWithStandardClaims()
	// GenerateTokenWithMapClaims()

	tokenStr, err := GenerateTokenByPrivateKey()
	if err != nil {
		panic(err)
	}
	fmt.Printf("tokenStr: %#v\n", tokenStr)

	token, err := ParseTokenByPublicKey(tokenStr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("token: %#v\n", token)
}
