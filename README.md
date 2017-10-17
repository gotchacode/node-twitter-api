# node-twitter-api


Node Twitter API written with Go and MongoDB

In order to use it, follow this example


```
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	NodeTwitterAPI "github.com/vinitkumar/node-twitter-api"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	http.HandleFunc("/", NodeTwitterAPI.BaseHandler)
	http.HandleFunc("/currentuser", NodeTwitterAPI.UserHandler)
	http.HandleFunc("/users", NodeTwitterAPI.UsersHandler)
	http.HandleFunc("/tweet", NodeTwitterAPI.TweetHandler)
	http.HandleFunc("/tweets", NodeTwitterAPI.TweetsHandler)
	http.HandleFunc("/analytic", NodeTwitterAPI.AnalyticHandler)
	http.HandleFunc("/analytics", NodeTwitterAPI.AnalyticsHandler)
	fmt.Println("Running the server on localhost" + port)

	http.ListenAndServe(":"+port, nil)
}

```

## Documentation

Godoc: https://godoc.org/github.com/vinitkumar/node-twitter-api
