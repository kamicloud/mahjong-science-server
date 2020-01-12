package commands

import (
	"context"
	"strconv"

	"github.com/EndlessCheng/mahjong-helper/platform/majsoul/proto/lq"
	"github.com/kamicloud/mahjong-science-server/app/utils"
	"github.com/kamicloud/mahjong-science-server/app/utils/majsoul"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// SpiderCommand 牌谱爬虫
type SpiderCommand struct {
	baseCommand BaseCommand
}

// Handle 牌谱爬虫
func (spiderCommand *SpiderCommand) Handle() {
	logrus.Info("Command Spider")

	spiderCommand.baseCommand.mutex.Lock()
	defer spiderCommand.baseCommand.mutex.Unlock()

	logrus.Info("SpiderCommand Start")
	spiderCommand.baseCommand.Handle(spiderCommand.handle)
	logrus.Info("SpiderCommand Done")
}

// handle 牌谱爬虫
func (spiderCommand *SpiderCommand) handle() {
	spider()
}

func spider() error {
	gameLiveTypes := utils.GetGameLiveTypes()

	client, err := majsoul.GetClient(false)

	if err != nil {
		return err
	}

	for _, gameLiveType := range *gameLiveTypes {
		resp, err := client.FetchGameLiveList(&lq.ReqGameLiveList{
			FilterId: uint32(gameLiveType.ID),
		})

		if err != nil {
			return err
		}

		id := strconv.Itoa(gameLiveType.ID)

		logrus.Info("Got " + id + " " + gameLiveType.Name1Chs + " " + gameLiveType.Name2Chs + " paipu " + strconv.Itoa(len(resp.LiveList)))
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
		logrus.Info("StoreGameLive", res, err)
	}

}
