///////////////////////////////////////////////////////////////////////////
// Tutorial - Webserver, Round 2
// ⭐️⭐️ DAN 1.15: this is my ACTIVE web server project (ignore any others)
//
// 🟢🟢>> go run main.go
//
// See Youtube tutorial here: https://www.youtube.com/watch?v=eqvDSkuBihs
// d
//
//  TODO 1.15.25: none right now (need to focus on Docker and TF)
//
//
// My Go Dev Journal:
//   dan chisholm, 8.27.24... about the first program i've written since May 2004 when i left EA
//
//   DAN 1.11.25 annnnd we're back after a few months and after the new year
//   to set up docker with go dont forget you have to run commandline: go mod init blue2woo.com/m
//    which i have no idea what it does.  but afterward your 'docker build" will work
//
//  Git Commands 2.4.25
//		git diff
//		git status
//		git 
//		git commit -am // the 'a' here is for all, but really just means 'add' so you can skip 'git add'
//		git config -l // lists all commit configs
//      // COMING SOON - Tony Teaches Tech - https://www.youtube.com/watch?v=yxvqLBHZfXk
//		TODO
//
///////////////////////////////////////////////////////////////////////////

package main

import (
	"encoding/json" // DANTODO: research this package
	"fmt"
	"log" // DANTODO: try using this package to log events
	"net/http"
	"os"
	"strconv"
	"sync"
)

// // GLOBAL VARS /////
type User struct {
	Name string `json:"name"` // see https://youtu.be/eqvDSkuBihs?si=n4Z6vJfAH6xmdZY9&t=482 for how this works for json
}

// the map KVP structure's format is create it via the 'make' cmd; whats in []s is the Key; whats after is the Value
var userCache = make(map[int]User)

// mutex - an RWMutex blocks both Read and Write
var cacheMutex sync.RWMutex

// just testing a KVP, ie a map

func returnAge() {
	emp := make(map[string]int)
	emp["Samia"] = 20
	emp["Sana"] = 23

	fmt.Println(emp["Sana"])
}

// /// MAIN /////
func main() {

	returnAge()

	const name, age = "Kim", 22
	fmt.Fprintf(os.Stdout, "%s is %d years old and that's all.\n", name, age)
	//fmt.Printf("Err msg: %s\n", err)
	log.Printf("DWC-LOG: %s", name)

	///// HANDLER FUNCTIONS - ie 3 or 4 API endpoints (yes each is a separate endpoint)
	fmt.Println("Creating a Mutex")
	danMux := http.NewServeMux()
	danMux.HandleFunc("/", handleRoot)
	danMux.HandleFunc("POST /users", createUser)
	danMux.HandleFunc("GET /users/{id}", getUser)
	danMux.HandleFunc("DELETE /users/{id}", deleteUser)

	///// THE SERVER STUFF
	const danPort = ":8072"
	fmt.Printf("Starting dah Server on port: %s", danPort)
	http.ListenAndServe(danPort, danMux)
	fmt.Println("Server started and listening")
}

// ////////////////////
// ////////////////////
func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "It is Janary 11 of 2025 and the fires are still burning in LA")
}

// DELETE USER
func deleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if _, ok := userCache[id]; !ok {
		http.Error(w, "still dont see that user", http.StatusBadRequest)
		return
	}

	cacheMutex.Lock()
	delete(userCache, id)
	cacheMutex.Unlock()

	// DANHERE 9.11.24 1008pm https://youtu.be/eqvDSkuBihs?si=ke4-eUXJM_9CVklA&t=1345
	w.WriteHeader(http.StatusNoContent)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	// this is the id from the 'GET /users' listener above
	// we're also casting it to an int
	// we're also doing error checking in case the string is actually "id"
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// not sure what we're doin here, i guess we're getting the struct of the user 'id'
	cacheMutex.RLock()
	user, ok := userCache[id]
	cacheMutex.RUnlock()

	// error check if the user was found
	if !ok {
		http.Error(w, "they aint no such user", http.StatusNotFound)
	}

	// set up so that our response is known by the requestor (ie client) that its json format
	w.Header().Set("Content-Type", "application/json")

	// now its time to take our valid user and marshal their info into a valid json payload
	j, err := json.Marshal(user)
	// now since we know the user is legit, if there is any error then its' got to be on 'our'
	// side, ie in the server.  b/c the request is legit... so InternalServerError
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// now time to write out the json response
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// ////////// from https://www.youtube.com/watch?v=eqvDSkuBihs&t=796s around minute 10
func createUser(w http.ResponseWriter, r *http.Request) {
	var thisHereUser User
	err := json.NewDecoder(r.Body).Decode(&thisHereUser)

	// Error Check 1: if the error is not nothing (ie it IS something real) then bail out
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print(err) // DAN 8.28: testing out this command which writes to line
		return
	}

	// Error Check 2: No empty names allowed
	if thisHereUser.Name == "" {
		http.Error(w, "Please provide a name", http.StatusBadRequest)
		return
	}

	// Add this new user to our fake database - the KVP map we created above
	// use a mutex to ensure we dont have collisions with other I/O on this map
	cacheMutex.Lock()
	userCache[len(userCache)+1] = thisHereUser
	cacheMutex.Unlock()

	// So i guess we're building a Response
	w.WriteHeader(http.StatusNoContent)
}
