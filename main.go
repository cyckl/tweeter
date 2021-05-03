package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
	"github.com/bwmarrin/discordgo"
)

// Set global variables / boot information
const version = "3.0.0"
var (
	s				*discordgo.Session
	token =			flag.String("t", "", "Bot token")
	register =		flag.Bool("r", false, "Register (or re-register) bot commands on startup")
	unregister =	flag.Bool("u", false, "Unregister bot commands on shutdown")
)

func init() {
	// Print boot "splash"
	fmt.Println("╔═══════════════════════╗")
	fmt.Println("║ Tweeter by cyckl      ║")
	fmt.Println(fmt.Sprintf("║ Running version %s ║", version))
	fmt.Println("╚═══════════════════════╝")
	log.Println("[Info] Minimum permissions are 75776")
	
	// Seed RNG
	rand.Seed(time.Now().UnixNano())
	log.Println("[Info] Random number generator seeded")
	
	// Pass args in
	flag.Parse()
}

// Establish available commands
var (
	commands = []*discordgo.ApplicationCommand{
		{
			Name:			"tweet",
			Description:	"Generate a fake tweet",
			Options:		[]*discordgo.ApplicationCommandOption{
				{
					Type:			discordgo.ApplicationCommandOptionString,
					Name:			"content",
					Description:	"Tweet content",
					Required:		true,
				},
				{
					Type:			discordgo.ApplicationCommandOptionUser,
					Name:			"user",
					Description:	"Select tweet author",
					Required:		false,
				},
			},
		},
		{
			Name:			"about",
			Description:	"About Tweeter",
		},
	}
)

func main() {
	// Declare error here so it can be set without :=
	var err error
	
	// Create bot client session
	log.Println("[Info] Logging in")
	s, err = discordgo.New("Bot " + *token)
	if err != nil {
		log.Fatalf("[Fatal] Error creating session: %v", err)
	}
	
	// Pass on command events to functions
	log.Println("[Info] Registering intents")
	s.AddHandler(tweet)
	s.AddHandler(about)

	// We only care about integration (command) intents
	s.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildIntegrations)

	// Open websocket connection to Discord and listen
	err = s.Open()
	if err != nil {
		log.Fatalf("[Fatal] Error opening connection: %v", err)
	}
	
	if *register == true {
		registerCmd(s)
	}

	// Close Discord session cleanly
	defer s.Close()
	
	stop := make(chan os.Signal)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-stop
	
	if *unregister == true {
		unregisterCmd(s)
	}
	
	log.Println("[Info] Shutting down")
}

// Pueudorandom numbers
func randInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func tweet(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Data.Name == "tweet" {
		// Set default endpoints for data (no selected user and no nickname)
		author := fmt.Sprintf("@%s", i.Member.User.Username)
		avatar := discordgo.EndpointUserAvatar(i.Member.User.ID, i.Member.User.Avatar)
		
		// Handle user selection and nicknames
		if len(i.Data.Options) == 2 {
			// If there is a selected user, set the author and avatar to the selected user instead of default
			author = fmt.Sprintf("@%s", i.Data.Options[1].UserValue(s).Username)
			avatar = discordgo.EndpointUserAvatar(i.Data.Options[1].UserValue(s).ID, i.Data.Options[1].UserValue(s).Avatar)
		} else {
			// If there is no selected user, at least check if the command author has a nickname
			if i.Member.Nick != "" {
				author = fmt.Sprintf("%s (@%s)", i.Member.Nick, i.Member.User.Username)
			}
		}
	
		// Respond to command event with embed
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type:	discordgo.InteractionResponseChannelMessageWithSource,
			Data:	&discordgo.InteractionApplicationCommandResponseData{
				// One entry in the array of embeds... Why did they make this a fucking array?
				Embeds: []*discordgo.MessageEmbed{
					{
						Author: &discordgo.MessageEmbedAuthor{
							Name:		author,
							IconURL:	avatar,
						},
						Color:			1942002,
						Description:	i.Data.Options[0].StringValue(),
						Footer:	&discordgo.MessageEmbedFooter{
							Text:		"Twitter",
							IconURL:	"https://abs.twimg.com/icons/apple-touch-icon-192x192.png",
						},
						Fields: []*discordgo.MessageEmbedField{
							{Name: "Retweets", Value: strconv.Itoa(randInt(5000, 50000)), Inline: true},
							{Name: "Likes", Value: strconv.Itoa(randInt(25000, 150000)), Inline: true},
						},
					},
				},
			},
		})
	}
}

func about(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Data.Name == "about" {
		// Respond to command event with embed
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type:	discordgo.InteractionResponseChannelMessageWithSource,
			Data:	&discordgo.InteractionApplicationCommandResponseData{
				Embeds: []*discordgo.MessageEmbed{
					{
						Thumbnail: &discordgo.MessageEmbedThumbnail{
							URL:	"https://github.com/cyckl/tweeter/raw/master/img/tweeter.png",
						},
						Title: "Tweeter by cyckl",
						Author: &discordgo.MessageEmbedAuthor{
							Name:    "About",
							IconURL: "https://github.com/cyckl/tweeter/raw/master/img/tweeter.png",
						},
						Color:       16729402,
						Description: "Tweeter is a mock Twitter embed generator for Discord.",
						Footer: &discordgo.MessageEmbedFooter{
							Text:    "Why do I still maintain this?",
						},

						Fields: []*discordgo.MessageEmbedField{
							{Name:	"Version",		Value: version,								Inline: true},
							{Name:	"Build date",	Value: "2021-05-02",						Inline: true},
							{Name:	"Github",		Value: "https://github.com/cyckl/tweeter",	Inline: false},
						},
					},
				},
			},
		})
	}
}

func registerCmd(s *discordgo.Session) {
	// Register commands
	for _, v := range commands {
		log.Println("[Info] Registering:", v.Name)
		_, err := s.ApplicationCommandCreate(s.State.User.ID, "", v)
		if err != nil {
			log.Panicf("[Error] Cannot create '%v' command: %v", v.Name, err)
		}
	}
}

func unregisterCmd(s *discordgo.Session) {
	// Get command list
	pubCommands, err := s.ApplicationCommands(s.State.User.ID, "")
	if err != nil {
		log.Panicf("[Error] Cannot get application commands: %v", err)
	}
	
	// Unregister commands by ID
	for _, v := range pubCommands {
		log.Println("[Info] Unregistering:", v.Name)
		err := s.ApplicationCommandDelete(s.State.User.ID, "", v.ID)
		if err != nil {
			log.Panicf("[Error] Cannot delete '%v' command: %v", v.Name, err)
		}
	}
	log.Println("[Info] Commands unregistered")
}
