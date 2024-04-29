package usermodel

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID            primitive.ObjectID `bson:"_id"`
	UserName      string             `bson:"username"`
	Password      string             `bson:"password"`
	TelegramBotId string             `bson:"telegram_bot_id"`
	ChatId        string             `bson:"chat_id"`
}

func NewUser(id primitive.ObjectID, userName string, pass string, TBId string, chatId string) *User {
	return &User{
		ID:            id,
		UserName:      userName,
		Password:      pass,
		TelegramBotId: TBId,
		ChatId:        chatId,
	}
}
