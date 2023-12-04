package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"

	"github.com/palp1tate/brevinect/service/cron/global"
	"github.com/palp1tate/brevinect/service/cron/initialize"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

func main() {
	initialize.InitConfig()
	initialize.InitLogger()

	zap.S().Info("cron service started")
	c := cron.New()
	_, err := c.AddFunc("@every 3s", func() {
		filepath := ".sql"
		now := time.Now().Format("20060102150405")
		filename := fmt.Sprintf("%s%s", now, filepath)
		cmd := exec.Command("mysqldump",
			"-u"+global.ServerConfig.MySQL.User,
			"-p"+global.ServerConfig.MySQL.Password,
			global.ServerConfig.MySQL.Database,
			"--result-file="+"service/cron/backup/"+filename)
		if err := cmd.Run(); err != nil {
			zap.S().Panic(err)
		}
		zap.S().Infof("%s:backup success", now)
	})
	if err != nil {
		zap.S().Panicf("failed to add cron job: %s", err.Error())
	}
	c.Start()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	c.Stop()
	zap.S().Info("cron service stopped")
}
