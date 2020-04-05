# 使用するGolangのイメージを指定する
FROM golang:1.14

# 必要なパッケージなどなどをインストールする
RUN go get -u github.com/go-sql-driver/mysql

# ワーキングディレクトリを指定する
WORKDIR /go/src/sample_docker_compose

# sample_docker_compose直下のディレクトリを全てコンテナ上に載せる
ADD . .

# ファイルをビルド
RUN go build ./docker/golang

# 環境変数を設定
ENV PATH /go/bin:$PATH
