package init

import (
	"log"

	"github.com/wrpota/go-gin/internal/global/variable"
	_ "github.com/wrpota/go-gin/internal/pkg/destroy" //退出信号监听
	"github.com/wrpota/go-gin/internal/pkg/log_hook"

	"github.com/wrpota/go-gin/configs"
	cgorm "github.com/wrpota/go-gin/pkg/grom"
	"github.com/wrpota/go-gin/pkg/zap_log"
)

func init() {

	// 初始化全局日志句柄，并载入日志钩子处理函数
	variable.ZapLog = zap_log.New(configs.Get().GetString("Logs.GoLogName"), log_hook.ZapLogHandler)
	variable.GinZapLog = zap_log.New(configs.Get().GetString("Logs.GinLogName"), log_hook.ZapLogHandler)
	//初始化数据库连接
	if dbRead, err := cgorm.GetDbReadDriver(configs.Get().GetString("Database.UseDbType")); err != nil {
		log.Fatal("Gorm 数据库驱动、连接初始化失败" + err.Error())
	} else {
		variable.GormReadDb = dbRead
	}
	if dbWrite, err := cgorm.GetDbWriteDriver(configs.Get().GetString("Database.UseDbType")); err != nil {
		log.Fatal("Gorm 数据库驱动、连接初始化失败" + err.Error())
	} else {
		variable.GormWriteDb = dbWrite
	}
}
