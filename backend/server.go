package main

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net/http"
	"server/database"
)

func main() {
	database.Connection();
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		(w).Header().Set("Access-Control-Allow-Origin", "*")
		
		r.ParseMultipartForm(0)

		file, fileHeader, err := r.FormFile("file")

		fmt.Printf("File being processed is: %s\n", fileHeader.Filename)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		fileHash := sha512.New()

		if _,err := io.Copy(fileHash, file); err != nil {
			log.Fatal(err)
		}

		hash := fileHash.Sum(nil)

		fmt.Fprint(w, hex.EncodeToString(hash))
	})

    fmt.Printf("Starting server at port 3000\n")
	if err := http.ListenAndServe(":3000", nil); err != nil {
        log.Fatal(err)
    }
}