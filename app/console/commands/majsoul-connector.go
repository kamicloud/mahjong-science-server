package commands

import "github.com/kamicloud/mahjong-science-server/app/utils/majsoul"

import "fmt"

// MajsoulConnector 连接雀魂服务器
func MajsoulConnector() {
	fmt.Println("Command MajsoulConnector")
	_, err := majsoul.GetClient()

	if err != nil {
		fmt.Println(err)
	}

	err = majsoul.Login()

	if err != nil {
		fmt.Println(err)
	}
}
