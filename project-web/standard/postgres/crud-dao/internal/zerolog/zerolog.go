package zerolog

import (
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// panic (zerolog.PanicLevel, 5)
// fatal (zerolog.FatalLevel, 4)
// error (zerolog.ErrorLevel, 3)
// warn (zerolog.WarnLevel, 2)
// info (zerolog.InfoLevel, 1)
// debug (zerolog.DebugLevel, 0)
// trace (zerolog.TraceLevel, -1)

func init() {
	// UNIX Time is faster and smaller than most timestamps
	// If you set zerolog.TimeFieldFormat to an empty string,
	// logs will write with UNIX time
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}

func Info(version, file string, line int, host, method, msg string) {
	log.Info().
		Str("data", time.Now().Format("2006-01-02 15:04:05")).
		Str("version", version).
		Str("file", file).
		Int("line", line).
		Str("host", host).
		Str("method", method).
		Msg(msg)
}

func Error(version, file string, line int, host, method, msg string) {
	log.Error().
		Str("data", time.Now().Format("2006-01-02 15:04:05")).
		Str("version", version).
		Str("file", file).
		Int("line", line).
		Str("host", host).
		Str("method", method).
		Msg(msg)
}
