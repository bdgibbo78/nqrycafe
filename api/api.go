package api

import (
    "net/http"
    "log"
    "fmt"
    "encoding/json"

    "../pos"
)

//
// An interface to allow objects conforming to it to handle requests from
// the API.
//
type APIHandler interface {

    // Get the current inventory to return to clients to render
    GetInventory() (interface{}, error)

    // Handle a client transaction and return an interface to the client
    HandleTransaction(transaction pos.Transaction) (interface{}, error)
}

//
// The endpoint for our API. It provides the particular handler required
// to handle requests for this endpoint.
//
type APIEndpoint struct {

    // A handler to handle client requests
	handler APIHandler
}

func ServeHTTP(handler APIHandler) {

    endpoint := APIEndpoint{handler: handler}

	http.HandleFunc("/api/v1/inventory", endpoint.InventoryHandler)
    http.HandleFunc("/api/v1/transaction", endpoint.TransactionHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (self *APIEndpoint) InventoryHandler(w http.ResponseWriter, req *http.Request) {
    if req.Method != http.MethodGet {
		// Sync requests must be GET
        w.WriteHeader(http.StatusNotAcceptable)
		return
	}

    inventory, err := self.handler.GetInventory()
    if err != nil {
        log.Printf("Decoding error: %e", err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")

    s, _ := json.Marshal(inventory)
    fmt.Fprintf(w, string(s))
}

func (self *APIEndpoint) TransactionHandler(w http.ResponseWriter, req *http.Request) {
    if req.Method != http.MethodPost {
		// Sync requests must be POST
        w.WriteHeader(http.StatusNotAcceptable)
		return
	}

    ct := req.Header.Get("Content-Type")
	if ct != "application/json" {
        log.Print("Error: not a JSON request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

    // Decode the JSON post into a transaction object
    transaction := pos.MakeTransaction()
    decoder := json.NewDecoder(req.Body)
    err := decoder.Decode(&transaction)
    if err != nil {
        log.Printf("Decoding error: %e", err)
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    sale, err := self.handler.HandleTransaction(transaction)
    if err != nil {
        log.Printf("Transaction error: %e", err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")

    s, _ := json.Marshal(&sale)
    fmt.Fprintf(w, string(s))
}
