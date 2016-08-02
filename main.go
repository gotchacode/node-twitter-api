package main

import (
	"fmt"
	"os"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	Name     string
	Username string
	provider string
}

func main() {
	var mongodbURL = os.Getenv("MONGODB_URL")
	session, err := mgo.Dial(mongodbURL)
	if err != nil {
		panic(err)
	}
	//session.SetMode(mgo.Monotonic, true)
	var results []User
	c := session.DB("ntwitter").C("users")
	err = c.Find(bson.M{"provider": "github"}).All(&results)
	if err != nil {
		panic(err)
	}
	// fmt.Println("username", result.Username)
	for _, result := range results {
		fmt.Println("username", result.Username)
	}
	fmt.Println("username", results)

}
