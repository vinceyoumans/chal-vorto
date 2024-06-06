package slogpkg

import (
	"log"
	"log/slog"
	"os"

	"github.com/vinceyoumans/chal-vorto/vorto/pkg/util"
)

func LogVortoP100(message string) {
	logDir := "../output/slog/P100"
	util.MakeDir(logDir)

	// Open or create log file 1
	logFile1000, err := os.OpenFile(logDir+"/log1000.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file 1: %s", err)
	}
	defer logFile1000.Close()

	handlerOpts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}

	logger1000 := slog.New(slog.NewJSONHandler(logFile1000, handlerOpts))

	regGroup1000 := slog.Group(
		"log1000",
	)

	ll1000 := logger1000.With(regGroup1000)

	slog.SetDefault(logger1000)
	ll1000.Info(message)

}
