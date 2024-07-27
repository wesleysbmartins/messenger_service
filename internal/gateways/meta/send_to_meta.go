package gateways

import (
	"errors"
	"messenger_service/internal/adapters/http"
	"messenger_service/internal/entities"
	"messenger_service/internal/shared/exceptions"
	"os"
)

type MetaGateway struct {
	url   string
	token string
}

type IMetaGateway interface {
	Load()
	SendMessage(message entities.MessageSend) error
}

var metaGateway *MetaGateway

func (m *MetaGateway) Load() {
	metaGatewayInfos := &MetaGateway{
		url:   os.Getenv("META_URL"),
		token: os.Getenv("META_TOKEN"),
	}

	metaGateway = metaGatewayInfos
}

func (m *MetaGateway) SendMessage(message entities.MessageSend) error {
	client := &http.Client{}
	exception := &exceptions.HttpException{}

	url := metaGateway.url
	//if metaGateway.token != "" {
	//	url = fmt.Sprintf("%s/message?%s", metaGateway.url, metaGateway.token)
	//}

	resp, err := client.Post(url, message)

	if resp.StatusCode > 201 || err != nil {
		exception.Handle("ERROR TO SEND MESSAGE TO META", resp, err)
		if err == nil {
			err = errors.New("ERROR TO SEND MESSAGE TO META")
		}
	}

	return err
}
