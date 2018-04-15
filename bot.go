package main

import (
	"github.com/nlopes/slack"
	"log"
	"math/rand"
	"time"
)

const (
	requestMessage         = "お題くれ"
	requestMessageClothing = "服のお題くれ"
	keepAlive              = "生きてる？"
	answerAlive            = "はい！！！元気です！！！"
)

func Run(apikey string) int {
	api := slack.New(apikey)

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
				} else if ev.Text == requestMessageClothing {
					theme := ChoiceClothing()
					rtm.SendMessage(rtm.NewOutgoingMessage(theme, ev.Channel))
				} else if ev.Text == keepAlive {
					rtm.SendMessage(rtm.NewOutgoingMessage(answerAlive, ev.Channel))
				}

			case *slack.InvalidAuthEvent:
				log.Print("Invalid credentials")
				return 1

			}
		}
	}
}

func ChoiceClothing() string {

	clothing_bottoms := []string{
		"ミニスカート",
		"フレアスカート",
		"ギャザースカート",
		"プリーツスカート",
		"ペンシルスカート",
		"コクーンスカート",
		"ペンシルスカート",
		"ペンシルスカート",
		"プリーツスカート",
		"タイトスカート",
		"カプリパンツ",
		"サルエルパンツ",
		"スラックス",
		"スカンツ",
		"デニム",
		"リジットデニム",
		"ベイカーパンツ",
		"ストレートデニム",
		"ボーイフレンド",
		"スキニー",
		"バギーパンツ",
		"テーパードパンツ",
		"パラッツォパンツ",
		"ジョガーパンツ",
		"チノ・パン"}
	clothing_bottom := clothing_bottoms[rand.Intn(len(clothing_bottoms))]

	clothing_tops := []string{
		"ワイシャツ",
		"作業着",
		"法被",
		"セーラー",
		"ワンピース",
		"ニットシャツ",
		"ブラウス",
		"ポロシャツ",
		"デニムシャツ",
		"ニットシャツ",
		"Tシャツ",
		"カットソー",
		"チュニック",
		"タンクトップ",
		"キャミソール",
		"ベアトップ",
		"ニット",
		"ベスト",
		"カーディガン",
		"スウェット",
		"パーカー",
		"ジャージ"}
	clothing_top1 := clothing_tops[rand.Intn(len(clothing_tops))]
	clothing_top2 := clothing_tops[rand.Intn(len(clothing_tops))]

	shoebox := []string{
		"ブーツ",
		"スニーカー",
		"サンダル",
		"パンプス",
		"ショートブーツ",
		"ニーハイブーツ",
		"ハイヒール",
		"厚底靴",
		"下駄",
		"はだし",
		"草履",
		"島草履",
		"雨靴",
		"ローラースケート",
		"ローラーブレード",
		"アイススケート",
		"足袋"}
	shoes := shoebox[rand.Intn(len(shoebox))]

	option := ChoiceOption()

	theme := clothing_top1 + "か" + clothing_top2 + "か、もしくは両方を合わせた服に" + clothing_bottom + "を描きましょう。履物は" + shoes + "で、" + option + "もおまけでいかがでしょうか。"

	return theme

}

func ChoiceOption() string {
	options := []string{
		"小動物",
		"鞄",
		"眼鏡",
		"帽子",
		"パンスト",
		"色黒",
		"手袋",
		"眼帯"}
	return options[rand.Intn(len(options))]
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

	clothing_list := []string{"警察官の服", "メイド服／執事服", "ドレス／タキシード", "体操服", "巫女服／神主服", "水着", "学生服", "看護服／白衣", "私服", "道着"}
	clothing := clothing_list[rand.Intn(len(clothing_list))]

	sex_list := []string{"男", "女"}
	sex := sex_list[rand.Intn(len(sex_list))]

	option := ChoiceOption()

	theme := expression + hairStyle + "の" + clothing + "を着た" + sex + "を描きましょう。おまけで" + option + "もいれてみては。"

	return theme
}
