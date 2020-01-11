package console

import (
	"time"

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

	cronInstance = cron.New(cron.WithLocation(time.FixedZone("CST", 8*3600)))

	// 每分钟心跳检查
	_, _ = cronInstance.AddFunc("* * * * *", commands.MajsoulConnector)
	// 每天3点同步排行
	_, _ = cronInstance.AddFunc("0 3 * * *", commands.SyncRank)
	// 每分钟拉取观战
	_, _ = cronInstance.AddFunc("* * * * *", commands.Spider)
	// 每30分钟拉取完整牌谱
	_, _ = cronInstance.AddFunc("*/30 * * * *", commands.RecordDownloader)

	go cronInstance.Start()
}

// Stop 停止定时任务
func Stop() {
	if cronInstance != nil {
		cronInstance.Stop()
	}
}
