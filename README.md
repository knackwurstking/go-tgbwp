# tgbwp (telegram-bot-with-plugins)

A in golang written telegram bot with plugins support
(using the golang plugins package).
I use this to get my (raspberry pi) server IP.

## Install Info (NOTES)

- first create a bot with the **[BotFather](https://telegram.me/botfather)**
  - telegram command: _`/newbot`_
  - type in the bots _name_
  - type in the bots _username_
  - grep the api token from the message
- now create a group and invite your bot into this group
- run telegram command (**BotFather**): _`/setjoingroups`_ to Disable

Now we need to do some stuff so that the bot only answer the group members

- inside the telegram group type: _`/hello`_
- then run `curl https://api.telegram.org/bot<token>/getUpdates | jq`
- get and store the _**chat id**_ from the curl command

- env: `export TGBWP_TOKEN=<token>`

- build with: `go build -buildmode=plugin -o ~/.local/share/tgbwp/plugins/ip.so ./plugins/ip`

## Configuration

Example configuration, replace \<number\> with the current chat numbers,
get it with `curl https://api.telegram.org/bot<token>/getUpdates | jq`

```json
{
  "id": {
    "user": {
      "knackwurstking": <number>
    },
    "chat": {
      "alice": <number>
    }
  }
}
```

## TODOs

- [ ] Adding new plugins for (rpi) getting journalctl logs (as files)
