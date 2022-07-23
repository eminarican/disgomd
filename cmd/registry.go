package cmd

import (
	"sync"

	"github.com/eminarican/safetypes"
)

var commands sync.Map

func Register(command Command) {
	commands.Store(command.name, command)
	for _, alias := range command.aliases {
		commands.Store(alias, command)
	}
}

func ByAlias(alias string) (opt safetypes.Option[Command]) {
	command, ok := commands.Load(alias)
	if !ok {
		return opt.None()
	}
	return opt.Some(command.(Command))
}

func Commands() map[string]Command {
	cmd := make(map[string]Command)
	commands.Range(func(key, value any) bool {
		cmd[key.(string)] = value.(Command)
		return true
	})
	return cmd
}
