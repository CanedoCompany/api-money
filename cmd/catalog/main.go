package main

import (
	"github.com/CanedoCompany/api-money/config"
	"github.com/CanedoCompany/api-money/router"
)

var (
	logger *config.Logger
)

func main() {

  logger = config.GetLogger("main")

  err := config.Init()
  if err != nil {
	logger.Errorf("Config initialization error: %v",err)
	return
  }

  router.Initialize()
}