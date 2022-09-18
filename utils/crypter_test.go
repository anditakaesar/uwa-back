package utils

import (
	"errors"
	"testing"
	"testing/iotest"

	"github.com/stretchr/testify/assert"
)

func Test_GetDefaultCrypter(t *testing.T) {
	t.Run("test GetDefaultCrypter not nil success", func(t *testing.T) {
		crypter := GetDefaultCrypter()

		assert.NotNil(t, crypter)
	})
}

func Test_GenerateHash(t *testing.T) {
	t.Run("test GenerateHash generate hash success", func(t *testing.T) {
		crypter := GetDefaultCrypter()

		hash := crypter.GenerateHash("some string")

		assert.NotEmpty(t, hash)
	})
}

func Test_CompareHash(t *testing.T) {
	t.Run("test GenerateHash generate hash success", func(t *testing.T) {
		crypter := GetDefaultCrypter()

		decodedStr := "some string"
		testedHash := "61d034473102d7dac305902770471fd50f4c5b26f6831a56dd90b5184b3c30fc"
		assert.Equal(t, true, crypter.CompareHash(testedHash, decodedStr))
	})
}

func Test_GenerateSecureToken(t *testing.T) {
	t.Run("test GenerateSecureToken not nil success", func(t *testing.T) {
		gotVal, gotErr := GenerateSecureToken(1)

		assert.Nil(t, gotErr)
		assert.NotNil(t, gotVal)
	})
}

func Test_GenerateSecureTokenWithReader(t *testing.T) {
	t.Run("test GenerateSecureTokenWithReader error", func(t *testing.T) {
		errEreader := iotest.ErrReader(errors.New("some error"))
		_, gotErr := GenerateSecureTokenWithReader(1, errEreader)

		assert.NotNil(t, gotErr)
	})
}
