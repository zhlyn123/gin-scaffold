package bootstrap

import (
	"os"
	"os/signal"
	"syscall"

	"gin-scaffold/internal/logger"
)

func WaitShutdown() {
	// 等待中断信号来优雅地关闭服务器
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	logger.Log.Info("Shutdown Server ...")
	logger.Close()
}
