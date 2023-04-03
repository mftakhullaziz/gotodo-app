package helpers

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) string {
	// Generate the bcrypt hash from the password string
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	hashedPassword := string(hashedPasswordBytes)
	return hashedPassword
}

func HashPasswordAndSalt(pwd []byte) string {
	log := LoggerParent()
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash password")
	}
	return string(hash)
}

func ComparedPassword(hashedPwd string, plainPassword []byte) (bool, error) {
	log := LoggerParent()
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		log.Println(err)
		return false, err
	}
	return true, nil
}
