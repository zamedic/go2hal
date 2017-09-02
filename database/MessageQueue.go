package database

import "gopkg.in/mgo.v2/bson"

type messageDB struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	MessageString   string
	ChatID    int64
	MessageID int
}

type MessageDTO struct {
	Message   string
	ChatID    int64
	MessageID int
}

func AddMessageToQueue(message string, chatID int64, messageID int) {
	msg := messageDB{MessageString: message, MessageID: messageID, ChatID: chatID}
	c := database.C("MessageQueue")
	c.Insert(msg)
}

func GetMessages() []MessageDTO {
	c := database.C("MessageQueue")
	var messages []messageDB
	c.Find(nil).All(&messages)

	result := make([]MessageDTO, len(messages))

	for i, x := range messages {
		result[i].MessageID = x.MessageID
		result[i].ChatID = x.ChatID
		result[i].Message = x.MessageString
		c.RemoveId(x.ID)
	}

	return result
}
