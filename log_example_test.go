package zerolog_test

import (
	"errors"
	"os"

	"github.com/rs/zerolog"
)

func ExampleNew() {
	log := zerolog.New(os.Stdout)

	log.Info().Msg("hello world")

	// Output: {"level":"info","message":"hello world"}
}

func ExampleLogger_With() {
	log := zerolog.New(os.Stdout).
		With().
		Str("foo", "bar").
		Logger()

	log.Info().Msg("hello world")

	// Output: {"level":"info","foo":"bar","message":"hello world"}
}

func ExampleLogger_Level() {
	log := zerolog.New(os.Stdout).Level(zerolog.WarnLevel)

	log.Info().Msg("filtered out message")
	log.Error().Msg("kept message")

	// Output: {"level":"error","message":"kept message"}
}

func ExampleLogger_Sample() {
	log := zerolog.New(os.Stdout).Sample(2)

	log.Info().Msg("message 1")
	log.Info().Msg("message 2")
	log.Info().Msg("message 3")
	log.Info().Msg("message 4")

	// Output: {"level":"info","sample":2,"message":"message 2"}
	// {"level":"info","sample":2,"message":"message 4"}
}

func ExampleLogger_Debug() {
	log := zerolog.New(os.Stdout)

	log.Debug().
		Str("foo", "bar").
		Int("n", 123).
		Msg("hello world")

	// Output: {"level":"debug","foo":"bar","n":123,"message":"hello world"}
}

func ExampleLogger_Info() {
	log := zerolog.New(os.Stdout)

	log.Info().
		Str("foo", "bar").
		Int("n", 123).
		Msg("hello world")

	// Output: {"level":"info","foo":"bar","n":123,"message":"hello world"}
}

func ExampleLogger_Warn() {
	log := zerolog.New(os.Stdout)

	log.Warn().
		Str("foo", "bar").
		Msg("a warning message")

	// Output: {"level":"warning","foo":"bar","message":"a warning message"}
}

func ExampleLogger_Error() {
	log := zerolog.New(os.Stdout)

	log.Error().
		Err(errors.New("some error")).
		Msg("error doing something")

	// Output: {"level":"error","error":"some error","message":"error doing something"}
}

func ExampleLogger_Log() {
	log := zerolog.New(os.Stdout)

	log.Log().
		Str("foo", "bar").
		Str("bar", "baz").
		Msg("")

	// Output: {"foo":"bar","bar":"baz"}
}
