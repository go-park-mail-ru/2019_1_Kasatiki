package main

import (
	"github.com/go-park-mail-ru/2019_1_Kasatiki/pkg/app/server"
)

func main() {
	advhater := server.App{}
	advhater.Initialize()
	advhater.Run(":8080")
}
