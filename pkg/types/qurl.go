package types

import (
	"pgxs.io/chassis"
	"time"
)

//QUrlDO qurl database object form orm
type QUrlDO struct {
	chassis.SampleBaseDO
	URL       string
	Hash      string
	CreatedAt time.Time  `json:"created_at,omitempty"`             // created time
	DeletedAt *time.Time `sql:"index" json:"deleted_at,omitempty"` //deleted time
}

func (qud QUrlDO) TableName() string {
	return QurlTableName
}

//QUrlCreateReq 创建短连接请求参数
type QUrlCreateReq struct {
	URL string `json:"url,omitempty"`
}

type QUrl struct {
	chassis.BaseDTO
	URL string `json:"url,omitempty" validate:"required,uri"`
}

type QurlResp struct {
	QUrl string `json:"qurl,omitempty"`
}
