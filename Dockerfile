# sudo docker build -t mudcord .
# sudo docker run -dit -e MUDCORD_TOKEN=foo -v /foo:/mudcord --name mudcord mudcord
FROM golang:1.13

WORKDIR /go/src/app
COPY . .

RUN go get github.com/sirupsen/logrus
RUN go get github.com/bwmarrin/discordgo
RUN go install .

RUN mkdir /mudcord
WORKDIR /mudcord

CMD ["app"]
