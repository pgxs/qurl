package main

import (
	"github.com/emicklei/go-restful"
	"pgxs.io/qurl/cmd/qurl/register"
	"pgxs.io/qurl/pkg/migrations"
	"time"

	"pgxs.io/chassis"
	"pgxs.io/chassis/config"
	"pgxs.io/chassis/log"

	cfg "pgxs.io/qurl/pkg/config"
)

func main() {
	//初始化配置
	time.LoadLocation("Asia/Shanghai")
	cfg.ResetEnvKey()
	config.LoadFromEnvFile()
	cfg.LoadServer()

	log := log.New().Category("cmd").Component("server")
	log.Infof("API server starting on http://127.0.0.1:%d", config.Server().Port)
	log.Infof("API docs on http://127.0.0.1:%d%s", config.Server().Port, config.Openapi().UI.Entrypoint)

	//connect DB: if error will exit
	chassis.DB()
	log.Info("DB connected")
	defer chassis.CloseDB()

	register.RegisterAdminWebService()
	register.RegisterWebService()

	// run db migrations 运行数据库迁移
	if err := migrations.Run(); err != nil {
		log.Info("DB Migrations failed")
		log.Fatalln(err)
	}
	log.Info("DB Migrations success")

	//listen run http server
	chassis.Serve(restful.RegisteredWebServices())

}
