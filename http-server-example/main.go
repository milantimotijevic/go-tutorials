package main

import (
	"fmt"
	"net/http"
	"time"
)

var print = fmt.Println

func healthCheck(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello there!"))
}

func doStuff(w http.ResponseWriter, req *http.Request) {
	print(req.Header)
	for id, header := range req.Header {
		print(id, " - ", header)
	}
	w.Write([]byte("Stuff got done"))
}

func doComplicatedStuff(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	select {
	case <-time.After(time.Second * 5):
		w.Write([]byte("Complicated stuff got done"))
	case <-ctx.Done():
		// this channel is only used to handle timeouts and cancelations/interruptions
		// it doesn't get triggered after a response is successfully sent
		print("The bastard probably canceled their request. Let's store their IP so we can DDOS them, teach them a lesson!")
		err := ctx.Err()
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getSound(w http.ResponseWriter, req *http.Request) {
	defer func() {
		// ward off any potential panics that may unwind the stack up to this function
		if r := recover(); r != nil {
			print("Server encountered the following error:", r)
			http.Error(w, "Uknown error on server side", http.StatusInternalServerError)
		}
	}()

	animal := req.URL.Query().Get("animal")

	switch animal {
	case "cat":
		w.Write([]byte("Meow"))
	case "dog":
		w.Write([]byte("Woof"))
	case "pig":
		w.Write([]byte("Oink"))
	case "fox":
		// a proper 400 error
		http.Error(w, "Oh no no no, we don't do that here!", http.StatusBadRequest)
	case "moose":
		panic("OMG, we don't know the plural form of 'moose'! Is it 'moose', 'mooses' or 'meese'???")
	default:
		w.Write([]byte("I don't know, bro... I don't know"))
	}
}

func main() {
	print("-- Server Example --")

	http.HandleFunc("/health", healthCheck)
	http.HandleFunc("/stuff", doStuff)
	http.HandleFunc("/complicated", doComplicatedStuff)
	http.HandleFunc("/sound", getSound)

	http.ListenAndServe(":3000", nil)
}
