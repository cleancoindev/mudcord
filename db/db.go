package db

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/tteeoo/mudcord/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

var ctx context.Context
var Cancel context.CancelFunc
var users *mongo.Collection
var servers *mongo.Collection

func init() {
	ctx, Cancel = context.WithCancel(context.Background())

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MUDCORD_MONGO_URI")))
	util.CheckFatal(err)

	logrus.Info("connected to db")

	users = client.Database("mudcord").Collection("users")
	servers = client.Database("mudcord").Collection("servers")
}

type itemQuan struct {
	ID   string
	Quan int
}

// NewServer creates a new server in our db
func NewServer(id string) (*Server, error) {
	server := Server{
		ID:     id,
		Prefix: ".",
	}
	_, err := servers.InsertOne(ctx, server)

	if err != nil {
		return &Server{}, err
	}

	return &server, nil
}

// GetServer gets a server from db based on discord id
func GetServer(id string) (*Server, error) {
	data := servers.FindOne(ctx, bson.M{"id": id})

	var server Server
	err := data.Decode(&server)

	if err != nil {
		return &Server{}, err
	}

	return &server, nil
}

// SetServer changes server data in the db based on discord id
func SetServer(newServer *Server) error {
	data := servers.FindOneAndUpdate(ctx, bson.M{"id": newServer.ID}, bson.M{"$set": newServer})

	var server Server
	err := data.Decode(&server)

	if err != nil {
		return err
	}

	return nil
}

// CheckServer checks if a server with the given id exists
func CheckServer(id string) bool {
	_, err := GetServer(id)

	if err == mongo.ErrNoDocuments {
		return false
	}

	return true
}

// NewUser creates a new user in our db
func NewUser(id string) (*User, error) {
	user := User{
		ID:      id,
		Level:   1,
		XP:      0,
		HP:      [2]int{20, 20},
		MP:      [2]int{20, 20},
		Gold:    0,
		Room:    "RoomGreatMarya",
		Hat:     "HatNone",
		Inv:     []*itemQuan{{ID: "ConsumableCanteen", Quan: 1}},
		Arsenal: []string{"WeaponBaseballBat"},
	}
	_, err := users.InsertOne(ctx, user)

	if err != nil {
		return &User{}, err
	}

	return &user, nil
}

// GetUser gets a user from db based on discord id
func GetUser(id string) (*User, error) {
	data := users.FindOne(ctx, bson.M{"id": id})

	var user User
	err := data.Decode(&user)

	if err != nil {
		return &User{}, err
	}

	return &user, nil
}

// SetUser changes user data in the db based on discord id
func SetUser(newUser *User) error {
	data := users.FindOneAndUpdate(ctx, bson.M{"id": newUser.ID}, bson.M{"$set": newUser})

	var user User
	err := data.Decode(&user)

	if err != nil {
		return err
	}

	return nil
}

// CheckStarted checks if a user with the given user id exists
func CheckStarted(id string) bool {
	_, err := GetUser(id)

	if err == mongo.ErrNoDocuments {
		return false
	}

	return true
}
