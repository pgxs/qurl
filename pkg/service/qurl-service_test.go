package service

import (
	"github.com/stretchr/testify/assert"
	"pgxs.io/qurl/pkg/types"
	"testing"
)

func TestQUrlService_Save(t *testing.T) {

	url := &types.QUrl{
		URL: "https://baidu.com",
	}
	err := qurlsvc.Save(url)
	assert.Error(t, err)
	url2 := &types.QUrl{
		URL: "https://i.liwei.im",
	}
	err = qurlsvc.Save(url2)
	assert.NoError(t, err)
	assert.NotEmpty(t, url2.ID)
}
