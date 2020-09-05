package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type MongoDoc struct {
	ID  string
	Doc map[string]interface{}
}
type OplogTimestamp struct {
	LatestOplogTimestamp primitive.Timestamp `json:"latest_oplog_timestamp"`
}

type OpLog struct {
	Timestamp    primitive.Timestamp    "ts"
	HistoryID    int64                  "h"
	MongoVersion int                    "v"
	Operation    string                 "op"
	Namespace    string                 "ns"
	Doc          map[string]interface{} "o"
	Update       map[string]interface{} "o2"
}


