package dotenv

import (
	"fmt"

	"github.com/Valgard/godotenv"
)

type DotEnv struct{}

type IDotEnv interface {
	Load()
}

func (d *DotEnv) Load() {
	dotenv := godotenv.New()
	if err := dotenv.Load(".env"); err != nil {
		panic(err)
	}

	fmt.Println("ENVIRONMENTS LOAD SUCCESS!")
}
