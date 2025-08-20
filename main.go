///////////////////////////////////////////////////////////////////////////
// Tutorial - Webserver, Round 2
// ⭐️⭐️ DAN 1.15: this is my ACTIVE web server project (ignore any others)
//
// 🟢🟢>> go run main.go
// TO BUILD DOCKER CONTAINER, see instructions at top of the 'Dockerfile'
//
// See Youtube tutorial here: https://www.youtube.com/watch?v=eqvDSkuBihs
//
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
//		git add  (basically moves these to the 'staging' area, which is all files that'll get committed)
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
	"time"
)

//
//
//

// // GLOBAL VARS 🥬🥬🥬🥬🥬🥬🥬🥬🥬🥬🥬🥬🥬🥬🥬🥬🥬🥬🥬🥬🥬🥬🥬🥬🥬🥬🥬🥬🥬🥬🥬🥬🥬🥬
type User struct {
	Name string `json:"name"` // see https://youtu.be/eqvDSkuBihs?si=n4Z6vJfAH6xmdZY9&t=482 for how this works for json
}

// the map KVP structure's format is create it via the 'make' cmd; whats in []s is the Key; whats after is the Value
var userCache = make(map[int]User)

// mutex - an RWMutex blocks both Read and Write
var cacheMutex sync.RWMutex

// // GLOBAL VARS 🥬🥬🥬🥬🥬🥬🥬🥬🥬🥬🥬🥬🥬🥬🥬🥬🥬🥬🥬🥬🥬🥬🥬🥬🥬🥬🥬🥬🥬🥬🥬🥬🥬🥬
//
//

// returnAge - 💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀
func returnAge() {
	emp := make(map[string]int)
	emp["Samia"] = 20
	emp["Sana"] = 23

	fmt.Println(emp["Sana"])
} // returnAge - 💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀
//
//

// main - 🍉🍉🍉🍉🍉🍉🍉🍉🍉🍉🍉🍉🍉🍉🍉🍉🍉🍉🍉🍉🍉🍉🍉🍉🍉🍉🍉🍉🍉🍉🍉🍉
//
// DAN 2.21.25 - reminder, in browser you can hit this endpoint with: http://localhost:8072/get_time
func main() {

	returnAge()

	const name, age = "Kim", 22
	fmt.Fprintf(os.Stdout, "%s is %d years old and that's all.\n", name, age)
	//fmt.Printf("Err msg: %s\n", err)
	log.Printf("feb woowhoooggg: %s", name)

	//✋🏼✋🏼✋🏼✋🏼✋🏼✋🏼✋🏼✋🏼✋🏼✋🏼✋🏼✋🏼✋🏼✋🏼✋🏼 HANDLER FUNCTIONS ✋🏼✋🏼✋🏼✋🏼✋🏼✋🏼✋🏼✋🏼✋🏼✋🏼✋🏼✋🏼✋🏼✋🏼✋🏼✋🏼✋🏼
	fmt.Println("Creating a Mutex")
	danMux := http.NewServeMux()
	danMux.HandleFunc("/", handleRoot)
	danMux.HandleFunc("POST /ip", IpAddy)
	danMux.HandleFunc("GET /ip", returnIpAddress)
	danMux.HandleFunc("POST /users", createUser)
	danMux.HandleFunc("GET /users/{id}", getUser)
	danMux.HandleFunc("DELETE /users/{id}", deleteUser)
	danMux.HandleFunc("/get_time", getTimeHandler)

	// make a handler function to handle a /ip endpoint and return the ip address of the client

	///// THE SERVER STUFF
	const danPort = ":8072"
	fmt.Printf("Feb 9 Super Bowl 🏮🏮🏮🏮: we is are now starting server: %s", danPort)
	http.ListenAndServe(danPort, danMux)
	fmt.Println("Server started and listening")
}

// main - 🍉🍉🍉🍉🍉🍉🍉🍉🍉🍉🍉🍉🍉🍉🍉🍉🍉🍉🍉🍉🍉🍉🍉🍉🍉🍉🍉🍉🍉🍉🍉🍉
// main - 🍉🍉🍉🍉🍉🍉🍉🍉🍉🍉🍉🍉🍉🍉🍉🍉🍉🍉🍉🍉🍉🍉🍉🍉🍉🍉🍉🍉🍉🍉🍉🍉 🦋

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "it's now february 9th 🏈🏈🏈  and the superbowl is on - eagles are demolishing")
}

func returnIpAddress(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is your GET /ip call.  immuna give you your ip")
}

// DAN 2.21.25 752am - replacement from AI Gemini
func readListHandler(w http.ResponseWriter, r *http.Request) {
	// Handle the incoming data for the read list
	// Since we're not using a database, you could just print it to the console for now
	fmt.Println("Received read list data:", r.Body)
	// You might want to add some validation or processing here
	// Send a static response
	fmt.Fprintf(w, "read list received")
}

