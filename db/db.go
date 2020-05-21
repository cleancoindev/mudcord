package db

import (
    "go.mongodb.org/mongo-driver/mongo"
    "github.com/tteeoo/mudcord/util"
    "go.mongodb.org/mongo-driver/mongo/options"
    "context"
    "time"
)

func init() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
	    "mongodb+srv://theo:<password>@cluster0-nosjr.mongodb.net/test?retryWrites=true&w=majority",
	))
	util.CheckFatal(err)

	users := client.Database("mudcord").Collection("users")
	servers := client.Database("mudcord").Collection("servers")
}

