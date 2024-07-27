package ioutil

import (
	"fmt"
	"os"
)

type IoUtil struct{}

type IIoUtil interface {
	WriteFile(path string, value any) error
}

func (i *IoUtil) WriteFile(path string, value any) error {

	if value != nil {
		file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)

		if err != nil {
			message := fmt.Sprintf("ERROR TO OPEN FILE: %s", path)
			fmt.Println(message, err)
			return err
		}

		defer file.Close()

		_, err = file.WriteString(fmt.Sprintf("\n%s\n", value))

		if err != nil {
			fmt.Println("ERROR TO WRITE FILE", err)
			return err
		}

		return nil
	} else {
		panic("Value to Read is Empty")
	}
}
