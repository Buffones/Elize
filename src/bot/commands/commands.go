package commands

import (
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
)

type ElizeCommand struct {
	Command discord.SlashCommandCreate
	Handler func(*handler.CommandEvent) error
}

var Commands = []ElizeCommand{
	pingCommand,
	sayCommand,
	rot13Command,
}

func GetCommandCreates() (commands []discord.ApplicationCommandCreate) {
	for _, c := range Commands {
		commands = append(commands, c.Command)
	}
	return
}
