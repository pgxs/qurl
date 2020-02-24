package repository

import (
	"errors"
	"sync"

	"pgxs.io/chassis"
	log2 "pgxs.io/chassis/log"

	"pgxs.io/qurl/pkg/types"
)

//QUrlRepository 短连接数据库持久化对象
type QUrlRepository struct {
	log *log2.Entry
}

var (
	qurlRepo     *QUrlRepository
	qurlRepoOnce sync.Once
)

//QurlRepositoryInstance qurl数据库操作单例
func QurlRepositoryInstance() *QUrlRepository {
	qurlRepoOnce.Do(func() {
		qurlRepo = new(QUrlRepository)
		qurlRepo.log = log2.New().Category("repo").Component("qurl")
	})
	return qurlRepo
}

//FindOne 根据ID查找短连接
func (qur QUrlRepository) FindOne(id uint) *types.QUrlDO {
	qurl := new(types.QUrlDO)
	chassis.DB().First(qurl, id)
	return qurl
}

//Save 保存短连接
func (qur QUrlRepository) Save(qurl *types.QUrlDO) error {
	db := chassis.DB().Save(qurl)
	err := db.Error
	ra := db.RowsAffected
	qur.log.Debugf("影响行数%d", ra)
	if err == nil && ra > 0 {
		return nil
	}
	qur.log.Error(err)
	return errors.New("save qurl failed")
}

//FindByHash 根据hash值查询短连接
func (qur QUrlRepository) FindByHash(hash string) (qurls []types.QUrlDO) {
	db := chassis.DB().Where("hash = ?", hash).Find(&qurls)
	if err := db.Error; err != nil {
		qur.log.Error(db)
	}
	return
}
