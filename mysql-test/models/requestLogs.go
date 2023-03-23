package models

import (
	"log"
	"mysql-test/db"
	"time"
)

type RequestLogs struct {
	Id      int64  `json:"-" gorm:"column:id;primaryKey;AUTO_INCREMENT"`
	Path    string `json:"path" gorm:"column:path"` //deposit（0） or withdraw（1）
	Payload string `json:"payload" gorm:"column:payload"`
	Header  string `json:"header" gorm:"column:header"`
	Time    int64  `json:"status" gorm:"column:time"`
}

func NewRequestLogs() *RequestLogs {
	return &RequestLogs{}
}

type ReqRequestLogs struct {
	Page     int `json:"page" `
	PageSize int `json:"pageSize" binding:"numeric"`
}

type AddRequestLogs struct {
	Path    string `json:"path" gorm:"column:path"` //deposit（0） or withdraw（1）
	Payload string `json:"payload" gorm:"column:payload"`
	Header  string `json:"header" gorm:"column:header"`
}

func (m *RequestLogs) Pagination(req *ReqRequestLogs, whereCondition string) (error, int64, []RequestLogs) {
	var total int64
	var requestLogs []RequestLogs

	db.Mysql.Table("request_logs").Where(whereCondition).Count(&total)

	err := db.Mysql.Table("request_logs").Where(whereCondition).Order("id desc").Limit(req.PageSize).Offset((req.Page - 1) * req.PageSize).Find(&requestLogs).Debug().Error
	if err != nil {
		return err, 0, nil
	}

	return nil, total, requestLogs
}

func (m *RequestLogs) AddRecord(req *AddRequestLogs) error {

	data := RequestLogs{}

	data.Path = req.Path
	data.Payload = req.Payload
	data.Header = req.Header
	data.Time = time.Now().Unix()

	err := db.Mysql.Table("request_logs").Create(&data)

	if err.Error != nil {
		log.Println(err.Error.Error())
		return err.Error
	}
	return nil
}
