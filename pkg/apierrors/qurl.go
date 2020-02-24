package apierrors

import (
	"pgxs.io/chassis"
)

var QurlInvalid = chassis.NewAPIError(10001, "url有误", "url格式不正确，协议为：https:// or http://, 并且已url encode")
var QurlSaveFailed = chassis.NewAPIError(10002, "保存url失败", "未知异常")
var QurlHasExisted = chassis.NewAPIError(10003, "url已存在", "已存在该URL的短连接，无需重复新增")
var QurlInBlackList = chassis.NewAPIError(10004, "URL在黑名单", "要生成的域名在黑名单，不能生成")
