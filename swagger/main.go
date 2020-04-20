// Testing go-swagger generation
//
// The purpose of this application is to test go-swagger in a simple GET request.
//
//     Schemes: http
//     Host: localhost:8080
//     Version: 0.0.1
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: Daniel<danielfs.ti@gmail.com>
//
//     Consumes:
//     - text/plain
//
//     Produces:
//     - text/plain
//
// swagger:meta
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type message struct {
	Description string `json:"description"`
}

// Success response
// swagger:response ok
type swaggerResponse struct {
	// in: body
	body struct {
		// HTTP status code 200 - Status OK
		message
	}
}

// response with static swagger file
func swagger(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	http.ServeFile(w, r, "swagger.json")
}

// HTTP status code 200 and an array of repository models in data
// swagger:response reposResp
type swaggReposResp struct {
	// in:body
	Body struct {
		// HTTP status code 200 - Status OK
		Code int `json:"code"`
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	// swagger:operation GET /hello/{name} hello Hello
	//
	// Returns a simple Hello message
	// ---
	// consumes:
	// - text/plain
	// produces:
	// - text/plain
	// parameters:
	// - name: name
	//   in: path
	//   description: Name to be returned.
	//   required: true
	//   type: string
	// responses:
	//   '200':
	//     description: The hello message
	//     type: string

	log.Println("Responsing to /hello request")
	log.Println(r.UserAgent())

	name := mux.Vars(r)["name"]

	fmt.Fprintln(w, "Hello:", name)
}

func handleMessage(w http.ResponseWriter, r *http.Request) {
	// swagger:operation GET /message message Message
	//
	// Returns a simple Message
	// ---
	// consumes:
	// - text/plain
	// produces:
	// - application/json
	// responses:
	//   "200":
	//     "$ref": "#/responses/ok"

	json.NewEncoder(w).Encode(message{Description: "this is a message"})
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/hello/{name}", index).Methods(http.MethodGet)
	router.HandleFunc("/message", handleMessage).Methods(http.MethodGet)
	router.HandleFunc("/swagger.json", swagger).Methods(http.MethodGet)

	// adding cors
	handler := cors.Default().Handler(router)

	log.Fatal(http.ListenAndServe(":8080", handler))
}
