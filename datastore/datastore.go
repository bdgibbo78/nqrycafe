package datastore

import (
    "../model"
)

//
// DataStore describes our interface for all configured types of persistent
// data storage.
//
type DataStore interface {

    // Do an initial load from the datastore. Return an error on failure.
    Init() error

    // Add a new Category to the datastore
    NewCategory(cat model.Category)

    // Add a new item to the database under the category 'cat'
    NewItem(cat model.CategoryId, item model.InventoryItem) error

    // Get the current Inventory or return error on failure
    GetInventory() (*model.Inventory, error)

    // Get the InventoryItem with a given item type within given CategoryId.
    GetInventoryItem(it model.OrderItem) (*model.InventoryItem, error)
}
