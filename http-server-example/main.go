package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
)

var print = fmt.Println

var validate = validator.New()

func healthCheck(w http.ResponseWriter, req *http.Request) {
	if _, err := w.Write([]byte("Hello there!")); err != nil {
		fmt.Println("Failed to send response. Error:", err)
	}
}

func doStuff(w http.ResponseWriter, req *http.Request) {
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

func reservationHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		jsonData, _ := json.Marshal(Reservations)
		w.Write(jsonData)

	case "POST":
		w.Header().Set("Content-Type", "application/json")
		// a general way of reading req.Body when it could be things other than JSON
		reqBody, _ := io.ReadAll(req.Body)
		var reservationInput Reservation
		if err := json.Unmarshal(reqBody, &reservationInput); err != nil {
			w.Write([]byte("Failed to unmarshal reservation"))
			return
		}

		if err := validate.Struct(reservationInput); err != nil {
			w.Write([]byte(fmt.Sprint(err)))
			return
		}

		reservationCreated := CreateReservation(reservationInput)
		reservationCreatedJson, _ := json.Marshal(reservationCreated)
		w.Write(reservationCreatedJson)

	case "PUT":
		w.Header().Set("Content-Type", "application/json")
		// idiomatic way of reading req.Body when we know it's JSON
		decoder := json.NewDecoder(req.Body)
		var reservationInput Reservation
		if err := decoder.Decode(&reservationInput); err != nil {
			w.Write([]byte("Failed to decode reservation input"))
		}

		if err := validate.Struct(reservationInput); err != nil {
			w.Write([]byte(fmt.Sprint(err)))
			return
		}

		reservationUpdated := UpdateReservation(reservationInput)
		reservationUpdatedJson, _ := json.Marshal(reservationUpdated)
		w.Write(reservationUpdatedJson)
	default:
		http.Error(w, "Disallowed HTTP method on /reservation endpoint", http.StatusBadRequest)
	}
}

func main() {
	print("-- Server Example --")

	http.HandleFunc("/health", healthCheck)
	http.HandleFunc("/stuff", doStuff)
	http.HandleFunc("/complicated", doComplicatedStuff)
	http.HandleFunc("/sound", getSound)
	// we don't specify the method, it routes all of them to the same handler
	http.HandleFunc("/reservation", reservationHandler)

	// blocking call
	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Println("Failed to start server, error:", err)
	}

	fmt.Println("hi")
}
