package pos

import (
    "../datastore"
)

type Product interface {

    GetTotal(ds datastore.DataStore) (float32, error)
}

//
// An order contains multiple products to be purchased
//
type Order struct {
    Products []Product
}

func MakeOrder() Order {
    return Order{Products: make([]Product, 0)}
}

//
// Transaction describes a customer transaction. It contains 1 or more orders and
// a transaction identifier.
//
type Transaction struct {
    Id int              `json:"id"`
    Order Order         `json:"order"`
}

func MakeTransaction() Transaction {
    t := Transaction{
        Id: 0,
        Order: MakeOrder(),
    }
    return t
}

//
// Sale describes a sale to the customer. It contains the identifier of the
// transaction and a total cost.
//
type Sale struct {
    TransactionId int   `json:"transactionid"`
    Cost float32        `json:"cost"`
}

func MakeSale(transactionId int, cost float32) Sale {
    return Sale{TransactionId: transactionId, Cost: cost}
}
