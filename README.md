# Tweeter 2
A mock Twitter embed generator bot for Discord, now written in Go.

## Building
1. Install Go
2. Install [discordgo](https://github.com/bwmarrin/discordgo) with `go get github.com/bwmarrin/discordgo`
3. Build the bot with `go build`
3. Start the bot and pass your token to it
   - `./tweeter -t <bot token>`

## Features
- User information added into tweet embed (with support for nicknames)
- Randomized likes and retweets
- Accurate recreation of Discord embed
- Written in Go, I guess

## Example
![Example of usage and embed](https://github.com/cyckl/tweeter/raw/master/img/example.png)

## Usage
There are three available commands:

`.t <message>`
This generates a tweet with the provided text.

`t.about`
About dialog.

## Bugs?
Report any bugs or other odd behaviors to the Issues page and I'll try to get it all patched up. Thanks for taking a look!
