# lgtm-kinako-api

## 環境構築

- pathの追加

```
$ nano ~/.zshrc
```

```
# golang
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

- ~/.zshrc を読み込み直す
```
$ source ~/.zshrc
```

- air と dlv のコマンドが実行できるか確認
```
$ air -v
$ dlv -h
```
- air と dlvをインストールしていない場合は以下でインストール

```
$ go install github.com/cosmtrek/air@latest
```

```
$ go install github.com/go-delve/delve/cmd/dlv@latest
```


- 起動
```
$ GO_ENV=dev air
```

- 動作確認

```
$ curl http://localhost:8080/album

[{"id":107,"title":"初めてのきなこ","image":"...
```
