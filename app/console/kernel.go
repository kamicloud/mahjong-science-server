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
	majsoulConnector := commands.MajsoulConnector{}
	recordDownloader := commands.RecordDownloader{}
	spiderCommand := commands.SpiderCommand{}
	syncRankCommand := commands.SyncRank{}

	// 初始化
	if app.Config.Runmode == "prod" {
		majsoulConnector.Handle()
		syncRankCommand.Handle()
		spiderCommand.Handle()
		recordDownloader.Handle()
	}

	cronInstance = cron.New(cron.WithLocation(time.FixedZone("CST", 8*3600)))

	// 心跳检查
	_, _ = cronInstance.AddFunc("* * * * *", majsoulConnector.Handle)
	// 每天3点同步排行
	_, _ = cronInstance.AddFunc("0 3 * * *", syncRankCommand.Handle)
	// 每分钟拉取观战
	_, _ = cronInstance.AddFunc("* * * * *", spiderCommand.Handle)
	// 拉取完整牌谱
	_, _ = cronInstance.AddFunc("* * * * *", recordDownloader.Handle)

	go cronInstance.Start()
}

// Stop 停止定时任务
func Stop() {
	if cronInstance != nil {
		cronInstance.Stop()
	}
}
