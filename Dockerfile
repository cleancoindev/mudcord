# sudo docker build -t mudcord .
# sudo docker run -dit -e MUDCORD_TOKEN=foo -e MUDCORD_MONGO_URI=bar --name mudcord mudcord
FROM golang:1.13

WORKDIR /go/src/app
COPY . .

RUN go get github.com/sirupsen/logrus
RUN go get github.com/bwmarrin/discordgo
RUN go install .

CMD ["app"]
