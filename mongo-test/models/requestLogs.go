package models

import (
	"github.com/kamva/mgm/v3"
	"log"
	"time"
)

type RequestLogs struct {
	mgm.IDField `json:",inline" bson:",inline"`
	Path        string `json:"path" bson:"path"` //deposit（0） or withdraw（1）
	Payload     string `json:"payload" bson:"payload"`
	Header      string `json:"header" bson:"header"`
	Time        int64  `json:"status" bson:"time"`
}

func NewRequestLogs() *RequestLogs {
	return &RequestLogs{}
}

type ReqRequestLogs struct {
	Page     int `json:"page" `
	PageSize int `json:"pageSize" binding:"numeric"`
}

type AddRequestLogs struct {
	Path    string `json:"path" gorm:"path"` //deposit（0） or withdraw（1）
	Payload string `json:"payload" gorm:"payload"`
	Header  string `json:"header" gorm:"header"`
}

//func (m *RequestLogs) Pagination(ctx context.Context, skip, limit int64, filter interface{}) (ats []*RequestLogs, count int64, err error) {
//	var requestLogs []*RequestLogs
//	opt := &options.FindOptions{}
//	err = mgm.Coll(&RequestLogs{}).SimpleFindWithCtx(ctx, &requestLogs, filter, opt)
//	if err != nil {
//		log.Printf("SimpleFindWithCtx error: %v", err)
//		return
//	}
//	count, err = mgm.Coll(&RequestLogs{}).CountDocuments(ctx, filter)
//	if err != nil {
//		log.Printf("CountDocuments error: %v", err)
//		return
//	}
//	if len(requestLogs) == 0 {
//		return nil, 0, nil
//	}
//	tokenColl := mgm.Coll(&commonModel.Token{}).Name()
//	matchStage := bson.D{{operator.Match, filter}}
//	sortStage := bson.D{{operator.Sort, bson.D{{"createdAt", -1}}}}
//	pageStage := bson.D{{operator.Skip, skip}}
//	limitStage := bson.D{{operator.Limit, limit}}
//	lookupStage := bson.D{{operator.Lookup, bson.D{
//		{"from", tokenColl},
//		{"localField", "token"},
//		{"foreignField", field.ID},
//		{"as", "tokenDetail"}}}}
//	unwindStage := bson.D{{"$unwind", bson.D{
//		{"path", "$tokenDetail"},
//		{"preserveNullAndEmptyArrays", false}}}}
//	cursor, err := mgm.Coll(&commonModel.TokenAccount{}).Aggregate(mgm.Ctx(), mongo.Pipeline{matchStage, sortStage, pageStage, limitStage, lookupStage, unwindStage})
//	if err != nil {
//		log.Errorf("Aggregate error: %v", err)
//		return nil, 0, err
//	}
//	defer func() {
//		_ = cursor.Close(nil)
//	}()
//	for cursor.Next(ctx) {
//		var value = &model.AccountWithToken{}
//		if err = cursor.Decode(value); err != nil {
//			return nil, 0, err
//		}
//		ats = append(ats, value)
//	}
//	return
//}

func (m *RequestLogs) AddRecord(req *AddRequestLogs) error {

	data := &RequestLogs{}

	data.Path = req.Path
	data.Payload = req.Payload
	data.Header = req.Header
	data.Time = time.Now().Unix()

	err := mgm.Coll(data).Create(data)
	if err != nil {
		log.Printf("insert transaction failed. err:%s", err.Error())
		return err
	}
	return nil
}
