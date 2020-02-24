package register

import (
	"github.com/emicklei/go-restful"

	qurlCtl "pgxs.io/qurl/pkg/controller"
)

//RegisterWebService 注册QURL web 服务
func RegisterWebService() {
	restful.Add(qurlCtl.QUrlControllerInstance().RedirectWebService())
}
func RegisterAdminWebService() {
	restful.Add(qurlCtl.QUrlControllerInstance().AdminWebService())
}
