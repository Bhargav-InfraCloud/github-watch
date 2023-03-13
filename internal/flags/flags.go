package flags

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/pflag"
)

type Manager struct {
	baseFlagSet *pflag.FlagSet
}

func NewManager(flagSetName string) *Manager {
	return &Manager{
		baseFlagSet: pflag.NewFlagSet(flagSetName, pflag.ContinueOnError),
	}
}

func (m *Manager) AddFlagSet(fs *FlagSet, flags ...FlagDetails) *Manager {
	for _, flag := range flags {
		flag.addToFlagSet(fs)
	}

	m.baseFlagSet.AddFlagSet(fs.pflagSet)

	return m
}

func (m *Manager) Parse(args []string) error {
	err := m.baseFlagSet.Parse(args)
	if err != nil {
		if errors.Is(err, pflag.ErrHelp) {
			os.Exit(0)
		}

		return fmt.Errorf("failed to parse input flags: %w", err)
	}

	return nil
}
