package main

import (
	"fmt"
	"net/http"
	"log"
	"crypto/sha1"
	"os"
	"encoding/hex"
	"bufio"
	"strconv"
)

var secret string

func main() {
	if os.Getenv("VERSION") == "" {
		log.Fatal("MUST DEFINE A VERSION")
	}
	i1, err := strconv.Atoi(os.Getenv("VERSION"))
	if i1 < 32110 {
		log.Fatal("VERSION MUST BE 32110 OR GREATER")
	}

	fmt.Println("running version", i1)

	file, err := os.Open("/myemailaddress.txt")
	if err != nil {
			log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
    h := sha1.New()
    h.Write([]byte(line + "super secret something"))
    sha1_hash := hex.EncodeToString(h.Sum(nil))
		secret += line + " " + sha1_hash
}

	if scanner.Err() != nil {
			fmt.Printf(" > Failed!: %v\n", scanner.Err())
	}

	http.HandleFunc("/", hello)

	listenOnASocket("0.0.0.0:8080")
}

func listenOnASocket(addr string) {
	fmt.Println("Listening on:", addr)
	http.ListenAndServe(addr, nil)
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, secret)
	defer req.Body.Close()
}