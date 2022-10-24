package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type MyClaims struct {
	UserName string `json:"user_name"`
	jwt.StandardClaims
}

func main() {
	mySignigkey := []byte("woshigoutou")
	claims := MyClaims{
		UserName: "goutouzi",
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 60,
			ExpiresAt: time.Now().Unix() + 60*60*2,
			Issuer:    "goutouzi",
		},
	}
	jw := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := jw.SignedString(mySignigkey)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(signedString)
	str, err := jwt.ParseWithClaims(signedString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySignigkey, nil
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(str)
}
