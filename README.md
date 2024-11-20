# tgbwp (telegram-bot-with-plugins)

A Golang Telegram bot with plugins support.

## Install Info

-   Create a bot with the [BotFather](https://telegram.me/botfather) command: _`/newbot`_, name, username, and grep the API token from the message.
-   Invite the bot into a group.
-   Run the [BotFather](https://telegram.me/botfather) command: _`/setjoingroups`_ to disable the bot from joining groups.

Now, the bot will only answer group members.

-   Inside the group, type: _`/hello`_.
-   Run `curl https://api.telegram.org/bot<token>/getUpdates | jq` to get and store the chat ID.

-   export TGBWP_TOKEN=<token>`Enable debug:`export debug. DEBUG=true`

Build with: `go` build -buildmode=plugin -o ~/.local/share/tgbwp/plugins/ip.so ./plugins/ip

## Configuration

Example configuration, replace \<number\> with the current chat numbers:

```json
{
  “id”: {
    “user”: {
      “knackwurstking”: <number>
    },
    “chat”: {
      “alice”: <number>
    }
  }
}
```
