package controller

import (
	"net/http"
	"net/url"
	"strings"
	"sync"

	"github.com/emicklei/go-restful"

	"pgxs.io/chassis"
	"pgxs.io/chassis/log"

	"pgxs.io/qurl/pkg/apierrors"
	"pgxs.io/qurl/pkg/config"
	"pgxs.io/qurl/pkg/service"
	"pgxs.io/qurl/pkg/types"
	"pgxs.io/qurl/pkg/util"
)

type QUrlController struct {
	log    *log.Entry
	svc    *service.QUrlService
	base62 *util.Base62
}

var (
	qurlController *QUrlController
	qurlOnce       sync.Once
)

func QUrlControllerInstance() *QUrlController {
	qurlOnce.Do(func() {
		qurlController = new(QUrlController)
		qurlController.log = log.New().Category("controller").Component("qurl")
		qurlController.svc = service.QUrlServiceInstance()
		qurlController.base62 = util.NewBase62()
	})
	return qurlController
}

//MobileWebService 移动端接口
func (u QUrlController) MobileWebService() (ws *restful.WebService) {
	ws = new(restful.WebService)
	return
}

//WebService web PC 浏览器接口
func (u QUrlController) WebService() (ws *restful.WebService) {
	ws = new(restful.WebService)
	ws.Path("/api")
	return ws
}

//RedirectWebService 短连接跳转服务 非接口 前缀特殊 故单独设置
func (u QUrlController) RedirectWebService() (ws *restful.WebService) {
	ws = new(restful.WebService)
	ws.Path("/qurl").Produces("text/html")
	ws.Route(ws.GET("{sName}").To(u.getURL))

	tags := []string{"QURLs", "Portal"}
	chassis.AddMetaDataTags(ws, tags)
	return
}

//AdminWebService 管理后台接口
func (u QUrlController) AdminWebService() (ws *restful.WebService) {
	ws = new(restful.WebService)
	ws.Path("/admin/api").Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)
	ws.Route(ws.POST("qurls").To(u.save).Doc("生成短网址").
		Reads(types.QUrlCreateReq{}).
		Returns(201, "success", types.QUrl{}))

	tags := []string{"QURLs", "Admin"}
	chassis.AddMetaDataTags(ws, tags)
	return
}

func (u QUrlController) getURL(req *restful.Request, resp *restful.Response) {
	u.log.Debug("get url")
	sName := req.PathParameter("sName")
	url := u.svc.GetURL(sName)
	if url != "" {
		http.Redirect(resp.ResponseWriter, req.Request, url, 302)
	}
}
func (u QUrlController) save(req *restful.Request, resp *restful.Response) {
	var qurl types.QUrl
	req.ReadEntity(&qurl)
	if err := chassis.ValidateEntityAndWriteResp(resp, &qurl, apierrors.QurlInvalid); err != nil {
		return
	}

	if !(strings.HasPrefix(qurl.URL, "http://") || strings.HasPrefix(qurl.URL, "https://")) {
		if strings.HasPrefix(qurl.URL, config.Server().Qurl.Prefix) {
			chassis.NewResponse(resp).Error(400, apierrors.QurlInBlackList)
			return
		}
		chassis.NewResponse(resp).Error(400, apierrors.QurlInvalid)
		return
	}
	if _, err := url.ParseRequestURI(qurl.URL); err != nil {
		u.log.Debug(err)
		chassis.NewResponse(resp).Error(400, apierrors.QurlInvalid)
		return
	}

	if err := u.svc.Save(&qurl); err != nil {
		if err.Error() == service.UrlExistedErr {
			shortUrl := config.Server().Qurl.Prefix + u.base62.Encode(qurl.ID)
			chassis.NewResponse(resp).Status(200).
				Entity(&chassis.Entity{Data: types.QurlResp{QUrl: shortUrl}, APIError: apierrors.QurlHasExisted})
			return
		}
		chassis.NewResponse(resp).Error(500, apierrors.QurlSaveFailed)
		return
	}
	shortUrl := config.Server().Qurl.Prefix + u.base62.Encode(qurl.ID)
	chassis.NewResponse(resp).Created(types.QurlResp{QUrl: shortUrl})
}
