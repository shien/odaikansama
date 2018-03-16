package main

import (
	"github.com/nlopes/slack"
	"log"
	"math/rand"
	"os"
	"time"
)

const (
	requestMessage = "お題くれ"
)

func run(api *slack.Client) int {
	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for {
		select {
		case msg := <-rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.HelloEvent:
				log.Print("Hello Event")

			case *slack.MessageEvent:
				log.Printf("Message: %v\n", ev)
				if ev.Text == requestMessage {
					theme := ChoiceTheme()
					rtm.SendMessage(rtm.NewOutgoingMessage(theme, ev.Channel))
				}

			case *slack.InvalidAuthEvent:
				log.Print("Invalid credentials")
				return 1

			}
		}
	}
}

func ChoiceTheme() string {
	rand.Seed(time.Now().Unix())

	expression := []string{
		"泣いている",
		"怒っている",
		"笑っている",
		"考えている",
		"喜んでいる",
		"からかっている",
		"企んでいる",
		"焦っている",
		"照れている"}

	theme := expression[rand.Intn(len(expression))]

	hairStyle := []string{"長髪", "ポニーテール", "ツインテール", "短髪", "サイドテール", "ぼさぼさ"}
	theme = theme + hairStyle[rand.Intn(len(hairStyle))] + "の"

	clothing := []string{"警察官", "メイド服／執事服", "ドレス／タキシード", "体操服", "巫女服／神主", "水着", "学生服", "看護士／医者"}
	theme = theme + clothing[rand.Intn(len(clothing))] + "を着た"

	sex := []string{"男", "女"}
	theme = theme + sex[rand.Intn(len(sex))] + "を描きましょう。"

	option := []string{"小動物", "眼鏡", "帽子", "パンスト", "ブーツ", "色黒", "手袋", "眼帯"}
	theme = theme + "おまけで" + option[rand.Intn(len(option))] + "もいれてみては。"

	return theme
}

func main() {
	api := slack.New("APPKEY")
	os.Exit(run(api))
}
