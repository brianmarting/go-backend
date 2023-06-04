package logging

import "github.com/rs/zerolog"

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}
