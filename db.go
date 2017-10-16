package NodeTwitterAPI

import (
	"fmt"
	"log"
	"os"

	mgo "gopkg.in/mgo.v2"
)

// Connect to the MongoDB
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
		log.Fatal("Sorry! Can't connect to servers right now!")
	}
	session.SetMode(mgo.Monotonic, true)
	return session, err
}

// Get session and returns it for a valid session, else throw an error
func getSession() (session *mgo.Session, err error) {
	var mongodbURL = os.Getenv("MONGODB_URL")
	sessionMongo, errMongo := connectDB(mongodbURL, false)
	if errMongo != nil {
		log.Fatal(errMongo)
	}
	return sessionMongo, errMongo
}
