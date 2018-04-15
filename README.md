# odaikansama

## how to use

Download slack repository.
```
go get -u github.com/nlopes/slack
```

Build source.
```
go build -o odaikansama.go
```

Run slack bot.
```
./odaikansama --apikey [apikey]
```

## TODO

- お題の種類を増やす
 - 服
 - デフォルトの適当なお題
- 人間以外も指定できるようにする
- conf にお題が書けるようにする
- 必要そうならテストを書く
- 時間で毎日一題出せるようにする
