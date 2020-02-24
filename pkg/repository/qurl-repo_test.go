package repository

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"pgxs.io/qurl/pkg/types"
	"pgxs.io/qurl/pkg/util"
	"testing"
)

func TestFindOne(t *testing.T) {
	url := qurlRepo.FindOne(10000)
	assert.Equal(t, "https://baidu.com", url.URL)
	hh := util.NewHashHelper(util.HashMd5)
	longHash := hh.EncodeToHexString([]byte(url.URL))
	assert.Equal(t, longHash[8:24], url.Hash)
}

func TestQUrlRepository_Save(t *testing.T) {
	qurl := &types.QUrlDO{URL: "https://liwei.co"}
	hh := util.NewHashHelper(util.HashMd5)
	longHash := hh.EncodeToHexString([]byte(qurl.URL))
	fmt.Println("long hash:", longHash)
	hash := longHash[8:24]
	fmt.Printf("hash \nlong:%s,\nshort:%s\n", longHash, hash)
	qurl.Hash = hash
	err := qurlRepo.Save(qurl)
	assert.NoError(t, err)
}

func TestQUrlRepository_FindByHash(t *testing.T) {
	urls := qurlRepo.FindByHash("5c360ce6a0c64f92")
	assert.NotEmpty(t, urls)
	assert.Equal(t, 1, len(urls))
	assert.Equal(t, uint(10000), urls[0].ID)
	assert.Equal(t, "https://baidu.com", urls[0].URL)
}
