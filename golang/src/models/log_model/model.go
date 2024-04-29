package logmodel

import "go.mongodb.org/mongo-driver/bson/primitive"

type Log struct {
	ID         primitive.ObjectID `bson:"_id"`
	LogType    string
	Actortype  string
	Actor      interface{}
	TargetType string
	Target     string
	Message    string
	Payload    interface{}
}

type ReportableLog struct {
	ENV     string `json:"env"`
	Message string `json:"message"`
	File    string `json:"file"`
	Line    string `json:"line"`
	Url     string `json:"url"`
	Phone   string `json:"phone"`
	Ip      string `json:"ip"`
	Time    string `json:"time"`
}

func NewLog(logType string, actorType string, actor interface{}, targetType string, target string, message string, payload interface{}) *Log {
	return &Log{
		LogType:    logType,
		Actortype:  actorType,
		Actor:      actor,
		TargetType: targetType,
		Target:     target,
		Message:    message,
		Payload:    payload,
	}
}
