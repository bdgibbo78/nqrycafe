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

    // The Register for this Controller
    register *Register
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
        register: MakeRegister(),
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

func (self *Controller) GetProduct(pType string) (Product, error) {

    return self.register.NewProduct(pType)
}

//
// Handle a transaction from the client. Return a Sale object or an error on failure.
//
func (self *Controller) HandleProductOrder(product Product) (*Sale, error) {

    sale := NewSale(self.register.nextTransactionId(), 0.0)

    // product is an interface to a Product which can be of any type that conforms
    // to that interface.
    t, err := product.GetTotal(self.datastore)
    if err == nil {
        sale.Cost = t
        return sale, nil
    }
    return nil, err
}
