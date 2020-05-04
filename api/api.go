package api

import (
    "net/http"
    "log"
    "fmt"
    "encoding/json"
    "../pos"

    "github.com/gorilla/mux"
)

//
// An interface to allow objects conforming to it to handle requests from
// the API.
//
type APIHandler interface {

    // Get the current inventory to return to clients to render
    GetInventory() (interface{}, error)

    // Get the Product of type given by 'ptype' or error on failure
    GetProduct(pType string) (pos.Product, error)

    // Handle a product order from the client returning a Sale or error on failure
    HandleProductOrder(product pos.Product) (*pos.Sale, error)
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

    router := mux.NewRouter()

	router.HandleFunc("/api/v1/inventory", endpoint.InventoryHandler)
    router.HandleFunc("/api/v1/order/{ptype}", endpoint.CoffeeHandler)

    http.Handle("/", router)

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

func (self *APIEndpoint) CoffeeHandler(w http.ResponseWriter, req *http.Request) {
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

    vars := mux.Vars(req)
	pType := vars["ptype"]

    // Get the object to be decoded
    product, err := self.handler.GetProduct(pType)

    decoder := json.NewDecoder(req.Body)
    err = decoder.Decode(product)
    if err != nil {
        log.Printf("Decoding error: %e", err)
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    sale, err := self.handler.HandleProductOrder(product)
    if err != nil {
        log.Printf("Transaction error: %e", err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")

    s, _ := json.Marshal(&sale)
    fmt.Fprintf(w, string(s))
}
