# goのバージョン
FROM golang:1.22.5-alpine3.20

# OSのインストール済みのパッケージをバージョンアップ&必要なパッケージをインストール
RUN apk upgrade --update && \
    apk --no-cache add git gcc musl-dev

# ワーキングディレクトリの設定
WORKDIR /go/src/app/

# ホストPCの[./src]フォルダーを[/go/src/app/]フォルダーにコピー
ADD ./mr-platform-recommend-contents-server/src .

CMD ["go", "run", "main.go"]