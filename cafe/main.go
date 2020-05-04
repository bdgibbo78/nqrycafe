package main

import (
    "log"
    "../pos"
    "../datastore"
    "../api"
)

func main() {

    // Create our datastore and initialise it
    datastore := datastore.MakeTestStore()
    err := datastore.Init()
    if err != nil {
        log.Fatal("Failed to initialise datastore.")
    }

    // Create our Controller configured to use our datastore
    controller, err := pos.MakeController(&datastore)
    if err != nil {
        log.Fatal("Failed to create the controller.")
    }

    // Finally, start serving HTTP JSON API using our controller
    api.ServeHTTP(controller)


}
