package util

import "golang.org/x/crypto/bcrypt"

// HashEncode 加密
func HashEncode(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// ComparePwd 解密
func ComparePwd(hashedPwd string, sourcePwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(sourcePwd))

	if err != nil {
		return false
	}
	return true
	// return err == nil
}
