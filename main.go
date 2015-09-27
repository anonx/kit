// Package main is used as an entry point of
// the framework. It validates user input parameters
// and runs subcommands (aka tools).
package main

import (
	"os"

	"github.com/colegion/goal/commands/create"
	"github.com/colegion/goal/commands/generate"
	"github.com/colegion/goal/commands/run"
	"github.com/colegion/goal/internal/command"
	"github.com/colegion/goal/log"
)

// handlers is a map of registered subcommands (aka tools)
// the framework supports.
var handlers = command.NewContext()

func main() {
	var trace bool

	// Do not show stacktrace if something goes wrong
	// in case tracing is turned off.
	defer func() {
		if !trace {
			if err := recover(); err != nil {
				// Do nothing, error message has already been printed
				// and we do not need stack trace.
			}
		}
	}()

	// Enabling tracing if that is requested by a user.
	command.Helpers["--trace"] = func(val string) {
		if val == "true" {
			trace = true
		}
	}

	// Try to run the subcommand user requested.
	err := Handlers.Process(os.Args[1:]...)
	if err == command.ErrIncorrectArgs { // The arguments were not correct.
		log.Warn.Printf(unknownCmd, os.Args[1])
		return
	}
	if err != nil { // The arguments were omitted.
		Handlers.Process("help", "info") // Show a help message.
		return
	}
}

func init() {
	// Register the supported subcommands.
	Handlers.Register(create.Handler)
	Handlers.Register(run.Handler)
	Handlers.Register(generate.Handler)
	Handlers.Register(helpHandler)

	// Show header message when using new or help
	// commands.
	command.Helpers["new"] = showHeader
	command.Helpers["run"] = showHeader
	command.Helpers["help"] = showHeader
}

// showHeader prints a header message
// with the name of the project.
func showHeader(val string) {
	log.Trace.Println(header)
}

var unknownCmd = `Unknown command "%s".
Run "goal help" for usage.`
