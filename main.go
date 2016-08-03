package main

import (
	"fmt"
	"log"
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
	User bson.ObjectId
}

func connectDB(dailURL string) (*mgo.Session, error) {
	mgo.SetDebug(true)

	var aLogger *log.Logger
	aLogger = log.New(os.Stderr, "", log.LstdFlags)
	mgo.SetLogger(aLogger)

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
	fmt.Print("userID", userResult.ID)

	var results []User
	fmt.Println("Trying to find all users of a certain provider")
	err = c.Find(bson.M{"provider": "github"}).All(&results)
	if err != nil {
		panic(err)
	}
	//for _, result := range results {
	//	fmt.Println("username", result.Username, result.Provider)
	//}
	// fmt.Println("username", results)
}

func getTweets(session *mgo.Session, err error) {
	connectionUser := session.DB("ntwitter").C("users")
	var userResult User
	err = connectionUser.Find(bson.M{"username": "vinitkumar"}).One(&userResult)
	fmt.Println("Finding user for a tweet query")
	fmt.Print("userID", userResult.ID)

	connectionTweet := session.DB("ntwitter").C("tweets")
	var tweet Tweet
	err = connectionTweet.Find(bson.M{"user": userResult.ID}).One(&tweet)
	fmt.Println("tweet", tweet)

	err = connectionTweet.Find(bson.M{"Body": "Kya chutiyaapa hai"}).One(&tweet)
	fmt.Println("tweet", tweet)
	var tweets []Tweet
	fmt.Println("sample request", bson.M{"user": userResult.ID})
	err = connectionTweet.Find(bson.M{"user": userResult.ID}).All(&tweets)
	fmt.Println("tweets", tweets)
}

func main() {
	var mongodbURL = os.Getenv("MONGODB_URL")
	session, err := connectDB(mongodbURL)
	getUsers(session, err)
	getTweets(session, err)
}
