# Tweeter
A mock Twitter embed generator bot for Discord.

## Getting started
1. Install Node.JS and npm
2. Install Discord.JS with `npm install discord.js`
3. Create a script with this format
   - `node index.js <bot token>`

## Features
- User information pasted into tweet embed (with support for nicknames)
- Randomized likes and retweets (in very high quantities!)
- Accurate recreation of Discord embed
- All commands and usage logged to console output.
- Lightweight (only 83 lines!)

## Example
![Example of usage and embed](https://github.com/cyckl/tweeter/raw/master/img/example.png)

## Usage
There are three available commands:

`.t <message>`
This generates a tweet with the provided argument.

`.about`
About page.

`.run <arg>`
Executes JS from the provided arg. This is designed to be hardcoded to one or more Discord user ID's for security.

## Bugs?
Report any bugs or other odd behaviors to the Issues page and I'll try to get it all patched up. Thanks for taking a look!