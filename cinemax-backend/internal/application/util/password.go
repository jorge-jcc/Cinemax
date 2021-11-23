package util

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func AuthenticatePassword(storedPassword, suppliedPassword string) bool {
	passwordCheck := []byte(suppliedPassword)
	passwordEmpleado := []byte(storedPassword)
	err := bcrypt.CompareHashAndPassword(passwordEmpleado, passwordCheck)
	return err == nil
}
