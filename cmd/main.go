package main

import (
	"log"

	"github.com/EvgenijMihievich/home/internal/config"
	"go.uber.org/zap"
)

func main() {
	cfg := config.InitConfig("./configs/config.yaml")

	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("Логгер сломан: %s", err)
	}
	defer logger.Sync()

	logger.Info(cfg.DB.Dbfolder)
}
