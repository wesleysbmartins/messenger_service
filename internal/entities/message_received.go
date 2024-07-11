package entities

type MessageReceived struct {
	Id        string `json:"id" bson:"id"`
	Time      int    `json:"time" bson:"time"`
	Messaging []struct {
		Sender    Sender    `json:"sender" bson:"sender"`
		Recipient Recipient `json:"recipient" bson:"recipient"`
		Timestamp int       `json:"timestamp" bson:"timestamp"`
		Message   message   `json:"message" bson:"message"`
		Postback  postback  `json:"postback" bson:"postback"`
	} `json:"messaging" bson:"messaging"`
}

type message struct {
	Mid  string `json:"mid" bson:"mid"`
	Text string `json:"text" bson:"text"`
}

type postback struct {
	Mid     string `json:"mid" bson:"mid"`
	Payload string `json:"payload" bson:"payload"`
}
