package main

import (
	"github.com/jessevdk/go-flags"
	"github.com/ukaznil/metipo/utils"
	"os"
	"github.com/ukaznil/metipo/core"
)

func main() {
	type Options struct {
		Init     bool `short:"i" long:"init" description:"Change the settings"`
		Exercise bool `short:"e" long:"exercise" description:"Do exercise with MeTipo"`
	}

	var opts Options
	var parser = flags.NewParser(&opts, flags.Default)
	var args, err = parser.Parse()
	utils.Perror(err)

	if len(args) != 0 {
		parser.WriteHelp(os.Stdout)
	} else {
		if opts.Init {
			core.Init()
		} else if opts.Exercise {
			core.Exercise()
		} else {
			parser.WriteHelp(os.Stdout)
		}
	}
}

/*
optarg.Header("Option '-h' or '--help'")
optarg.Add("h", "help", "Show help", "")
optarg.Header("Option '-i' or '--init'")
optarg.Add("i", "init", "Change the settings", "")
var ch = optarg.Parse()
<-ch

for opt := range optarg.Parse() {
	fmt.Print(opt)
}

fmt.Println(len(optarg.Remainder))
switch len(optarg.Remainder) {
case 1:
	for opt := range optarg.Parse() {
		switch opt.ShortName {
		case "h":
			optarg.Usage()
		case "i":
			core.Init()
		}

		switch opt.Name {
		case "help":
			optarg.Usage()
		case "init":
			core.Init()
		}
	}
default:
	optarg.Usage()
}
*/
