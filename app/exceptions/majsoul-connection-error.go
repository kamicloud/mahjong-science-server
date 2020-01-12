package exceptions

type MajsoulConnectionError struct {
}

func (e *MajsoulConnectionError) Error() string {
	return "连接雀魂服务器失败"
}
