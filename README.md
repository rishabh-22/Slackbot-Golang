# slackbot-golang
This repository holds the code for a slackbot made using golang which replies to the queries made by the user using wit.ai at the backend.

### Install the dependencies:

```
go get github.com/Krognol/go-wolfram
go get github.com/christianrondeau/go-wit
go get github.com/nlopes/slack
```

make a file with your access tokens in it and source the file

```
SLACK_ACCESS_TOKEN=<token>
WIT_AI_ACCESS_TOKEN=<token>
WOLFRAM_APP_ID=<token>
```

run the file:

`go run main.go`

## Prerequisites:
1. Make a bot in your slack workspace
2. Make an app on wit.ai and train it for basic questions. (I've used just greetings apart from general questions here)
make intents as greeting and wolfram_search_query and train the latter intent with basic questions.
3. Make an app on wolfram developer portal.
