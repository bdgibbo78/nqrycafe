package pos

import (
    "errors"
    "../datastore"
)

//
// A interface from which all Products need to conform to.
//
type Product interface {

    GetTotal(ds datastore.DataStore) (float32, error)
}

//
// A Register contains a map of functions to create a specific product.
// Each product function is keyed by its product type.
// New product types are added via the init() function.
//
type Register struct {
    products map[string]interface{}
    transactionId int
}

func MakeRegister() *Register {
    reg := Register{
        products: make(map[string]interface{}),
        transactionId: 0,
    }
    reg.init()
    return &reg
}

func (self *Register) init() {
    self.products["coffee"] = NewCoffee

    // Register new products here
}

func (self *Register) nextTransactionId() int {
    self.transactionId += 1
    return self.transactionId
}

func (self *Register) NewProduct(pType string) (Product, error) {
    fn, ok := self.products[pType]
    if !ok {
        return nil, errors.New("Product '" + pType + "' does not exist in the register")
    }

    c, ok := fn.(func() Product)
    if !ok {
        return nil, errors.New("Product '" + pType + "' does not conform to Product")
    }
    obj := c()
    return obj, nil
}

//
// Sale describes a sale to the customer. It contains the identifier of the
// transaction and a total cost.
//
type Sale struct {
    TransactionId int   `json:"transactionid"`
    Cost float32        `json:"cost"`
}

func NewSale(transactionId int, cost float32) *Sale {
    return &Sale{TransactionId: transactionId, Cost: cost}
}
