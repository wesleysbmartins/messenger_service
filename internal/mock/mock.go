package mock

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"messenger_service/internal/entities"
	"os"
)

type Mock struct {
	Messages []entities.MessageSend
}

type IMock interface {
	Load()
}

var MockSingleton Mock

func (m *Mock) Load() {
	path_file := os.Getenv("MOCK_PATH_FILE")

	file, err := os.Open(path_file)

	if err != nil {
		fmt.Println("ERROR TO LOAD MOCK")
		panic(err)
	}

	bytes, _ := ioutil.ReadAll(file)

	entity := &[]entities.MessageSend{}

	json.Unmarshal(bytes, entity)

	MockSingleton.Messages = *entity

	fmt.Println("MOCK LOAD SUCCESS!")
}
