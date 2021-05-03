# Tweeter 3
A mock Twitter embed generator bot for Discord, written in Go.

## Building
1. Install Go
2. Install [discordgo](https://github.com/bwmarrin/discordgo) with `go get github.com/bwmarrin/discordgo`
3. Build the bot with `go build`
4. Set the `$TOKEN_TWEETER` environment variable with your token
	- `export TOKEN_TWEETER=token`
5. Start the bot and pass your token to it
	- `./tweeter`

## Features
- User information added into tweet embed (with support for nicknames)
- You can select other users
- Slash commands
- Randomized likes and retweets
- Accurate recreation of Discord embed
- Written in Go, I guess

## Other commands
- `tweeter -r`
	- Register global slash commands (may take one hour to propagate)
- `tweeter -u`
	- Unregister any previously registered commands

## Example
![Example of usage and embed](https://github.com/cyckl/tweeter/raw/master/img/example.png)

## Usage
There are two available commands:

`/tweet <message>`
This generates a tweet with the provided text.

`/about`
About dialog.

## Bugs?
Report any bugs or other odd behaviors to the Issues page and I'll try to get it all patched up. Thanks for taking a look!
