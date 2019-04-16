package server

import (
	"github.com/go-park-mail-ru/2019_1_Kasatiki/internal/pkg/app"
)

func main() {
	advhater := App{}
	advhater.Initialize()
	advhater.Run(":8080")
}
