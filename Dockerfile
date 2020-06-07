# sudo docker build -t mudcord .
# sudo docker run -dit -e MUDCORD_TOKEN=foo -e MUDCORD_MONGO_URI=bar -v /var/log/mudcord:/go/src/app/log --name mudcord mudcord
FROM golang:1.13

WORKDIR /go/src/app
COPY . .

RUN go get github.com/bwmarrin/discordgo
RUN go install .

CMD ["mudcord"]
