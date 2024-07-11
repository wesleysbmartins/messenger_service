package console

import (
	"fmt"
	"os"
	"strconv"
)

type Console struct{}

type IConsole interface {
	Print(value any)
}

func (c *Console) Print(value any) {
	print_enabled, _ := strconv.ParseBool(os.Getenv("PRINT_ENABLED"))

	if print_enabled {
		fmt.Println(value)
	}
}
