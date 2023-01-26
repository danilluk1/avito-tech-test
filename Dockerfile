FROM alpine as base

RUN apk add git curl wget upx

WORKDIR /app

COPY --from=golang:1.19.5-alpine /usr/local/go /usr/local/go
ENV PATH="$PATH:/usr/local/go/bin"
ENV PATH="$PATH:/root/go/bin"

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -ldflags="-s -w" -o ./api ./cmd/main.go && upx -9 -k ./api

FROM base as api
COPY --from=base /app/api /bin/api
CMD ["/bin/api"]