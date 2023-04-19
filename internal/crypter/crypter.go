package crypter

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hash"
	"io"
)

type Crypter interface {
	GenerateHash(str string) string
	CompareHash(hashStr string, str string) bool
}

type CustomCrypter struct {
	hasher hash.Hash
}

func GetDefaultCrypter() *CustomCrypter {
	hasher := sha256.New()
	return &CustomCrypter{
		hasher: hasher,
	}
}

func (c *CustomCrypter) GenerateHash(str string) string {
	c.hasher.Write([]byte(str))
	defer c.hasher.Reset()
	return fmt.Sprintf("%x", c.hasher.Sum(nil))
}

func (c *CustomCrypter) CompareHash(hashStr string, str string) bool {
	toCompare := c.GenerateHash(str)
	return hashStr == toCompare
}

func GenerateSecureToken(length int) (string, error) {
	return GenerateSecureTokenWithReader(length, nil)
}

func GenerateSecureTokenWithReader(length int, randObj io.Reader) (string, error) {
	b := make([]byte, length)
	if randObj == nil {
		randObj = rand.Reader
	}

	if _, err := randObj.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}
