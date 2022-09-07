# Disgomd

## Todo
- write a readme :P
- add slash commands support

## Example usage
```go
func main() {
	cli, err := disgo.New(os.Getenv("bot_token"),
		bot.WithGatewayConfigOpts(
			gateway.WithIntents(
				gateway.IntentGuildMessages,
				gateway.IntentMessageContent,
			),
		),
	)
	defer cli.Close(context.TODO())
	if err != nil {
		fmt.Println("Couldn't create client")
		os.Exit(1)
	}

	cmd.Init(cli, "!")
	cmd.Register(cmd.New("test", "just a test command", nil, TestCommand{}))

	err = cli.OpenGateway(context.TODO())
	if err != nil {
		fmt.Println("Couldn't open gateway")
		os.Exit(1)
	}

	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-s
}

type TestCommand struct{}

func (TestCommand) Run(ctx cmd.Context) discord.MessageCreate {
	return discord.NewMessageCreateBuilder().
		SetContent("Hi this is a test message :)").Build()
}
```

thanks to [dragonfly](https://github.com/df-mc/dragonfly) contributors for command system
