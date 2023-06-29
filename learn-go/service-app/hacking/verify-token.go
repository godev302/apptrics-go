package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"os"
)

const token = `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c`

type claims struct {
	jwt.RegisteredClaims
	Roles []string `json:"roles"`
}

func main() {
	PublicPEM, err := os.ReadFile("pubkey.pem")
	if err != nil {
		log.Fatalln("not able to read pem file")
	}
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(PublicPEM)
	if err != nil {
		log.Fatalln(err)
	}
	var c claims // claims struct will hold jwt data after parsing the token

	//jwt.ParseWithClaims verify the token against the public key
	token, err := jwt.ParseWithClaims(token, &c, func(t *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})

	if err != nil {
		fmt.Println("parsing token", err)
	}
	if !token.Valid {
		fmt.Println("invalid token")
		return
	}
	fmt.Println(c)
}
