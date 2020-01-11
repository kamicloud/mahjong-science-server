package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"reflect"
	"strconv"
	"time"

	"github.com/EndlessCheng/mahjong-helper/platform/majsoul/api"
	"github.com/EndlessCheng/mahjong-helper/platform/majsoul/proto/lq"
	"github.com/EndlessCheng/mahjong-helper/platform/majsoul/tool"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/golang/protobuf/proto"
	"github.com/kamicloud/mahjong-science-server/app"
	"github.com/kamicloud/mahjong-science-server/app/utils"
	"github.com/kamicloud/mahjong-science-server/app/utils/majsoul"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
)

type messageWithType struct {
	Name string        `json:"name"`
	Data proto.Message `json:"data"`
}

type FullRecord struct {
	Head    *lq.RecordGame    `json:"head"`
	Details []messageWithType `json:"details"`
}

// RecordDownloader 下载牌谱
func RecordDownloader() {
	now := time.Now()
	gameLiveTypes := utils.GetGameLiveTypes()

	for _, gameLiveType := range *gameLiveTypes {
		id := strconv.Itoa(gameLiveType.ID)
		collection := utils.GetCollection("majsoul", "paipu_"+id)
		cur, err := collection.Find(context.TODO(), bson.M{
			"starttime": bson.M{
				"$lt": now.Unix() - 7200,
			},
			"Done": bson.M{
				"$exists": false,
			},
		})
		if err != nil {
			continue
		}

		for cur.Next(context.TODO()) {
			var result *lq.GameLiveHead
			err := cur.Decode(&result)
			if err != nil {
				log.Fatal(err)
			}
			// do something with result....
			// uuid := "200111-cc4cfd9e-bac9-45f7-abfd-59b5e472d1bc"
			uuid := result.Uuid
			err = download(id, uuid)

			if err != nil {
				fmt.Println(err, uuid)
				continue
			}
			ossClient, err := oss.New("oss-cn-hangzhou.aliyuncs.com", app.Config.Osskey, app.Config.Osssecret)
			if err != nil {
				fmt.Println(err)
				continue
			}

			bucket, err := ossClient.Bucket("kamicloud")
			if err != nil {
				fmt.Println(err)
				continue
			}

			err = bucket.PutObjectFromFile("mahjong-science/records/"+uuid+".json", app.Config.Recordspath+"/"+uuid+".json")
			if err != nil {
				fmt.Println(err)
				continue
			}
			updateRes, err := collection.UpdateOne(context.TODO(), bson.M{
				"uuid": bson.M{
					"$eq": uuid,
				},
			}, bson.M{
				"$set": bson.M{
					"Done": true,
				},
			})
			if err != nil {
				fmt.Println(updateRes, err)
			}
		}

		if err := cur.Err(); err != nil {
			fmt.Println(err)
		}
	}
}

func download(roomID string, uuid string) error {
	client, err := majsoul.GetClient()
	if err != nil {
		fmt.Println(err)
	}

	// 获取具体牌谱内容
	reqGameRecord := lq.ReqGameRecord{
		GameUuid: uuid,
	}
	respGameRecord, err := client.FetchGameRecord(&reqGameRecord)
	if err != nil {
		return err
	}

	// 解析
	data := respGameRecord.Data
	if len(data) == 0 {
		dataURL := respGameRecord.DataUrl
		if dataURL == "" {
			return err
		}
		data, err = tool.Fetch(dataURL)
		if err != nil {
			return err
		}
	}
	detailRecords := lq.GameDetailRecords{}
	if err := api.UnwrapMessage(data, &detailRecords); err != nil {
		return err
	}
	details := []messageWithType{}
	for _, detailRecord := range detailRecords.GetRecords() {
		name, data, err := api.UnwrapData(detailRecord)
		if err != nil {
			return err
		}

		name = name[1:] // 移除开头的 .
		mt := proto.MessageType(name)
		if mt == nil {
			return fmt.Errorf("未找到 %s，请检查代码！", name)
		}
		messagePtr := reflect.New(mt.Elem())
		if err := proto.Unmarshal(data, messagePtr.Interface().(proto.Message)); err != nil {
			return err
		}

		details = append(details, messageWithType{
			Name: name[3:], // 移除开头的 lq.
			Data: messagePtr.Interface().(proto.Message),
		})
	}

	// 保存至本地（JSON 格式）
	parseResult := FullRecord{
		Head:    respGameRecord.Head,
		Details: details,
	}
	jsonData, err := json.Marshal(&parseResult)
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(app.Config.Recordspath+"/"+uuid+".json", jsonData, 0644); err != nil {
		return err
	}

	collection := utils.GetCollection("majsoul", "heads_"+roomID)
	err = storeHead(collection, respGameRecord.Head)
	if err != nil {
		return err
	}

	collection = utils.GetCollection("majsoul", "records_"+roomID)

	err = storeRecord(collection, &parseResult)

	return err
}

func storeHead(collection *mongo.Collection, head *lq.RecordGame) error {
	exists := &lq.GameLiveHead{}

	err := collection.FindOne(context.TODO(), bson.M{
		"uuid": head.Uuid,
	}).Decode(exists)

	if err != nil {
		//data := utils.Struct2Map(*resp1.LiveList[i])
		res, err := collection.InsertOne(context.TODO(), head)
		//res, err := collection.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})
		//id := res.InsertedID
		fmt.Println("StoreGameHead", res, err)
	}

	return err
}

func storeRecord(collection *mongo.Collection, record *FullRecord) error {
	exists := &lq.GameLiveHead{}

	err := collection.FindOne(context.TODO(), bson.M{
		"uuid": record.Head.Uuid,
	}).Decode(exists)

	if err != nil {
		//data := utils.Struct2Map(*resp1.LiveList[i])
		res, err := collection.InsertOne(context.TODO(), record)
		//res, err := collection.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})
		//id := res.InsertedID
		fmt.Println("StoreGameRecord", res, err)
	}

	return err
}
