// Goals:
// The goals for this API is following.
// - Provide API endpints for tweets, users, analytics, comments
// - Small isolated modular functions that provide an easy interface to build
// upon.

package NodeTwitterAPI

import (
	"encoding/json"
	"fmt"
	"net/http"

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

// Analytic struct
type Analytic struct {
	ID   bson.ObjectId `bson:"_id,omitempty"`
	IP   string
	User bson.ObjectId
	URL  string
}

// Routes   routes of the API
type Routes struct {
	Title     string
	FirstURL  string
	SecondURL string
	ThirdURL  string
	ForthURL  string
	FifthURL  string
	SixthURL  string
}

// Query Map
type Query map[string]string

// Get Analytic and return it.
func getAnalytic(session *mgo.Session, err error) Analytic {
	AnalyticConnection := session.DB("ntwitter").C("analytics")
	var analytic Analytic
	err = AnalyticConnection.Find(nil).One(&analytic)
	return analytic
}

// Get list of all analytics and returns it.
func getAnalytics(session *mgo.Session, err error) []Analytic {
	AnalyticConnection := session.DB("ntwitter").C("analytics")
	var analyticsList []Analytic
	err = AnalyticConnection.Find(nil).All(&analyticsList)
	return analyticsList
}

// Get current user and return it.
func getUser(session *mgo.Session, err error) User {
	userConnection := session.DB("ntwitter").C("users")
	var user User
	fmt.Println("Trying to query one users")
	err = userConnection.Find(bson.M{"username": "vinitkumar"}).One(&user)
	return user
}

// Get list of all users and returns it.
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

// Gets tweets by an user and reurns it.
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

// Gets list of all tweets and returns it.
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

// UserHandler   : Returns the current user of the session.
func UserHandler(w http.ResponseWriter, r *http.Request) {
	session, err := getSession()
	var user User
	user = getUser(session, err)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

// UsersHandler   : Returns list of all users in the database.
func UsersHandler(w http.ResponseWriter, r *http.Request) {
	session, err := getSession()
	var users []User
	users = getUsers(session, err)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

// TweetsHandler   : Returns list of all tweets in the database.
func TweetsHandler(w http.ResponseWriter, r *http.Request) {
	session, err := getSession()
	var tweets []Tweet
	tweets = getTweets(session, err)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tweets)
}

// TweetHandler  : Returns tweet by an user in the database.
func TweetHandler(w http.ResponseWriter, r *http.Request) {
	session, err := getSession()
	var tweet Tweet
	tweet = getTweet(session, err)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tweet)
}

// AnalyticHandler   : Return one analytic entry from the database.
func AnalyticHandler(w http.ResponseWriter, r *http.Request) {
	session, err := getSession()
	var analytic Analytic
	analytic = getAnalytic(session, err)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(analytic)
}

// AnalyticsHandler   : Returns list of all analytics in the database.
func AnalyticsHandler(w http.ResponseWriter, r *http.Request) {
	session, err := getSession()
	var analytics []Analytic
	analytics = getAnalytics(session, err)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(analytics)
}

// BaseHandler   Show a landing page with links of the APIs.
func BaseHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Doctor", "Strange")
	var homePage = `<h1>Welcome to node twitter API</h1>
	
	<br>
	For more details: Please visit the following:
	<ul>
	<li>
	<a href="/currentuser">Current user</a>
	</li>
	<li>
	<a href="/users">Users List</a> 	
	</li>
	<li>
	<a href="/tweet">Tweet</a> 
	</li>
	<li>
	<a href="/tweets">Tweets</a> 
	</li>
	<li>
	<a href="/analytic">Analytic</a> 
	</li>
	<li>	
	<a href="/analytics">Analytics</a> 
	</li>
	`
	fmt.Fprintf(w, homePage)
}
