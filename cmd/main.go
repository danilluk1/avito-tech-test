package cmd

import (
	"github.com/danilluk1/avito-tech/config"
)

func main() {
	config, err := config.New(true)
	if err != nil {
		panic(err)
	}

}
