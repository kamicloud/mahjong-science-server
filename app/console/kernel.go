package console

import (
	"github.com/kamicloud/mahjong-science-server/app"
	"github.com/kamicloud/mahjong-science-server/app/console/commands"
	"github.com/robfig/cron/v3"
)

func init() {
	c := cron.New()
	// 每天3点同步排行
	_, _ = c.AddFunc("* 3 * * *", commands.SyncRank)
	_, _ = c.AddFunc("* * * * *", commands.Spider)
	go c.Start()
	defer c.Stop()

	go startUp()
}

func startUp() {
	if app.Config.Runmode == "prod" {
		commands.Spider()
		commands.SyncRank()
	}
}
