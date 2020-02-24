package util

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"hash"
)

const (
	HashMd5 = iota
	HashSha1
	HashSha256
)

type HashHelper struct {
	Algorithm int
}

func NewHashHelper(alg int) *HashHelper {
	return &HashHelper{alg}
}

func (hp HashHelper) EncodeToHexString(data []byte) string {
	var h hash.Hash
	switch hp.Algorithm {
	case HashMd5:
		h = md5.New()
	case HashSha1:
		h = sha1.New()
	case HashSha256:
		h = sha256.New()
	}
	h.Write(data)
	sum := h.Sum(nil)
	return hex.EncodeToString(sum)
}
