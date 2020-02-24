package util

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBase62(t *testing.T) {
	b62 := NewBase62()
	assert.NotNil(t, b62)
}

func TestBase62_Encode(t *testing.T) {
	b62 := NewBase62()
	fmt.Println(b62.Encode(uint(10000)))
	b62str := b62.Encode(uint(1915810))
	assert.Equal(t, "82oa", b62str)
	assert.Equal(t, uint(36526445), b62.Decode("2tgcd"))
}
