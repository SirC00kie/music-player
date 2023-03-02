FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -o music-player cmd/music-player/main.go

ENTRYPOINT [ "./music-player" ]
