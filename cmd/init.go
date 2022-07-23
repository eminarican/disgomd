package cmd

import (
	"fmt"
	"strings"

	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
)

func Init(cli bot.Client, prefix string) {
	cli.AddEventListeners(bot.NewListenerFunc(func(event *events.MessageCreate) {
		if event.Message.Author.Bot || !strings.HasPrefix(event.Message.Content, prefix) {
			return
		}
		mem := *event.Message.Member
		rmem := discord.ResolvedMember{
			Permissions: cli.Caches().GetMemberPermissions(mem),
			Member:      mem,
		}
		ctx := Context{
			Client: cli,
			Member: &rmem,
		}
		args := strings.Split(event.Message.Content, " ")
		opt := ByAlias(args[0][1:])

		_, _ = event.Client().Rest().CreateMessage(event.ChannelID, func() discord.MessageCreate {
			if opt.IsNone() {
				return discord.NewMessageCreateBuilder().SetContent(fmt.Sprintf("Unknown command: %v. Please check that the command exists and that you have permission to use it.", args[0])).Build()
			}
			res := opt.Unwrap().Execute(strings.Join(args[1:], " "), ctx)
			if res.IsErr() {
				return discord.NewMessageCreateBuilder().SetContent(res.Error().Error()).Build()
			}
			return res.Unwrap()
		}())
	}))
}
