package commands

import (
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
)

var sayCommand = ElizeCommand{
	Command: discord.SlashCommandCreate{
		Name:        "say",
		Description: "I'll send the message you want.",
		Options: []discord.ApplicationCommandOption{
			discord.ApplicationCommandOptionString{
				Name:        "message",
				Description: "What should I say?",
				Required:    true,
			},
			discord.ApplicationCommandOptionBool{
				Name:        "ephemeral",
				Description: "Whether the response should only be visible to you",
				Required:    false,
			},
		},
	},
	Handler: handleSay,
}

func handleSay(event *handler.CommandEvent) error {
	data := event.SlashCommandInteractionData()
	return event.Respond(
		discord.InteractionResponseTypeCreateMessage,
		discord.NewMessageCreateBuilder().
			SetContent(data.String("message")).
			SetEphemeral(data.Bool("ephemeral")).
			Build(),
	)
}
