package main

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		(w).Header().Set("Access-Control-Allow-Origin", "*")
		greeting := "Hello!"

		hash := sha512.Sum512([]byte(greeting))

		fmt.Fprint(w, hex.EncodeToString(hash[:]))
	})

    fmt.Printf("Starting server at port 3000\n")
	if err := http.ListenAndServe(":3000", nil); err != nil {
        log.Fatal(err)
    }
}