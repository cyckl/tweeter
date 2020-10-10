package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	token   string
	version string
)

func init() {
	version = "2.0.2"
	flag.StringVar(&token, "t", "", "Bot token")
	flag.Parse()
}

func main() {
	// Pseudorandom num gen seed

	// Create new Discord session using bot token
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("error creating discord session,", err)
		return
	}

	// Register the tweet, about func as a callback for MessageCreate events
	dg.AddHandler(tweet)
	dg.AddHandler(about)

	// For this example we only care about recieving message events
	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)

	// Open websocket conn to Discord and listen
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until EOF
	fmt.Println("╔═══════════════════════╗")
	fmt.Println("║ Tweeter by cyckl      ║")
	fmt.Println(fmt.Sprintf("║ Running version %s ║", version))
	fmt.Println("╚═══════════════════════╝")
	rand.Seed(time.Now().UnixNano())
	fmt.Println("[Info] Random number generator seeded")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Close Discord session cleanly
	dg.Close()
	fmt.Println("[Info] Program terminated")
}

// Pueudorandom numbers
func randInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func tweet(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore bot messages, good practice
	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.HasPrefix(m.Content, ".t ") {
		// Nickname handling
		tweetAuthor := "undefined"
		if m.Member.Nick != "" {
			tweetAuthorValues := []string{m.Member.Nick, " (@", m.Author.Username, ")"}
			tweetAuthor = strings.Join(tweetAuthorValues, "")
		} else {
			tweetAuthor = m.Author.Username
		}

		// Fill embed and send it
		s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			Author: &discordgo.MessageEmbedAuthor{
				Name:		tweetAuthor,
				IconURL:	discordgo.EndpointUserAvatar(m.Author.ID, m.Author.Avatar),
			},
			Color:			1942002,
			Description:	strings.TrimPrefix(m.Content, ".t "),
			Footer:	&discordgo.MessageEmbedFooter{
				Text:		"Twitter",
				IconURL:	"https://abs.twimg.com/icons/apple-touch-icon-192x192.png",
			},

			Fields: []*discordgo.MessageEmbedField{
				{Name: "Retweets", Value: strconv.Itoa(randInt(5000, 50000)), Inline: true},
				{Name: "Likes", Value: strconv.Itoa(randInt(25000, 150000)), Inline: true},
			},
		})
		fmt.Println(fmt.Sprintf("[Tweet] (%s) %s", m.Author.ID, strings.TrimPrefix(m.Content, ".t ")))
	}
}

func about(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore bot messages, good practice
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "t.about" {
		// Send embed
		s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
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
				Text:    "This is beta software. Please be patient.",
			},

			Fields: []*discordgo.MessageEmbedField{
				{Name: "Version", Value: version, Inline: true},
				{Name: "Build date", Value: "2020-10-09", Inline: true},
				{Name: "Github", Value: "https://github.com/cyckl/tweeter", Inline: false},
			},
		})
	}
}
