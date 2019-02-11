package helpers

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func GetPwd(pwd string) []byte {
	_, err := fmt.Scan(pwd)
	if err != nil {
		log.Println(err)
	}

	return []byte(pwd)
}

func HashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func ComparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
