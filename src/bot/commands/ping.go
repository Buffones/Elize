package commands

import (
	"log/slog"
	"time"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
)

var pingCommand = ElizeCommand{
	Command: discord.SlashCommandCreate{
		Name:        "ping",
		Description: "I'll just reply with Pong!",
	},
	Handler: handlePing,
}

func handlePing(event *handler.CommandEvent) error {
	gatewayPing := event.Client().Gateway().Latency().String()

	eb := discord.NewEmbedBuilder().
		SetTitle("Pong!").
		AddField("Rest", "loading...", false).
		AddField("Gateway", gatewayPing, false).
		SetColor(0x5c5fea)

	defer func() {
		start := time.Now().UnixNano()
		event.Client().Rest().GetBotApplicationInfo()
		duration := time.Now().UnixNano() - start
		eb.SetField(0, "Rest", time.Duration(duration).String(), false)
		if _, err := event.Client().Rest().UpdateInteractionResponse(
			event.ApplicationID(),
			event.Token(),
			discord.MessageUpdate{Embeds: &[]discord.Embed{eb.Build()}},
		); err != nil {
			event.Client().Logger().Error("Failed to update ping embed: ", slog.Any("err", err))
		}
	}()

	return event.Respond(
		discord.InteractionResponseTypeCreateMessage,
		discord.NewMessageCreateBuilder().
			SetEmbeds(eb.Build()).
			SetEphemeral(true).
			Build(),
	)
}
