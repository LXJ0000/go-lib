package ratelimit

import (
	"context"
	_ "embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"log/slog"
	"net/http"
	"time"
)

const (
	defaultWindow = time.Second
	defaultLimit  = 100
)

//go:embed slice_window.lua
var LuaSliceWindow string

type Config struct {
	RedisAddr     string
	RedisPassword string
	RedisDB       int
	Window        time.Duration
	Limit         int
}

func DefaultConfig() Config {
	return Config{
		RedisAddr: "127.0.0.1:6379",
		Window:    defaultWindow,
		Limit:     defaultLimit,
	}
}

func New(config Config) gin.HandlerFunc {
	cmd := redis.NewClient(&redis.Options{
		Addr:     config.RedisAddr,
		DB:       config.RedisDB,
		Password: config.RedisPassword,
	})
	if _, err := cmd.Ping(context.Background()).Result(); err != nil {
		slog.Error("Redis Connect Fail With Error Config", "err", err.Error())
		return nil
	}
	return func(context *gin.Context) {
		key := fmt.Sprintf("ip:%s", context.ClientIP())
		isLimit, err := cmd.Eval(context, LuaSliceWindow, []string{key}, config.Window.Milliseconds(), config.Limit, time.Now().UnixMilli()).Bool()
		if err != nil {
			slog.Warn("Redis Op Fail", "error", err.Error())
			context.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		if isLimit {
			slog.Warn("To Many Requests From", "ip", context.ClientIP())
			context.AbortWithStatus(http.StatusServiceUnavailable)
			return
		}
		context.Next()
	}
}
