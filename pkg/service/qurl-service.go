package service

import (
	"errors"
	lru "github.com/hashicorp/golang-lru"
	"pgxs.io/qurl/pkg/config"
	"sync"

	"pgxs.io/chassis"
	log2 "pgxs.io/chassis/log"

	"pgxs.io/qurl/pkg/repository"
	"pgxs.io/qurl/pkg/types"
	"pgxs.io/qurl/pkg/util"
)

type QUrlService struct {
	log        *log2.Entry
	repo       *repository.QUrlRepository
	base62     *util.Base62
	hashHelper *util.HashHelper
	cache      *lru.Cache
}

const (
	UrlExistedErr = "url existed"
)

var (
	qurlsvc     *QUrlService
	qurlsvcOnce sync.Once
)

//QUrlServiceInstance qurl service 单例
func QUrlServiceInstance() *QUrlService {
	qurlsvcOnce.Do(func() {
		qurlsvc = new(QUrlService)
		qurlsvc.log = log2.New().Category("service").Component("qurl")
		qurlsvc.repo = repository.QurlRepositoryInstance()
		qurlsvc.base62 = util.NewBase62()
		qurlsvc.hashHelper = util.NewHashHelper(util.HashMd5)
		var err error
		cacheSize := 100
		if config.Server().Qurl.CacheSize != 0 {
			cacheSize = config.Server().Qurl.CacheSize
			qurlsvc.log.Infof("使用配置文件中的cache size: %d", cacheSize)
		} else {
			qurlsvc.log.Infof("使用默认cache size: %d", cacheSize)
		}

		qurlsvc.cache, err = lru.New(cacheSize)
		if err != nil {
			qurlsvc.log.Error("qurl new cache error:")
		}
	})

	return qurlsvc
}

//GetURL 获取URL
func (qs QUrlService) GetURL(sName string) (url string) {

	//cache 初始化成功 先从cache里取出
	if qs.cache != nil {
		if cUrl, ok := qs.cache.Get(sName); ok {
			url = cUrl.(string)
			qs.log.Info("url found in cache")
			return
		}
		//从数据库查询
		url = qs.getFromDB(sName)
		if url != "" {
			qs.cache.Add(sName, url)
		}
		return
	}

	url = qs.getFromDB(sName)
	return
}
func (qs QUrlService) getFromDB(sName string) (url string) {
	id := qs.base62.Decode(sName)
	urlDO := qs.repo.FindOne(id)
	if urlDO != nil {
		url = urlDO.URL
	}
	return
}

//Save 保存短连接
//对URL进行MD5 hash去中间16位
//保存时从数据库判重，如存在hash相同的则比对其url是否相同（防hash碰撞）
//如已存在则返回已存在url
func (qs QUrlService) Save(url *types.QUrl) error {
	//截取md5 hash值 中间16位
	hash := qs.hashHelper.EncodeToHexString([]byte(url.URL))[8:24]
	urls := qs.repo.FindByHash(hash)
	if len(urls) > 0 {
		for _, u := range urls {
			if url.URL == u.URL {
				chassis.Copy(url, u)
				return errors.New(UrlExistedErr)
			}
		}
	}
	var urlDO types.QUrlDO
	chassis.Copy(&urlDO, url)
	(&urlDO).Hash = hash
	if err := qs.repo.Save(&urlDO); err != nil {
		return errors.New("save url to db failed")
	}
	chassis.Copy(url, &urlDO)
	return nil
}
