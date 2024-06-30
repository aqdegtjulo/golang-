package pwd

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

// 加密密码
func HashPwd(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

// 验证密码，hash之后的密码
func CheckPwd(hashPwd string, Pwd string) bool {
	byteHash := []byte(hashPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, []byte(Pwd))
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
