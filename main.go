// Goals:
// The goals for this API is following.
// - Provide API endpints for tweets, users, analytics, comments
// - Small isolated modular functions that provide an easy interface to build
// upon.

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Query Map
type Query map[string]string

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

// Analytic struct
type Analytic struct {
	ID   bson.ObjectId `bson:"_id,omitempty"`
	IP   string
	User bson.ObjectId
	URL  string
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

func getAnalytic(session *mgo.Session, err error) Analytic {
	AnalyticConnection := session.DB("ntwitter").C("analytics")
	var analytic Analytic
	err = AnalyticConnection.Find(nil).One(&analytic)
	return analytic
}

func getAnalytics(session *mgo.Session, err error) []Analytic {
	AnalyticConnection := session.DB("ntwitter").C("analytics")
	var analyticsList []Analytic
	err = AnalyticConnection.Find(nil).All(&analyticsList)
	return analyticsList
}

func getUser(session *mgo.Session, err error) User {
	userConnection := session.DB("ntwitter").C("users")
	var user User
	fmt.Println("Trying to query one users")
	err = userConnection.Find(bson.M{"username": "vinitkumar"}).One(&user)
	return user
}

func getUsers(session *mgo.Session, err error) []User {
	userConnection := session.DB("ntwitter").C("users")
	var users []User
	fmt.Println("Trying to find all users of a certain provider")
	err = userConnection.Find(bson.M{"provider": "github"}).All(&users)
	if err != nil {
		panic(err)
	}
	return users
}

func getTweet(session *mgo.Session, err error) Tweet {
	UserConnection := session.DB("ntwitter").C("users")
	var userResult User
	err = UserConnection.Find(bson.M{"username": "vinitkumar"}).One(&userResult)
	fmt.Println("Finding user for a tweet query")
	fmt.Print("userID", userResult.ID)
	TweetConnection := session.DB("ntwitter").C("tweets")
	var tweet Tweet
	err = TweetConnection.Find(bson.M{"user": userResult.ID}).One(&tweet)
	return tweet
}

func getTweets(session *mgo.Session, err error) []Tweet {
	UserConnection := session.DB("ntwitter").C("users")
	var userResult User
	err = UserConnection.Find(bson.M{"username": "vinitkumar"}).One(&userResult)
	TweetConnection := session.DB("ntwitter").C("tweets")
	var tweets []Tweet
	fmt.Println("sample request", bson.M{"user": userResult.ID})
	err = TweetConnection.Find(bson.M{"user": userResult.ID}).All(&tweets)
	return tweets
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	var mongodbURL = os.Getenv("MONGODB_URL")
	session, err := connectDB(mongodbURL, false)
	var user User
	user = getUser(session, err)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div><div>%s</div>", user.Name, user.Email, user)
}

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	var mongodbURL = os.Getenv("MONGODB_URL")
	session, err := connectDB(mongodbURL, false)
	var users []User
	users = getUsers(session, err)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", "Returning Users", users)
}

func TweetsHandler(w http.ResponseWriter, r *http.Request) {
	var mongodbURL = os.Getenv("MONGODB_URL")
	session, err := connectDB(mongodbURL, false)
	var tweets []Tweet
	tweets = getTweets(session, err)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", "Returning Tweets", tweets)
}

func TweetHandler(w http.ResponseWriter, r *http.Request) {
	var mongodbURL = os.Getenv("MONGODB_URL")
	session, err := connectDB(mongodbURL, false)
	var tweet Tweet
	tweet = getTweet(session, err)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", "Returning Tweet", tweet)
}

func AnalyticHandler(w http.ResponseWriter, r *http.Request) {
	var mongodbURL = os.Getenv("MONGODB_URL")
	session, err := connectDB(mongodbURL, false)
	var analytic Analytic
	analytic = getAnalytic(session, err)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", "Returning Analytic", analytic)
}

func AnalyticsHandler(w http.ResponseWriter, r *http.Request) {
	var mongodbURL = os.Getenv("MONGODB_URL")
	session, err := connectDB(mongodbURL, false)
	var analytics []Analytic
	analytics = getAnalytics(session, err)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", "Returning Analytics", analytics)
}

func main() {
	http.HandleFunc("/", UserHandler)
	http.HandleFunc("/users", UsersHandler)
	http.HandleFunc("/tweet", TweetHandler)
	http.HandleFunc("/tweets", TweetsHandler)
	http.HandleFunc("/analytic", AnalyticHandler)
	http.HandleFunc("/analytics", AnalyticsHandler)
	http.ListenAndServe(":8080", nil)
}
