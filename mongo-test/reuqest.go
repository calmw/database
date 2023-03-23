package main

import (
	"log"
	"mongo-test/models"
	"runtime/debug"
)

func InsertRequest(record models.AddRequestLogs) error {
	err := models.NewRequestLogs().AddRecord(&record)
	if err != nil {
		log.Println(debug.Stack(), err)
		return err
	}

	return nil
}
