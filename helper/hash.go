package helper

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"os"
	"time"
)

var signature = []byte("myPrivateSignateure")

// GetAESEncrypted to Encrypt Password using CBC 256
// plaintext as string
// return encryptedPwd as String, error
func GetAESEncrypted(plaintext string) (string, error) {
	err := godotenv.Load("config/.env")
	PanicIfError(err)
	key := os.Getenv("MYKEY")
	iv := os.Getenv("MYINITVECTOR")

	var plainTextBlock []byte
	length := len(plaintext)

	if length%16 != 0 {
		extendBlock := 16 - (length % 16)
		plainTextBlock = make([]byte, length+extendBlock)
		copy(plainTextBlock[length:], bytes.Repeat([]byte{uint8(extendBlock)}, extendBlock))
	} else {
		plainTextBlock = make([]byte, length)
	}

	copy(plainTextBlock, plaintext)
	block, err := aes.NewCipher([]byte(key))

	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, len(plainTextBlock))
	mode := cipher.NewCBCEncrypter(block, []byte(iv))
	mode.CryptBlocks(ciphertext, plainTextBlock)

	str := hex.EncodeToString(ciphertext)
	return str, nil
}

// GetAESDecrypted to Decrypt Password using CBC 256
// encrypted as string
// return decryptedPwd as slice byte, error
func GetAESDecrypted(encrypted string) ([]byte, error) {
	err := godotenv.Load("config/.env")
	PanicIfError(err)
	key := os.Getenv("MYKEY")
	iv := os.Getenv("MYINITVECTOR")

	ciphertext, err := hex.DecodeString(encrypted)
	PanicIfError(err)

	block, err := aes.NewCipher([]byte(key))
	PanicIfError(err)

	if len(ciphertext)%aes.BlockSize != 0 {
		return nil, fmt.Errorf("block size cant be zero")
	}

	mode := cipher.NewCBCDecrypter(block, []byte(iv))
	mode.CryptBlocks(ciphertext, ciphertext)
	ciphertext = PKCS5UnPadding(ciphertext)

	return ciphertext, nil
}

// PKCS5UnPadding to
// src as slice byte
// return decryptedPwd as slice byte
func PKCS5UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])

	return src[:(length - unpadding)]
}

func GenerateJWT(email string, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_email": email,
		"exp":        time.Now().Add(24 * time.Hour).Unix(),
		"role":       role,
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	stringToken, err := accessToken.SignedString(signature)
	return stringToken, err
}

func DecryptJWT(token string) (map[string]interface{}, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}
		return signature, nil
	})
	data := make(map[string]interface{})
	if err != nil {
		return data, err
	}

	if !parsedToken.Valid {
		return data, errors.New("invalid token")
	}
	return parsedToken.Claims.(jwt.MapClaims), nil
}
