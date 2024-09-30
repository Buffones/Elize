package commands

import (
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"

	"buffones/elize/cyphers"
)

var caesarCypherCommand = ElizeCommand{
	Command: discord.SlashCommandCreate{
		Name:        "caesar",
		Description: "Caesar cypher your message with any key you want.",
		Options: []discord.ApplicationCommandOption{
			discord.ApplicationCommandOptionString{
				Name:        "message",
				Description: "The text you want cyphered.",
				Required:    true,
			},
			discord.ApplicationCommandOptionInt{
				Name:        "key",
				Description: "How many places would you like to add for each letter?",
				Required:    true,
			},
		},
	},
	Handler: handleCaesarCypher,
}

func handleCaesarCypher(event *handler.CommandEvent) error {
	data := event.SlashCommandInteractionData()
	return event.Respond(
		discord.InteractionResponseTypeCreateMessage,
		discord.NewMessageCreateBuilder().
			SetContent(cyphers.NewCaesarString(data.String("message"), data.Int("key")).Encoded()).
			SetEphemeral(true).
			Build(),
	)
}