// DAN 2.21.25 752am - replacement from AI Gemini
func readEntryHandler(w http.ResponseWriter, r *http.Request) {
	// Handle the incoming data for the read entry
	fmt.Println("Received read entry data:", r.Body)
	// Send a static response
	fmt.Fprintf(w, "Read entry received")
}

// DAN 2.21.25 752am - replacement from AI Gemini
func getTimeHandler(w http.ResponseWriter, r *http.Request) {
	// Send the current time as a response
	fmt.Fprintf(w, "/get_time: Current time: %s", time.Now().Format(time.RFC3339))
}

//
//
//
//

// deleteUser 🛑🛑🛑🛑🛑🛑🛑🛑🛑🛑🛑🛑🛑🛑🛑🛑🛑🛑🛑🛑🛑🛑🛑
//
//	2.9.25 with comments added in from GPT
func deleteUser(w http.ResponseWriter, r *http.Request) {
	// Extract User ID from the request path and convert it to an integer.
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		// If conversion fails, respond with HTTP 400 Bad Request and the error message.
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if the user with the given ID exists in the userCache.
	if _, ok := userCache[id]; !ok {
		// If the user does not exist, respond with HTTP 400 Bad Request and a message.
		http.Error(w, "still dont see that user", http.StatusBadRequest)
		return
	}

	// Lock the cacheMutex to ensure thread-safe access to the userCache.
	cacheMutex.Lock()
	// Delete the user from the userCache.
	delete(userCache, id)
	// Unlock the cacheMutex after modifying the userCache.
	cacheMutex.Unlock()

	// Respond with HTTP 204 No Content to indicate successful deletion.
	w.WriteHeader(http.StatusNoContent)
} // 🛑🛑🛑🛑🛑🛑🛑🛑🛑🛑🛑🛑🛑🛑🛑🛑🛑🛑🛑🛑🛑🛑🛑
//
//
//
//

// IpAddy - 🗺️🗺️🗺️🗺️🗺️🗺️🗺️🗺️🗺️🗺️🗺️🗺️🗺️🗺️🗺️🗺️🗺️🗺️🗺️🗺️🗺️🗺️
func IpAddy(w http.ResponseWriter, r *http.Request) {
	ip := r.RemoteAddr
	fmt.Fprintf(w, "POST /ip.  Your IP address is: %s", ip)
} // IpAddy - 🗺️🗺️🗺️🗺️🗺️🗺️🗺️🗺️🗺️🗺️🗺️🗺️🗺️🗺️🗺️🗺️🗺️🗺️🗺️🗺️🗺️🗺️

// DELETE USER = tje version before gpt added comments
// func deleteUser(w http.ResponseWriter, r *http.Request) {
// 	id, err := strconv.Atoi(r.PathValue("id"))
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	if _, ok := userCache[id]; !ok {
// 		http.Error(w, "still dont see that user", http.StatusBadRequest)
// 		return
// 	}

// 	cacheMutex.Lock()
// 	delete(userCache, id)
// 	cacheMutex.Unlock()

// 	// DANHERE 9.11.24 1008pm https://youtu.be/eqvDSkuBihs?si=ke4-eUXJM_9CVklA&t=1345
// 	w.WriteHeader(http.StatusNoContent)
// }
//
//

// getUser - 🙋🙋🙋🙋🙋🙋🙋🙋🙋🙋🙋🙋🙋🙋🙋🙋🙋🙋🙋🙋🙋🙋🙋🙋🙋🙋🙋🙋🙋
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
		return
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
} // getUser - 🙋🙋🙋🙋🙋🙋🙋🙋🙋🙋🙋🙋🙋🙋🙋🙋🙋🙋🙋🙋🙋🙋🙋🙋🙋🙋🙋🙋🙋
//
//

// // createUser - 💥💥💥💥💥💥💥💥💥💥💥💥💥💥💥💥💥💥💥💥💥💥💥💥💥💥💥💥💥💥
//
//	from https://www.youtube.com/watch?v=eqvDSkuBihs&t=796s around minute 10
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
		http.Error(w, "createUser: Please provide a name", http.StatusBadRequest)
		return
	}

	// Add this new user to our fake database - the KVP map we created above
	// use a mutex to ensure we dont have collisions with other I/O on this map
	cacheMutex.Lock()
	userCache[len(userCache)+1] = thisHereUser
	cacheMutex.Unlock()

	// So i guess we're building a Response
	// DAN 2.9 watchin the superbowl - commmetning this out so i can test other endpoitns
	// 	w.WriteHeader(http.StatusNoContent)
	fmt.Fprintf(w, "its feb 9 and almost halftime of the superbowl - eagles are crushing")

} // createUser - 💥💥💥💥💥💥💥💥💥💥💥💥💥💥💥💥💥💥💥💥💥💥💥💥💥💥💥💥💥💥
