package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/lamecksilva/leaking-bucket-go/ratelimit"
)

func main() {
	log.Println("Leaking Bucket Go")

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hello world route")
		fmt.Fprintln(w, "ok")
	})

	limiter := ratelimit.NewLeakyBucketLimiter(
		10,
		5,
	)

	handler := ratelimit.LeakyBucketMiddleware(limiter)(mux)

	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}
