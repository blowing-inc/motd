package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/gobuffalo/packr"
)

type Quote struct {
	Message string `json:"message"`
}

var box packr.Box = packr.NewBox("./static")

func randomString(s []string) string {
	rand.Seed(time.Now().Unix())
	return s[rand.Intn(len(s))]
}

func pickMessage() string {
	f, err := box.FindString("quotes.txt")
	if err != nil {
		log.Fatal(err)
	}
	parts := strings.Split(f, "\n")

	return strings.TrimSpace(randomString(parts))
}

func quoteHandler(w http.ResponseWriter, r *http.Request) {
	resp := Quote{Message: pickMessage()}

	data, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func main() {
	// Http Server
	http.HandleFunc("/quote", quoteHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
