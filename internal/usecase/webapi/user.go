package webapi

import "golang.org/x/crypto/bcrypt"

func (g *GopherMartRepoWebAPI) GetPasswordHash(
	password string) (hashedPassword string, err error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	hashedPassword = string(hash)
	return
}

func (g *GopherMartRepoWebAPI) VerifyPassword(
	password, hashedPassword string) (err error) {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
