package app

import (
	"fmt"

	"github.com/jinzhu/configor"
)

func init() {
	configor.Load(&Config, "conf/config.yml")
	fmt.Printf("config: %#v", Config)
}
