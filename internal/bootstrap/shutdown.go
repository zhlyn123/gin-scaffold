package bootstrap

import (
	"os"
	"os/signal"
	"syscall"

	"gin-scaffold/internal/logger"

	"go.uber.org/zap"
)

type ShutdownFunc func() error

var Shutdowns []ShutdownFunc

func RegisterShutdown(f ShutdownFunc) {
	Shutdowns = append(Shutdowns, f)
}

func Shutdown() error {
	for _, f := range Shutdowns {
		if err := f(); err != nil {
			logger.Log.Error(
				"Shutdown error", 
				zap.Error(err),
			)
			return err
		}
	}
	return nil
}

func WaitShutdown() {
	// 等待中断信号来优雅地关闭服务器
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	logger.Log.Info("Shutdown Server ...")
	logger.Close()
}
