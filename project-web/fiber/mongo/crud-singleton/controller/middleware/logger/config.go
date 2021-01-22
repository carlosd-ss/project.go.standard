package logger

import (
	"os"

	"github.com/gofiber/fiber/middleware"
)

// Config is the configuration for the logs
var Config = middleware.LoggerConfig{
	Next:       nil,
	Format:     "${time} ${method} ${path} - ${ip} - ${status} - ${latency}\n",
	TimeFormat: "15:04:05",
	Output:     os.Stderr,
}
