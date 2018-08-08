# odaikansama

## how to use

Download slack repository.
```
go get -u github.com/nlopes/slack
```

Build source.
```
go build -o odaikansama
```

Run slack bot.
```
./odaikansama --apikey [apikey]
```

## TODO

- 人間以外も指定できるようにする
- csv にお題が書けるようにする [済]
- お題投稿ができるようにする [済]
 - 同じお題が投稿されてしまう
- 全てのお題を csv に対応させる
- 必要そうならテストを書く
- 時間で毎日一題出せるようにする
