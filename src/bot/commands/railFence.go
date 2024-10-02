package commands

import (
	"buffones/elize/cyphers"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
)

var railFenceCypherCommand = ElizeCommand{
	Command: discord.SlashCommandCreate{
		Name:        "rail-fence",
		Description: "Cypher your message using the rail fence strategy.",
		Options: []discord.ApplicationCommandOption{
			discord.ApplicationCommandOptionString{
				Name:        "message",
				Description: "The text you want cyphered.",
				Required:    true,
			},
			discord.ApplicationCommandOptionInt{
				Name:        "rails",
				Description: "How many rails do you want to use to cypher this message? (Min: 2, Max: 15)",
				Required:    true,
			},
		},
	},
	Handler: handleRailFenceCypher,
}

func handleRailFenceCypher(event *handler.CommandEvent) error {
	data := event.SlashCommandInteractionData()
	message, rails := data.String("message"), data.Int("rails")
	if rails < 2 || rails > 15 {
		return event.Respond(
			discord.InteractionResponseTypeCreateMessage,
			discord.NewMessageCreateBuilder().
				SetContent("desired rails are outside boundaries. Minimum: 2. Maximum: 15").
				SetEphemeral(true).
				Build(),
		)
	}
	cypheredString := cyphers.NewRailFenceString(message, rails).Encoded()
	return event.Respond(
		discord.InteractionResponseTypeCreateMessage,
		discord.NewMessageCreateBuilder().
			SetContent(cypheredString).
			SetEphemeral(true).
			Build(),
	)
}
