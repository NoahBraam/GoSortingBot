package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"GoSortingBot/commands"

	"github.com/bwmarrin/discordgo"
)

// Parameters from flag.
var accountToken string

func init() {
	// Parse command line arguments.
	flag.StringVar(&accountToken, "t", "", "Bot account token")
	flag.Parse()
	if accountToken == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
}

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	var err error
	var session *discordgo.Session
	session, err = discordgo.New("Bot " + accountToken)
	setupHandlers(session)
	panicOnErr(err)
	err = session.Open()
	panicOnErr(err)
	<-make(chan struct{})
}

func setupHandlers(sess *discordgo.Session) {
	fmt.Printf("Setup.")
	sess.AddHandler(messageHandler)
}

func messageHandler(sess *discordgo.Session, evt *discordgo.MessageCreate) {
	message := evt.Message
	switch strings.ToLower(strings.TrimSpace(message.Content)) {
	case "!help":
		commands.HelpCommand(sess, message)
	case "!houses":
		commands.HouseCommand(sess, message)
	case "!gryffindor":
		commands.GryffindorCommand(sess, message)
	case "!ravenclaw":
		commands.RavenclawCommand(sess, message)
	case "!hufflepuff":
		commands.HufflepuffCommand(sess, message)
	case "!slytherin":
		commands.SlytherinCommand(sess, message)
	}
}
