package main

import (
	"log"
	"log/slog"
	"os"
	"path/filepath"
)

var appEnv = os.Getenv("APP_ENV") // $APP_ENV=production go run main.go

// go vet ./...
func main() {
	InitSlog()
	slog.Error("error")
	slog.Warn("warn")
	slog.Info("info")
	slog.Debug("debug")
}

func InitSlog() {

	opts := &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelInfo,
	}

	var handler slog.Handler = slog.NewJSONHandler(os.Stdout, opts)

	if appEnv == "production" {
		const LogDir = "/var/log/lksense/lk-data"
		opts.Level = slog.LevelWarn
		fileName := filepath.Join(LogDir, "test.log")
		file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatal(err)
		}

		handler = slog.NewJSONHandler(file, opts)
	}

	slog.SetDefault(slog.New(handler))
}
