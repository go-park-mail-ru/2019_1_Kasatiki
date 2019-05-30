package main

import (
	"fmt"
	"github.com/go-park-mail-ru/2019_1_Kasatiki/pkg/app/server"
	"github.com/go-park-mail-ru/2019_1_Kasatiki/pkg/models"
	log "gopkg.in/inconshreveable/log15.v2"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Crit("you must pass the path to the configuration file as the first argument")
		os.Exit(1)
	}
	configBytes, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Crit("unable to open configuration file '" + os.Args[1] + "': " + err.Error())
		os.Exit(1)
	}
	config := &models.Config{}
	err = config.UnmarshalJSON(configBytes)
	if err != nil {
		log.Crit("unable to parse configuration file ./config.json : " +
			"it should be json with all fields : " + err.Error())
		os.Exit(1)
	}
	advhater := server.App{}
	advhater.Initialize(config)
	advhater.Run(fmt.Sprintf(":%d", config.ServerPort))
}
