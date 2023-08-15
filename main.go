package main

import (
	"github.com/kanrichan/clipboard/model"
	"github.com/kanrichan/clipboard/router"
)

func main() {

	r := router.Setup()
	err := model.Setup()
	if err != nil {
		panic(err)
	}
	r.Run("127.0.0.1:8080")
}
