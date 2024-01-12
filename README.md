# lgtm-kinako-api

## 概要

きなこ（愛犬）の LGTM 画像を共有出来るサービスです。画像をクリックすると Markdown がコピーされ使用することができます。

- サービス URL
  - https://lgtm-kinako.com/
- フロントリポジトリ
  - https://github.com/Kazuya-Sakamoto/lgtm-kinako

## 環境

- golang 1.20.7
- echo v4.11.1
- mysql v8.0
- gormigrate v2.1.1

## その他環境

- API: Render
- Storage: AWS s3
- CDN: AWS cloudfront
- DB: PlanetScale
- Watch: UptimeRobot

## 環境構築

- コンテナ起動

```
$ docker compose up -d
```

- path の追加

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

- 不要な依存関係の削除 / 必要な依存関係の追加

```
$ go mod tidy
```

- air と dlv のコマンドが実行できるか確認

```
$ air -v
$ dlv -h
```

- air と dlv をインストールしていない場合は以下でインストール

```
$ go install github.com/cosmtrek/air@latest
```

```
$ go install github.com/go-delve/delve/cmd/dlv@latest
```

- マイグレーション

```
$ GO_ENV=dev go run migrate/migrate.go
```

- 起動

```
$ GO_ENV=dev air
```

- 動作確認

```
$ curl http://localhost:8080/album/random

[{"id":107,"title":"初めてのきなこ","image":"...
```
