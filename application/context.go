package application

import (
	"crypto/sha256"
	"fmt"
	"hash"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Context struct {
	Log     *zap.Logger
	Crypter Crypter
	DB      *gorm.DB
	TimeNow *time.Time
}

type Crypter interface {
	GenerateHash(str string) string
	CompareHash(hashStr string, str string) bool
}

type CustomCrypter struct {
	hasher hash.Hash
}

func BuildCustomCrypter() *CustomCrypter {
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
