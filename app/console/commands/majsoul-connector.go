package commands

import (
	"github.com/kamicloud/mahjong-science-server/app/utils/majsoul"
	"github.com/sirupsen/logrus"
)

type MajsoulConnector struct {
	baseCommand BaseCommand
}

func (majsoulConnector *MajsoulConnector) Handle() {
	logrus.Info("Command MajsoulConnector")

	majsoulConnector.baseCommand.mutex.Lock()
	defer majsoulConnector.baseCommand.mutex.Unlock()

	majsoulConnector.baseCommand.Handle(majsoulConnector.handle)
	logrus.Info("Command MajsoulConnector Done")
}

// MajsoulConnector 连接雀魂服务器
func (majsoulConnector *MajsoulConnector) handle() {
	defer func() {
		if info := recover(); info != nil {
			majsoul.Close()
		}
	}()

	_, err := majsoul.GetClient(true)

	if err != nil {
		logrus.Error("WS连接失败", err)
		return
	}

	if !majsoul.CheckLogin() {
		err = majsoul.Login()
	}

	if err != nil {
		logrus.Error("Login检查失败", err)
	}
}
