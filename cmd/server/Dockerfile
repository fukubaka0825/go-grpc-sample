# Goのバージョンを1.11に指定
ARG GO_VERSION=1.12
# ベースイメージを設定する
FROM golang:${GO_VERSION}-alpine AS builder
# 必要パッケージのインストールを行う
RUN apk add --no-cache git
# 作業ディレクトリを指定する
WORKDIR /
# go.modとgo.sumをコピーする
COPY ./go.mod ./go.sum ./
# 依存go.modとgo.sumをもとに必要パッケージをダウンロード
RUN go mod download
# 作成したソースコードをコピー
COPY . .
# ソースコードからバイナリをビルドする
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/grpc-server cmd/server/main.go || RET=$? || true
# scratchをベースイメージとして指定
FROM scratch
# バイナリをコピー
COPY --from=builder /bin/grpc-server /bin/grpc-server
# 公開用のポートとして8080を指定する
EXPOSE 8080
# バイナリを実行
ENTRYPOINT ["./bin/grpc-server"]
