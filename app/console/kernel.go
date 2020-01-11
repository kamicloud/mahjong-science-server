package console

import (
	"github.com/kamicloud/mahjong-science-server/app"
	"github.com/kamicloud/mahjong-science-server/app/console/commands"
	"github.com/robfig/cron/v3"
)

var cronInstance *cron.Cron = nil

func init() {
	go startUp()
}

func startUp() {
	// 初始化
	if app.Config.Runmode == "prod" {
		commands.MajsoulConnector()
		commands.SyncRank()
		commands.Spider()
	}
	cronInstance = cron.New()
	// 每分钟心跳检查
	_, _ = cronInstance.AddFunc("* * * * *", commands.MajsoulConnector)
	// 每天3点同步排行
	_, _ = cronInstance.AddFunc("0 3 * * *", commands.SyncRank)
	// 每分钟拉取观战
	_, _ = cronInstance.AddFunc("* * * * *", commands.Spider)

	go cronInstance.Start()
}

func Stop() {
	if cronInstance != nil {
		cronInstance.Stop()
	}
}
