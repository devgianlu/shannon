package shannon

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var key = []byte{0x65, 0x87, 0xd8, 0x8f, 0x6c, 0x32, 0x9d, 0x8a, 0xe4, 0x6b}

func TestShannonCreate(t *testing.T) {
	assert := assert.New(t)
	s := ShannonNew(key)
	assert.NotNil(s)
}

func TestShannon_Encrypt(t *testing.T) {
	assert := assert.New(t)
	plaintext := []byte("Hello World")
	s := ShannonNew(key)
	s.Encrypt(plaintext)
	assert.Equal([]byte{0x94, 0x81, 0xe5, 0xa9, 0x5f, 0x93, 0x5e, 0xcb, 0x6c, 0xb5, 0x24}, plaintext)

}

func TestShannon_Decrypt(t *testing.T) {
	assert := assert.New(t)
	ciphertext := []byte{0x94, 0x81, 0xe5, 0xa9, 0x5f, 0x93, 0x5e, 0xcb, 0x6c, 0xb5, 0x24}
	s := ShannonNew(key)
	s.Decrypt(ciphertext)
	assert.Equal([]byte("Hello World"), ciphertext)
}

func TestShannon_Finish(t *testing.T) {
	assert := assert.New(t)
	plaintext := []byte("Hello World")
	s := ShannonNew(key)
	s.Encrypt(plaintext)
	mac := make([]byte, n)
	s.Finish(mac)
	assert.Equal([]byte{0x43, 0x23, 0x86, 0x24, 0xf3, 0xc9, 0xc, 0x58, 0x79, 0xf4, 0xd3, 0xef, 0x83, 0x98, 0x2e, 0x4e}, mac)
}

func TestShannon_CheckMac(t *testing.T) {
	assert := assert.New(t)
	ciphertext := []byte{0x94, 0x81, 0xe5, 0xa9, 0x5f, 0x93, 0x5e, 0xcb, 0x6c, 0xb5, 0x24}
	s := ShannonNew(key)
	s.Decrypt(ciphertext)
	err := s.CheckMac([]byte{0x43, 0x23, 0x86, 0x24, 0xf3, 0xc9, 0xc, 0x58, 0x79, 0xf4, 0xd3, 0xef, 0x83, 0x98, 0x2e, 0x4e})
	assert.Nil(err)
}
