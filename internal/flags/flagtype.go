package flags

import "time"

type FlagType interface {
	string | bool | int | time.Duration
}

type flag[T FlagType] struct {
	target     *T
	name       string
	shorthand  string
	defaultVal T
	desc       string
}

func NewFlag[T FlagType](target *T, name, shorthand, description string, defaultVal T) *flag[T] {
	return &flag[T]{
		target:     target,
		name:       name,
		shorthand:  shorthand,
		defaultVal: defaultVal,
		desc:       description,
	}
}

type FlagDetails interface {
	addToFlagSet(fs *FlagSet)
}

func (f *flag[T]) addToFlagSet(fs *FlagSet) {
	switch any(f.defaultVal).(type) {
	case string:
		defaultVal := any(f.defaultVal).(string)
		target := any(f.target).(*string)

		if f.shorthand == "" {
			fs.pflagSet.StringVar(target, f.name, defaultVal, f.desc)

			return
		}

		fs.pflagSet.StringVarP(target, f.name, f.shorthand, defaultVal, f.desc)
	case bool:
		defaultVal := any(f.defaultVal).(bool)
		target := any(f.target).(*bool)

		if f.shorthand == "" {
			fs.pflagSet.BoolVar(target, f.name, defaultVal, f.desc)

			return
		}

		fs.pflagSet.BoolVarP(target, f.name, f.shorthand, defaultVal, f.desc)
	case int:
		defaultVal := any(f.defaultVal).(int)
		target := any(f.target).(*int)

		if f.shorthand == "" {
			fs.pflagSet.IntVar(target, f.name, defaultVal, f.desc)

			return
		}

		fs.pflagSet.IntVarP(target, f.name, f.shorthand, defaultVal, f.desc)
	case time.Duration:
		defaultVal := any(f.defaultVal).(time.Duration)
		target := any(f.target).(*time.Duration)

		if f.shorthand == "" {
			fs.pflagSet.DurationVar(target, f.name, defaultVal, f.desc)

			return
		}

		fs.pflagSet.DurationVarP(target, f.name, f.shorthand, defaultVal, f.desc)
	}
}
