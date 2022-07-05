package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

// LogConfig defines the config for middleware.
type LogConfig struct {
	// Next defines a function to skip this middleware when returned true.
	//
	// Optional. Default: nil
	Next func(c *fiber.Ctx) bool

	// Format defines the logging tags
	//
	// Optional. Default: [${time}] ${status} - ${latency} ${method} ${path}\n
	Format string

	// TimeFormat https://programming.guide/go/format-parse-string-time-date-example.html
	//
	// Optional. Default: 15:04:05
	TimeFormat string

	// TimeZone can be specified, such as "UTC" and "America/New_York" and "Asia/Chongqing", etc
	//
	// Optional. Default: "Local"
	TimeZone string

	// TimeInterval is the delay before the timestamp is updated
	//
	// Optional. Default: 500 * time.Millisecond
	TimeInterval time.Duration

	// Output is a writer where logs are written
	//
	// Default: os.Stdout
	//Output io.Writer

	//enableColors     bool
	enableLatency    bool
	timeZoneLocation *time.Location
}

// LogConfigDefault is the default config
var LogConfigDefault = LogConfig{
	Next:         nil,
	Format:       "[${time}] ${status} - ${latency} ${method} ${path}\n",
	TimeFormat:   "15:04:05",
	TimeZone:     "Local",
	TimeInterval: 500 * time.Millisecond,
}

// Helper function to set default values
func logConfigDefault(config ...LogConfig) LogConfig {
	// Return default config if nothing provided
	if len(config) < 1 {
		return LogConfigDefault
	}

	// Override default config
	cfg := config[0]

	// Set default values
	if cfg.Next == nil {
		cfg.Next = LogConfigDefault.Next
	}
	if cfg.Format == "" {
		cfg.Format = LogConfigDefault.Format
	}
	if cfg.TimeZone == "" {
		cfg.TimeZone = LogConfigDefault.TimeZone
	}
	if cfg.TimeFormat == "" {
		cfg.TimeFormat = LogConfigDefault.TimeFormat
	}
	if int(cfg.TimeInterval) <= 0 {
		cfg.TimeInterval = LogConfigDefault.TimeInterval
	}
	return cfg
}
