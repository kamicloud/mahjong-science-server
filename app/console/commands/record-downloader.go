package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/EndlessCheng/mahjong-helper/platform/majsoul/api"
	"github.com/EndlessCheng/mahjong-helper/platform/majsoul/proto/lq"
	"github.com/EndlessCheng/mahjong-helper/platform/majsoul/tool"
	"github.com/golang/protobuf/proto"
	"github.com/kamicloud/mahjong-science-server/app"
	"github.com/kamicloud/mahjong-science-server/app/utils"
	"github.com/kamicloud/mahjong-science-server/app/utils/majsoul"
	"github.com/sirupsen/logrus"
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
type RecordDownloader struct {
	baseCommand BaseCommand
}

func (recordDownloader *RecordDownloader) Handle() {
	logrus.Info("Command RecordDownloader")

	recordDownloader.baseCommand.mutex.Lock()
	defer recordDownloader.baseCommand.mutex.Unlock()

	logrus.Info("Command RecordDownloader Start")
	recordDownloader.baseCommand.Handle(recordDownloader.handle)
	logrus.Info("Command RecordDownloader Done")
}

// Handle 处理公共方法
func (recordDownloader *RecordDownloader) handle() {
	now := time.Now()
	gameLiveTypes := utils.GetGameLiveTypes()

	for _, gameLiveType := range *gameLiveTypes {
		roomID := strconv.Itoa(gameLiveType.ID)
		logrus.Info("处理观战类型 " + roomID)
		collection := utils.GetCollection("majsoul", "paipu_"+roomID)
		cur, err := collection.Find(context.TODO(), bson.M{
			"starttime": bson.M{
				"$lt": now.Unix() - 7200,
			},
			"Done": bson.M{
				"$ne": true,
			},
		})
		if err != nil {
			continue
		}

		for cur.Next(context.TODO()) {
			time.Sleep(time.Microsecond * 500)
			var result *lq.GameLiveHead
			err := cur.Decode(&result)
			if err != nil {
				logrus.Error(err)
			}
			// uuid := "200111-cc4cfd9e-bac9-45f7-abfd-59b5e472d1bc"
			uuid := result.Uuid
			logrus.Info("处理完整牌谱 " + roomID + " " + uuid)
			err = recordDownloader.download(roomID, uuid)

			if err != nil {
				logrus.Error(err, uuid)
				return
			}
			logrus.Info("下载成功")

			_, err = collection.UpdateOne(context.TODO(), bson.M{
				"uuid": bson.M{
					"$eq": uuid,
				},
			}, bson.M{
				"$set": bson.M{
					"Done": true,
				},
			})
			if err != nil {
				logrus.Error("更新失败 "+uuid, err)
			} else {
				logrus.Info("更新成功 " + uuid)
			}
		}

		if err := cur.Err(); err != nil {
			fmt.Println(err)
		}
		logrus.Info("处理观战类型完毕 " + roomID)
	}
}

func (recordDownloader *RecordDownloader) download(roomID string, uuid string) error {
	client, err := majsoul.GetClient(false)
	if err != nil {
		return err
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

	parseResult := FullRecord{
		Head:    respGameRecord.Head,
		Details: details,
	}
	_, err = json.Marshal(&parseResult)
	if err != nil {
		return err
	}

	collection := utils.GetCollection("majsoul", "heads_"+roomID)
	err = storeHead(collection, respGameRecord.Head)
	if err != nil {
		return err
	}

	// collection = utils.GetCollection("majsoul", "accounts")
	// storeAccount(collection, respGameRecord.Head)

	collection = utils.GetCollection("majsoul", "records_"+roomID)

	err = recordDownloader.storeRecord(roomID, collection, &parseResult)

	return err
}

func storeHead(collection *mongo.Collection, head *lq.RecordGame) error {
	var err error
	exists := &lq.GameLiveHead{}

	err = collection.FindOne(context.TODO(), bson.M{
		"uuid": head.Uuid,
	}).Decode(exists)

	if err != nil {
		_, err = collection.InsertOne(context.TODO(), head)
	}

	return err
}

func storeAccount(collection *mongo.Collection, head *lq.RecordGame) {
	var err error
	for _, account := range head.Accounts {
		exists := &lq.RecordGame_AccountInfo{}
		filter := bson.M{
			"accountid": account.AccountId,
		}
		err = collection.FindOne(context.TODO(), filter).Decode(exists)

		if err != nil {
			collection.InsertOne(context.TODO(), account)
		} else {
			collection.UpdateOne(context.TODO(), filter, account)
		}
	}
}

func (recordDownloader *RecordDownloader) storeRecord(roomID string, collection *mongo.Collection, record *FullRecord) error {
	var err error
	exists := &lq.GameLiveHead{}

	err = collection.FindOne(context.TODO(), bson.M{
		"uuid": record.Head.Uuid,
	}).Decode(exists)

	if err != nil {
		//data := utils.Struct2Map(*resp1.LiveList[i])
		_, err = collection.InsertOne(context.TODO(), record)
	}

	return err
}

func (recordDownloader *RecordDownloader) buildStorePath(roomID string, uuid string) string {
	// recordsPath/roomID/uuid
	return app.Config.Recordspath + "/" + roomID + "/" + strings.Replace(uuid, "-", "/", -1)
}

func (recordDownloader *RecordDownloader) storeFile(filePath string, data []byte) error {
	var err error
	filePathArr := strings.Split(filePath, "/")
	folderPath := strings.Join(filePathArr[0:len(filePathArr)-1], "/")
	_, err = os.Stat(folderPath)

	if err != nil {
		err = os.MkdirAll(folderPath, os.ModeDir|os.ModePerm)
		if err != nil {
			return err
		}
	}

	err = ioutil.WriteFile(filePath, data, 0644)

	return err
}
