FROM golang:1.13

WORKDIR /go/src/app
COPY . .

RUN go get github.com/Sirupsen/logrus
RUN go get github.com/bwmarrin/discordgo
RUN go build

CMD ["./mudcord"]
