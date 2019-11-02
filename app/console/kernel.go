package console

import (
	"github.com/robfig/cron/v3"
	"mahjong-science-server/app/console/commands"
)

func init() {
	c := cron.New()
	// 每天3点同步排行
	_, _ = c.AddFunc("0 0 3 * *", commands.SyncRank)
	c.Start()

	startUp()
}

func startUp() {
	commands.SyncRank()
}
