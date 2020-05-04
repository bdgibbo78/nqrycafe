package pos

import (
    "../datastore"
)

//
// A controller coordinates the flow of data between the api and the datastore
// by conforming to the APIHandler interface.
//
type Controller struct {

    // The datastore to use for this controller
    datastore datastore.DataStore
}

//
// Create our controller and return a pointer to it or an error on failure
//
func MakeController(datastore datastore.DataStore) (*Controller, error) {

    // Check that we can retrieve an inventory from the datastore
    _, error := datastore.GetInventory()
    if error != nil {
        return nil, error
    }

    controller := Controller{
        datastore: datastore,
    }
    return &controller, nil
}

//
// Return a pointer to the current inventory or nil and an error on failure.
//
func (self *Controller) GetInventory() (interface{}, error) {
    inventory, error := self.datastore.GetInventory()
    if error != nil {
        return nil, error
    }
    return inventory, nil
}

//
// Handle a transaction from the client. Return a Sale object or an error on failure.
//
func (self *Controller) HandleTransaction(transaction Transaction) (interface{}, error) {

    sale := Sale{TransactionId: transaction.Id, Cost: 0.0}

    // Calculate the total cost given the current inventory
    var total float32 = 0.0
    for _, product := range transaction.Order.Products {

        // product is an interface to a Product which can be of any type that conforms
        // to that interface.
        t, err := product.GetTotal(self.datastore)
        if err != nil {
            return 0.0, err
        }
        total += t
    }

    // Otherwise, update our sale and return it
    sale.Cost = total
    return sale, nil
}
