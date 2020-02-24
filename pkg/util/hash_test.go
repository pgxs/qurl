package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHashHelper(t *testing.T) {
	hh := NewHashHelper(HashMd5)
	assert.NotNil(t, hh)
	assert.Equal(t, HashMd5, hh.Algorithm)
}

func TestHashHelper_EncodeToHexString(t *testing.T) {
	md5 := NewHashHelper(HashMd5)
	rs1 := md5.EncodeToHexString([]byte("123456"))
	assert.Equal(t, "e10adc3949ba59abbe56e057f20f883e", rs1)
}
