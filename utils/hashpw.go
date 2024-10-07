package utils

import "golang.org/x/crypto/bcrypt"

func GenerateHashPassword(pwd string) string {
	password, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		return string(password)
	}

	return string(password)
}

func CompareHashPassword(hashpwd, pwd string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashpwd), []byte(pwd))
}
