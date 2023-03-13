package flags

import "github.com/spf13/pflag"

type FlagSet struct {
	pflagSet *pflag.FlagSet
}

func NewFlagSet(name string) *FlagSet {
	return &FlagSet{
		pflagSet: pflag.NewFlagSet(name, pflag.ContinueOnError),
	}
}
