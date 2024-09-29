package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/gateway"
	"github.com/disgoorg/disgo/handler"

	ecommands "buffones/elize/commands"
)

var token = os.Getenv("TOKEN")

func main() {
	cr := handler.New()
	for _, c := range ecommands.Commands {
		cr.Command("/"+c.Command.Name, c.Handler)
	}
	client, err := disgo.New(
		token,
		bot.WithGatewayConfigOpts(
			gateway.WithIntents(
				gateway.IntentGuilds,
				gateway.IntentGuildMessages,
				gateway.IntentDirectMessages,
			),
		),
		bot.WithEventListeners(cr),
	)
	if err != nil {
		panic(err)
	}
	defer client.Close(context.TODO())

	if _, err = client.Rest().SetGlobalCommands(client.ApplicationID(), ecommands.GetCommandCreates()); err != nil {
		slog.Error("error while registering commands", slog.Any("err", err))
	}

	if err = client.OpenGateway(context.TODO()); err != nil {
		panic(err)
	}

	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-s
}
