package main

import (
	"encoding/json"
	"flag"
	"os"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

var (
	//BotID is the Discord Bot ID
	BotID string
	//ShowConfig is part of the init process
	ShowConfig string
	//response is the bot response on the channel
	response string
)

//Config structure
type Config struct {
	Token  string `json:"token"`
	Client string `json:"client"`
	Owner  string `json:"owner"`
	Prefix string `json:"prefix"`
	Cool   int    `json:"cooldown"`
	PerC   bool   `json:"per_chan"`
}

func init() {

	verbose := flag.String("v", "info", "set the console verbosity of the bot")
	flag.Parse()
	if *verbose == "debug" {
		log.SetLevel(log.DebugLevel)
		log.Debug("Log level set to debug")
	}
}

func getConfig(a string) string {
	var b string

	//Opens config.json and returns values
	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	config := Config{}
	err := decoder.Decode(&config)

	if err != nil {
		log.Error("error", err)
	}

	switch {
	case a == "token":
		b = config.Token
	case a == "client":
		b = config.Client
	case a == "owner":
		b = config.Owner
	case a == "prefix":
		b = config.Prefix
	default:
		b = "error"
	}

	return b
}

func getCooldown() int {
	var b int

	//Opens config.json and returns values
	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	config := Config{}
	err := decoder.Decode(&config)

	if err != nil {
		log.Error("error", err)
	}

	b = config.Cool

	return b
}

func getChannelStat() bool {
	var b bool

	//Opens config.json and returns values
	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	config := Config{}
	err := decoder.Decode(&config)

	if err != nil {
		log.Error("error", err)
	}

	b = config.PerC

	return b
}

func main() {
	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + getConfig("token"))
	if err != nil {
		log.Fatal("error creating Discord session,", err)
		return
	}

	// Get the account information.
	u, err := dg.User("@me")
	if err != nil {
		log.Fatal("error obtaining account details,", err)
	}

	// Store the account ID for later use.
	BotID = u.ID

	// Register messageCreate as a callback for the messageCreate events.
	dg.AddHandler(messageCreate)

	// Open the websocket and begin listening.
	err = dg.Open()
	if err != nil {
		log.Fatal("error opening connection,", err)
		return
	}

	log.Info("Bot is now running.  Press CTRL-C to exit.")
	// Simple way to keep program running until CTRL-C is pressed.
	<-make(chan struct{})
	return
}
