package entities

type MessageSend struct {
	Recipient Recipient `json:"recipient" bson:"recipient"`
	Message   struct {
		Text       string     `json:"text" bson:"text"`
		Attachment attachment `json:"attachment" bson:"attachment"`
	} `json:"message" bson:"message"`
}

type attachment struct {
	Type    string  `json:"type" bson:"type"`
	Payload payload `json:"payload" bson:"payload"`
}

type payload struct {
	TemplateType string   `json:"template_type" bson:"template_type"`
	Text         string   `json:"text" bson:"text"`
	Buttons      []button `json:"buttons" bson:"buttons"`
}

type button struct {
	Type    string `json:"type" bson:"type"`
	Title   string `json:"title" bson:"title"`
	Payload string `json:"payload" bson:"payload"`
}
