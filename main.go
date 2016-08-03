// Goals:
// The goals for this API is following.
// - Provide API endpints for tweets, users, analytics, comments
// - Small isolated modular functions that provide an easy interface to build
// upon.

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

type Analytics struct {
	ID   bson.ObjectId `bson:"_id,omitempty"`
	IP   string
	User bson.ObjectId
	Url  string
}

func connectDB(dailURL string, debug bool) (*mgo.Session, error) {
	if debug {
		mgo.SetDebug(debug)

		var aLogger *log.Logger
		aLogger = log.New(os.Stderr, "", log.LstdFlags)
		mgo.SetLogger(aLogger)
	}
	session, err := mgo.Dial(dailURL)
	fmt.Println("Trying to connect to DB")
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	return session, err
}

func getAnalytics(session *mgo.Session, err error) {
	AnalyticConnection := session.DB("ntwitter").C("analytics")
	var analytics Analytics
	err = AnalyticConnection.Find(nil).One(&analytics)
	fmt.Println("analytics", analytics)
	var analyticsList []Analytics
	err = AnalyticConnection.Find(nil).All(&analyticsList)
	for _, analytic := range analyticsList {
		fmt.Println("analytic", analytic.User, analytic.IP, analytic.Url)
	}
}

func getUsers(session *mgo.Session, err error) {
	userConnection := session.DB("ntwitter").C("users")
	var userResult User
	fmt.Println("Trying to query one users")
	err = userConnection.Find(bson.M{"username": "vinitkumar"}).One(&userResult)
	fmt.Println("user query", userResult)
	fmt.Print("userID", userResult.ID)

	var results []User
	fmt.Println("Trying to find all users of a certain provider")
	err = userConnection.Find(bson.M{"provider": "github"}).All(&results)
	if err != nil {
		panic(err)
	}
	//for _, result := range results {
	//	fmt.Println("username", result.Username, result.Provider)
	//}
	// fmt.Println("username", results)
}

func getTweets(session *mgo.Session, err error) {
	UserConnection := session.DB("ntwitter").C("users")
	var userResult User
	err = UserConnection.Find(bson.M{"username": "vinitkumar"}).One(&userResult)
	fmt.Println("Finding user for a tweet query")
	fmt.Print("userID", userResult.ID)

	TweetConnection := session.DB("ntwitter").C("tweets")
	var tweet Tweet
	err = TweetConnection.Find(bson.M{"user": userResult.ID}).One(&tweet)
	fmt.Println("tweet", tweet)

	err = TweetConnection.Find(bson.M{"Body": "Kya chutiyaapa hai"}).One(&tweet)
	fmt.Println("tweet", tweet)
	var tweets []Tweet
	fmt.Println("sample request", bson.M{"user": userResult.ID})
	err = TweetConnection.Find(bson.M{"user": userResult.ID}).All(&tweets)
	fmt.Println("tweets", tweets)
}

func main() {
	var mongodbURL = os.Getenv("MONGODB_URL")
	session, err := connectDB(mongodbURL, false)
	//getUsers(session, err)
	//getTweets(session, err)
	getAnalytics(session, err)
}
