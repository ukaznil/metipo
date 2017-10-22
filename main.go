package main

import (
	"github.com/jessevdk/go-flags"
	"os"
	"github.com/ukaznil/metipo/core"
)

func main() {
	type Options struct {
		Init     bool `short:"i" long:"init" description:"Change the settings" hidden:"true"`
		Exercise bool `short:"e" long:"exercise" description:"Do exercise with \u001B[0;33;3mMeTipo\u001B[0m" hidden:"false"`
	}

	var opts Options
	var parser = flags.NewParser(&opts, flags.PrintErrors|flags.PassDoubleDash)
	parser.Usage = "[OPTIONS]"
	var help struct {
		ShowHelp bool `short:"h" long:"help" description:"Show this help message"`
	}
	parser.AddGroup("Help Options", "", &help)
	parser.Parse()

	if opts.Init {
		core.InitConfig()
	} else if opts.Exercise {
		core.Exercise()
	} else {
		parser.WriteHelp(os.Stdout)
	}
}
