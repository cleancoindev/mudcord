package db

import (
	"context"
	"github.com/tteeoo/mudcord/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

var ctx context.Context
var cancel context.CancelFunc
var users *mongo.Collection

func init() {
	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MUDCORD_MONGO_URI")))
	util.CheckFatal(err)

	users = client.Database("mudcordDEV").Collection("users")
}

type itemQuan struct {
	ID   string
	Quan int
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

// GetUser gets a use from db based on discord id
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
