package commands

import (
	"sync"

	"github.com/sirupsen/logrus"
)

// BaseCommand 基类
type BaseCommand struct {
	mutex sync.Mutex
}

// Handle 公共
func (baseCommand *BaseCommand) Handle(callback func()) {
	defer func() {
		if info := recover(); info != nil {
			logrus.Error(info)
		}
	}()
	callback()
}
