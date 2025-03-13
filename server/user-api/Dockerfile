FROM golang:1.24-alpine AS builder
WORKDIR /server/user-api
# Goの依存関係をダウンロード
COPY go.mod go.sum ./
RUN go mod download

# ソースコードをコピーしてビルド
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o user-api ./cmd/api/

# 最小限のランタイムイメージ（Distroless）を利用
FROM gcr.io/distroless/base-debian10
COPY --from=builder /server/user-api /user-api
EXPOSE 8080
ENV PORT=8080
CMD ["/user-api"]
