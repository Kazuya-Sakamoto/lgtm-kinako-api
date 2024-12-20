<img src="https://d18g0hf2wnz3gs.cloudfront.net/20240413001121.JPG" alt="LGTM Image" width="830">

# lgtm-kinako-api

<div style="display: flex; gap: 15px; flex-wrap: wrap; align-items: center;">
  <img src="https://img.shields.io/badge/License-MIT-blue" alt="License">
  <img src="https://img.shields.io/badge/GoLang-1.20.7-blue?logo=go&logoColor=white" alt="GoLang">
  <img src="https://img.shields.io/badge/Echo-v4.11.1-green?logo=go&logoColor=white" alt="Echo">
  <img src="https://img.shields.io/badge/MySQL-v8.0-orange?logo=mysql&logoColor=white" alt="MySQL">
  <img src="https://img.shields.io/badge/Gormigrate-v2.1.1-yellow?logo=github&logoColor=white" alt="Gormigrate">
</div>

<br/>

<div style="display: flex; gap: 15px; flex-wrap: wrap; align-items: center;">
  <img src="https://img.shields.io/badge/API-Render-46E3B7?logo=render&logoColor=white" alt="Render">
  <img src="https://img.shields.io/badge/Storage-AWS%20S3-orange?logo=amazonaws&logoColor=white" alt="AWS S3">
  <img src="https://img.shields.io/badge/CDN-AWS%20CloudFront-FF9900?logo=amazonaws&logoColor=white" alt="AWS CloudFront">
  <img src="https://img.shields.io/badge/DB-TiDB-blue?logo=tidb&logoColor=white" alt="TiDB">
  <img src="https://img.shields.io/badge/Watch-UptimeRobot-2ECC71?logo=uptimerobot&logoColor=white" alt="UptimeRobot">
</div>

## Overview

This is a service that allows you to share LGTM images featuring Kinako (my beloved dog). By clicking on an image, the corresponding Markdown code is copied, making it easy to use.

- 🌐 [service URL](https://lgtm-kinako.com/)  
- 💻 [frontend repository](https://github.com/Kazuya-Sakamoto/lgtm-kinako)  

## Environment Setup

- Start the Container

```bash
$ make up
```

- Add Paths

```bash
$ nano ~/.zshrc
```

```bash
# golang
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

- Reload ~/.zshrc

```bash
$ source ~/.zshrc
```

- Remove Unnecessary Dependencies / Add Required Dependencies

```bash
$ go mod tidy
```

- Verify air and dlv Commands

```bash
$ air -v
$ dlv -h
```

- Install air and dlv if Not Already Installed

```bash
$ go install github.com/cosmtrek/air@latest
```

```bash
$ go install github.com/go-delve/delve/cmd/dlv@latest
```

- Run Migrations

```bash
$ make migration
```

- Start the Application

```bash
$ make dev
```

- Verify Functionality

```bash
$ curl http://localhost:8081/api/v1/albums

[{"id":107,"title":"初めてのきなこ","image":"...
```

## Tips

- If the Application Fails to Start

```bash
$ make down
$ make up
```

## Additional Documentation

Access to this documentation is restricted. Please contact @Kazuya-Sakamoto to request access. 🙇‍♂️

- [planetscale でデプロイ方法](https://www.notion.so/planetscale-c49789ce45c741f495a5861312592a21)
- [【Sequel Ace】MySQL GUI クライアントアプリの接続方法](https://www.notion.so/Sequel-Ace-MySQL-GUI-b5f8159e78f043a1beec7d083116da44)
- [.env ファイルについて](https://www.notion.so/env-ad6e94f9e5ef4247ab9e5295bfb00c13)
- [コーディング規約](docs/CODING_GUIDELINES.md)
