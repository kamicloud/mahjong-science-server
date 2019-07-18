package manager

import (
	"github.com/EndlessCheng/mahjong-helper/util"
)

func RandomTile34() []int {
	var tile34 = make([]int, 34)
	for i := 0; i < 14; i++ {
		util.RandomAddTile(tile34)
	}
	return tile34
}
