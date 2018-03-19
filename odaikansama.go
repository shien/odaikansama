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

	expressions := []string{
		"泣いている",
		"怒っている",
		"笑っている",
		"考えている",
		"喜んでいる",
		"からかっている",
		"企んでいる",
		"焦っている",
		"照れている"}
	expression := expressions[rand.Intn(len(expressions))]

	hairStyles := []string{"長髪", "ポニーテール", "ツインテール", "短髪", "サイドテール", "ぼさぼさ"}
	hairStyle := hairStyles[rand.Intn(len(hairStyles))]

	clothing_list := []string{"警察官の服", "メイド服／執事服", "ドレス／タキシード", "体操服", "巫女服／神主", "水着", "学生服", "看護士／医者", "私服"}
	clothing := clothing_list[rand.Intn(len(clothing_list))]

	sex_list := []string{"男", "女"}
	sex := sex_list[rand.Intn(len(sex_list))]

	options := []string{"小動物", "眼鏡", "帽子", "パンスト", "ブーツ", "色黒", "手袋", "眼帯"}
	option := options[rand.Intn(len(options))]

	theme := expression + hairStyle + "の" + clothing + "を着た" + sex + "を描きましょう。おまけで" + option + "もいれてみては。"

	return theme
}

func main() {
	api := slack.New("APIKEY")
	os.Exit(run(api))
}
