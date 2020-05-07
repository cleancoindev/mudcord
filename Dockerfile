# sudo docker build -t mudcord .
# sudo docker run -dit -e MUDCORD_TOKEN=foo -v /foo:/go/src/app --name mudcord mudcord
FROM golang:1.13

WORKDIR /go/src/app
COPY . .

RUN go get github.com/sirupsen/logrus
RUN go get github.com/bwmarrin/discordgo
RUN go install .

CMD ["app"]
