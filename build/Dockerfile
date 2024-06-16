FROM golang:1.18.1-alpine

# アップデートとgitのインストール
RUN apk update && \
    apk --no-cache add git

# appディレクトリの作成
RUN mkdir /go/src/app

# ワーキングディレクトリの設定
WORKDIR /go/src/app/

# ホストのファイルをコンテナの作業ディレクトリに移行
ADD . /go/src/app

# Goモジュールの初期化と依存関係の解決
RUN go mod tidy

# OpenAIのGoライブラリをインストール
RUN go get github.com/sashabaranov/go-openai

# アプリケーションの実行
CMD ["go", "run", "./app/main.go"]

# RUN go install github.com/cosmtrek/air@v1.29.0
# CMD ["air", "-c", ".air.toml"]