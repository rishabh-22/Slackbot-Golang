package main

import (
	"log"
	"os"

	"github.com/Krognol/go-wolfram"
	"github.com/christianrondeau/go-wit"
	"github.com/nlopes/slack"
)

const confidenceThreshold = 0.5

var (
	slackClient   *slack.Client
	witClient     *wit.Client
	wolframClient *wolfram.Client
)

func main() {
	slackClient = slack.New(os.Getenv("SLACK_ACCESS_TOKEN"))
	witClient = wit.NewClient(os.Getenv("WIT_AI_ACCESS_TOKEN"))
	wolframClient = &wolfram.Client{AppID: os.Getenv("WOLFRAM_APP_ID")}

	rtm := slackClient.NewRTM()
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.MessageEvent:
			if len(ev.BotID) == 0 {
				go handleMessage(rtm, ev)
			}
		}
	}
}

func handleMessage(rtm *slack.RTM, ev *slack.MessageEvent) {
	result, err := witClient.Message(ev.Msg.Text)
	if err != nil {
		log.Printf("unable to get wit.ai result: %v", err)
		return
	}

	var (
		topEntity    wit.MessageEntity
	)

	for _, entityList := range result.Entities {
		for _, entity := range entityList {
			if entity.Confidence > confidenceThreshold && entity.Confidence > topEntity.Confidence {
				topEntity = entity

			}
		}
	}

	replyToUser(rtm, ev, topEntity)
}

func replyToUser(rtm *slack.RTM, ev *slack.MessageEvent, topEntity wit.MessageEntity) {
	switch topEntity.Value {
	case "greeting":
		rtm.SendMessage(rtm.NewOutgoingMessage("Yo what up!?", ev.Channel))
		return
	case "wolfram_search_query":
		res, err := wolframClient.GetSpokentAnswerQuery(ev.Text, wolfram.Metric, 1000)
		if err == nil {
			rtm.SendMessage(rtm.NewOutgoingMessage(res, ev.Channel))
			return
		}

		log.Printf("unable to get data from wolfram: %v", err)
	}
	rtm.SendMessage(rtm.NewOutgoingMessage("¯\\_(o_o)_/¯", ev.Channel))
}