package main

import (
	"fmt"
	"os"
	"log"
	"time"
	"path/filepath"
	"github.com/golang-jwt/jwt/v5"
)

var sampleSecretKey = []byte("SecretYouShouldHide")
//var sampleSecretKey = "SecretYouShouldHide"

func setLogFile() {
	dir, _ := filepath.Abs(".")
	fmt.Println("app.log path -> " + dir)

	f, err := os.OpenFile(dir+"/app.log", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatalf("error opening log file: %v", err)
	}
	//defer f.Close()
	log.SetOutput(f)
}

func generateJWT(mail string) (string, error) {
	// JWT Token olusturmaca
	var token = jwt.New(jwt.SigningMethodHS256)

	var claims = token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(10 * time.Minute)
	claims["auth"] = true
	claims["email"] = mail

	tokenStr, err := token.SignedString(sampleSecretKey)
	return tokenStr, err
}