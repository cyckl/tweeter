// Tweeter bot by cyckl

// Call in dependencies / set up global variables
const { Client, MessageEmbed } = require('discord.js');
const client = new Client();
const version = '1.2.2';
var token = process.argv[2];

// Bootup sequence
client.on('ready', () => {
    console.log('╔════════════════════════════════════╗');
    console.log('║                                    ║');
    console.log('║  Tweeter by cyckl                  ║');
    console.log('║  Running version ' + version + '             ║');
    console.log('║  RIP 3-inch-largest                ║');
    console.log('║                                    ║');
    console.log('╚════════════════════════════════════╝');
    console.log('Online!\n');
    client.user.setActivity('.t msg');
});

// Command init
client.on('message', (message) => {
    if (message.content.startsWith('.t ')) tweet(message);
    if (message.content.startsWith('.run ')) exec(message);
    if (message.content.startsWith('.about')) about(message);
});

// Global Functions
function getRandomInt(x) {
    return Math.floor(Math.random() * x);
}

// Commands
function tweet(message) {
    // Handles removing nickname if no nickname
    if (message.member.nickname == null) tweetUserString = message.author.username;
    else tweetUserString = message.member.nickname + ' (@' + message.author.username + ')';

    var content = message.content.replace('.t ', '');
    var embed = new MessageEmbed()
        .setAuthor(tweetUserString, message.author.displayAvatarURL(), 'https://www.youtube.com/watch?v=dQw4w9WgXcQ')
        .setColor(0x1DA1F2)
        .setDescription(content)
        .addField('Retweets', getRandomInt(50000), true)
        .addField('Likes', getRandomInt(1000000), true)
        .setFooter('Twitter', 'https://abs.twimg.com/icons/apple-touch-icon-192x192.png');
    // Send tweet and log
    message.channel.send(embed);
    console.log('[Tweet] (' + message.author.tag + ') ' + content);
}

function exec(message) {
    // User ID hardcode
    if (message.author.id != '248948821881520138' && message.author.id != '249052912762880000') return;
    
    // Try code, if broken, move to catch()
    try {
        var runCode = message.content.replace('.run ', '');
        var result = eval(runCode);
        // Send result and log
        message.channel.send('**Result:** ```js\n' + result + '```');
        console.log('[Exec] (' + message.author.tag + ') ' + runCode);
        console.log('[Exec Result] ' + result);
    }
    
    // If error, handle it instead of crashing.
    catch (e) {
        console.log('[Exec Error] ' + e);
        message.channel.send('**Error:** ```js\n' + e + '```');
    }
}

function about(message) {
    var embed = new MessageEmbed()
        .setTitle('Tweeter by cyckl')
        .setAuthor('About', 'https://github.com/cyckl/tweeter/raw/master/img/tweeter.png')
        .setColor(0xFF453A)
        .setDescription('Tweeter is a mock Twitter embed generator for Discord.')
        .addField('Version', version, true)
        .addField('Build date', '2020-04-12', true)
        .setFooter('This is alpha software. Please be patient!')
        .setThumbnail('https://github.com/cyckl/tweeter/raw/master/img/tweeter.png');
    // Send tweet and log
    message.channel.send(embed);
    console.log('[About] (' + message.author.tag + ') ' + 'About dialogue triggered.');
}

client.login(token);