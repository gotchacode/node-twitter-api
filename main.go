package main

import (
	"fmt"
	"os"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// User struct
type User struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	Name     string
	Username string
	Email    string
	Provider string
}

// Tweet Struct
type Tweet struct {
	ID   bson.ObjectId `bson:"_id,omitempty"`
	Body string
	User bson.ObjectId `bson:"_id,omitempty"`
}

func connectDB(dailURL string) (*mgo.Session, error) {
	session, err := mgo.Dial(dailURL)
	fmt.Println("Trying to connect to DB")
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	return session, err
}

func getUsers(session *mgo.Session, err error) {

	c := session.DB("ntwitter").C("users")
	var userResult User
	fmt.Println("Trying to query one users")
	err = c.Find(bson.M{"username": "vinitkumar"}).One(&userResult)
	fmt.Println("user query", userResult)
	var results []User
	fmt.Println("Trying to find all users of a certain provider")
	err = c.Find(bson.M{"provider": "github"}).All(&results)
	if err != nil {
		panic(err)
	}
	for _, result := range results {
		fmt.Println("username", result.Username)
	}
	fmt.Println("username", results)
}

func main() {
	var mongodbURL = os.Getenv("MONGODB_URL")
	session, err := connectDB(mongodbURL)
	getUsers(session, err)
}
