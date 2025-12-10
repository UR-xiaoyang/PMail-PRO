package password

import (
	"crypto/md5"
	"encoding/hex"

	"golang.org/x/crypto/bcrypt"
)

// Encode 对密码两次md5加盐
func Encode(password string) string {
	encodePwd := Md5Encode(Md5Encode(password+"pmail") + "pmail2023")
	return encodePwd
}

func Md5Encode(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// Hash generates a bcrypt hash of the password
func Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// Verify compares a raw password with a stored hash (supports MD5 legacy and Bcrypt)
func Verify(rawPwd, storedHash string) bool {
	// Legacy MD5 check (32 chars hex)
	if len(storedHash) == 32 {
		return Encode(rawPwd) == storedHash
	}

	// Bcrypt check
	err := bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(rawPwd))
	return err == nil
}
