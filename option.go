package main

import (
	"flag"
)

type Option struct {
	Verbose bool

	flagSet *flag.FlagSet
}

const commandName = "hclcheck"

func NewOption() *Option {
	ret := &Option{
		flagSet: flag.NewFlagSet(commandName, flag.ContinueOnError),
	}
	ret.flagSet.BoolVar(&ret.Verbose, "v", false, "verbose")
	return ret
}

func (o *Option) Parse(args []string) ([]string, error) {
	if err := o.flagSet.Parse(args); err != nil {
		return nil, err
	}
	return o.flagSet.Args(), nil
}

func (o *Option) PrintDefaults() {
	o.flagSet.PrintDefaults()
}

