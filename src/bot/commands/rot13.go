package commands

import (
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"

	"buffones/elize/cyphers"
)

var rot13Command = ElizeCommand{
	Command: discord.SlashCommandCreate{
		Name:        "rot13",
		Description: "Caesar cypher your message with a key of 13.",
		Options: []discord.ApplicationCommandOption{
			discord.ApplicationCommandOptionString{
				Name:        "message",
				Description: "The text you want cyphered.",
				Required:    true,
			},
		},
	},
	Handler: handleRot13,
}

func handleRot13(event *handler.CommandEvent) error {
	data := event.SlashCommandInteractionData()
	return event.Respond(
		discord.InteractionResponseTypeCreateMessage,
		discord.NewMessageCreateBuilder().
			SetContent(cyphers.NewCaesarString(data.String("message")).Encoded()).
			SetEphemeral(true).
			Build(),
	)
}
