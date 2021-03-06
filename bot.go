package main

import (
	"github.com/nlopes/slack"
	"log"
	"math/rand"
	"strings"
	"time"
)

const (
	requestMessage         = "お題くれ"
	requestClothingMessage = "服のお題くれ"
	requestAddOdaiMessage  = "お題追加して"
	keepAlive              = "生きてる？"
	answerAlive            = "はい！！！元気です！！！"
	answerAddOdai          = "追加しました"
	answerAddOdaiHelp      = "お題追加して [大項目] [小項目] [お題] と書いてください"
)

type OdaiCache struct {
	Data []Odai
}

type Odai struct {
	OdaiType    string
	OdaiSubtype string
	OdaiList    []string
}

func Run(apikey string) int {
	api := slack.New(apikey)

	rtm := api.NewRTM()
	odai := OdaiCache{}
	go rtm.ManageConnection()

	for {
		select {
		case msg := <-rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.HelloEvent:
				log.Print("Hello Event")

			case *slack.MessageEvent:
				log.Printf("Message: %v\n", ev)
				messageList := strings.Fields(ev.Text)
				if messageList[0] == requestMessage {
					theme := ChoiceTheme(odai)
					rtm.SendMessage(rtm.NewOutgoingMessage(theme, ev.Channel))
				} else if messageList[0] == requestClothingMessage {
					theme := ChoiceClothing(odai)
					rtm.SendMessage(rtm.NewOutgoingMessage(theme, ev.Channel))
				} else if messageList[0] == keepAlive {
					rtm.SendMessage(rtm.NewOutgoingMessage(answerAlive, ev.Channel))
				} else if messageList[0] == requestAddOdaiMessage {
					if len(messageList) == 4 {
						odai.AddOdai(messageList[1], messageList[2], messageList[3])
						rtm.SendMessage(rtm.NewOutgoingMessage(answerAddOdai, ev.Channel))
					} else {
						rtm.SendMessage(rtm.NewOutgoingMessage(answerAddOdaiHelp, ev.Channel))
					}
				}

			case *slack.InvalidAuthEvent:
				log.Print("Invalid credentials")
				return 1

			}
		}
	}
}

func ChoiceClothing(odai OdaiCache) string {

	clothing_bottoms := odai.GetOdai("服", "ボトムス")
	clothing_bottom := clothing_bottoms.OdaiList[rand.Intn(len(clothing_bottoms.OdaiList))]

	clothing_tops := odai.GetOdai("服", "トップス")

	clothing_top1 := clothing_tops.OdaiList[rand.Intn(len(clothing_tops.OdaiList))]
	clothing_top2 := clothing_tops.OdaiList[rand.Intn(len(clothing_tops.OdaiList))]

	shoebox := odai.GetOdai("服", "靴")
	shoes := shoebox.OdaiList[rand.Intn(len(shoebox.OdaiList))]

	option := ChoiceOption(odai)

	theme := clothing_top1 + "か" + clothing_top2 + "か、もしくは両方を合わせた服に" + clothing_bottom + "を描きましょう。履物は" + shoes + "で、" + option + "もおまけでいかがでしょうか。"

	return theme

}

func ChoiceOption(odai OdaiCache) string {
	options := odai.GetOdai("テーマ", "オプション")
	return options.OdaiList[rand.Intn(len(options.OdaiList))]
}

/* テーマを選ぶ */
func ChoiceTheme(odai OdaiCache) string {
	rand.Seed(time.Now().Unix())

	expressions := odai.GetOdai("テーマ", "オプション")
	expression := expressions.OdaiList[rand.Intn(len(expressions.OdaiList))]

	hairStyles := odai.GetOdai("テーマ", "髪型")
	hairStyle := hairStyles.OdaiList[rand.Intn(len(hairStyles.OdaiList))]

	clothing_list := []string{"警察官の服", "メイド服／執事服", "ドレス／タキシード", "体操服", "巫女服／神主服", "水着", "学生服", "看護服／白衣", "私服", "道着"}
	clothing := clothing_list[rand.Intn(len(clothing_list))]

	sex_list := []string{"男", "女"}
	sex := sex_list[rand.Intn(len(sex_list))]

	option := ChoiceOption(odai)

	theme := expression + hairStyle + "の" + clothing + "を着た" + sex + "を描きましょう。おまけで" + option + "もいれてみては。"

	return theme
}
