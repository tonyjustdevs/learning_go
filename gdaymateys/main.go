package main

import (
	"fmt"
	"net/http"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "G'day Mateys 🦘 via Golang!\n\n")
	fmt.Fprintf(w, "Saigon Time 🍜: %s", time.Now())
}

func main() {
	http.HandleFunc("/", greet)
	http.ListenAndServe(":8080", nil)
}

// go install module/... : build + place executable into location '~/go/bin'
// go build . : builds executable locally
// go run . : builds and runs locally
// ./ExecutableName : to run it if not in pathq
// ExecutableName : if in ~/go/bin AND ~/go/bin is in path
