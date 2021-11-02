package main

import (
	"fmt"

	"github.com/Binaretech/classroom-main/internal/server"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.Fatalln(server.App().Listen(fmt.Sprintf(":%s", viper.GetString("port"))))
}
