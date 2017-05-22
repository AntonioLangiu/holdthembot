package common

import (
	"fmt"
	"net/http"
	"os"
)

// This small webserver is needed to bind to port 80 in heroku
func WebServer() {
	http.HandleFunc("/", hello)
	fmt.Println("listening...")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "HoldThemBot")
}
