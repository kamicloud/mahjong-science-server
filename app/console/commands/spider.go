package commands

import (
	"context"
	"fmt"
	"strconv"
	"sync"

	"github.com/EndlessCheng/mahjong-helper/platform/majsoul/proto/lq"
	"github.com/kamicloud/mahjong-science-server/app/utils"
	"github.com/kamicloud/mahjong-science-server/app/utils/majsoul"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var mutex sync.Mutex

// Spider 牌谱爬虫
func Spider() {
	fmt.Println("Command Spider")

	mutex.Lock()
	defer mutex.Unlock()

	fmt.Println("Command Spider Start")
	spider()
}

func spider() error {
	gameLiveTypes := utils.GetGameLiveTypes()

	client, err := majsoul.GetClient()

	if err != nil {
		return err
	}

	for i := 0; i < len(gameLiveTypes); i++ {
		gameLiveType := gameLiveTypes[i]
		resp, err := client.FetchGameLiveList(&lq.ReqGameLiveList{
			FilterId: uint32(gameLiveType.ID),
		})

		if err != nil {
			return err
		}

		id := strconv.Itoa(gameLiveType.ID)

		fmt.Println("Got " + id + " " + gameLiveType.Name1Chs + " " + gameLiveType.Name2Chs + " paipu " + strconv.Itoa(len(resp.LiveList)))
		for j := 0; j < len(resp.LiveList); j++ {
			gameLive := resp.LiveList[j]
			collection := utils.GetCollection("majsoul", "paipu_"+id)

			storeGameLiveList(collection, gameLive)
		}

	}

	return nil
}

func storeGameLiveList(collection *mongo.Collection, head *lq.GameLiveHead) {
	exists := &lq.GameLiveHead{}

	err := collection.FindOne(context.TODO(), bson.M{
		"uuid": head.Uuid,
	}).Decode(exists)

	if err != nil {
		//data := utils.Struct2Map(*resp1.LiveList[i])
		res, err := collection.InsertOne(context.TODO(), head)
		//res, err := collection.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})
		//id := res.InsertedID
		fmt.Println("StoreGameLive", res, err)
	}

}
