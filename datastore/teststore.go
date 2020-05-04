package datastore

import (
    "errors"
    "../model"  // import our models
)

//
// TestStore conforms to the DataStore interface thus allowing
// consumers to load or modify coffee orders
//
type TestStore struct {

    // Reuse Inventory here for our storage purposes
    inventory model.Inventory

}

func MakeTestStore() TestStore {
    ts := TestStore{
        inventory: model.MakeInventory(),
    }
    return ts
}

func (self *TestStore) Init() error {

    // Add available Categories
    ctypes := model.MakeCategory("coffeetype", "Coffee Types")
    cpreps := model.MakeCategory("coffeeprep", "Coffee Preparation")
    csizes := model.MakeCategory("coffeesize", "Coffee Size")
    cconds := model.MakeCategory("coffeeconds", "Coffee Condiments")

    self.NewCategory(ctypes)
    self.NewCategory(cpreps)
    self.NewCategory(csizes)
    self.NewCategory(cconds)

    // Add Coffee types
    self.NewItem(ctypes.Id, model.MakeInventoryItem("house_blend", "House Blend", 1.0))
    self.NewItem(ctypes.Id, model.MakeInventoryItem("dark_roast", "Dark Roast", 1.5))
    self.NewItem(ctypes.Id, model.MakeInventoryItem("robusta", "Robusta", 2.0))
    self.NewItem(ctypes.Id, model.MakeInventoryItem("arabica", "Arabica", 2.5))

    // Add Coffee preparations
    self.NewItem(cpreps.Id, model.MakeInventoryItem("espresso", "Espresso", 1.0))
    self.NewItem(cpreps.Id, model.MakeInventoryItem("latte", "Latte", 1.25))
    self.NewItem(cpreps.Id, model.MakeInventoryItem("cappuccino", "Cappuccino", 1.5))
    self.NewItem(cpreps.Id, model.MakeInventoryItem("macchiato", "Macchiato", 1.75))
    self.NewItem(cpreps.Id, model.MakeInventoryItem("mocha", "Mocha", 2.0))

    // Add Coffee sizes
    self.NewItem(csizes.Id, model.MakeInventoryItem("standard", "Standard", 1.0))
    self.NewItem(csizes.Id, model.MakeInventoryItem("child", "Child", 0.75))
    self.NewItem(csizes.Id, model.MakeInventoryItem("large", "Large", 1.5))
    self.NewItem(csizes.Id, model.MakeInventoryItem("addict", "Addict", 2.0))

    // Add Condiments
    self.NewItem(cconds.Id, model.MakeInventoryItem("milk", "Milk", 1.0))
    self.NewItem(cconds.Id, model.MakeInventoryItem("sugar", "Sugar", 0.25))
    self.NewItem(cconds.Id, model.MakeInventoryItem("coco", "Coco Powder", 0.1))
    return nil
}

func (self *TestStore) NewCategory(cat model.Category) {
    self.inventory.Categories[cat.Id] = cat
    self.inventory.Items[cat.Id] = make(map[model.ItemType]model.InventoryItem)
}

// Add a new item to the database under the category 'cat'
func (self *TestStore) NewItem(cat model.CategoryId, item model.InventoryItem) error {
    // first check that the category exists
    _, ok := self.inventory.Items[cat]
    if !ok {
        return errors.New("Category with id '" + string(cat) + "' does not exist")
    }
    self.inventory.Items[cat][item.Type] = item
    return nil
}

func (self *TestStore) GetInventory() (*model.Inventory, error) {
    return &self.inventory, nil
}

func (self *TestStore) GetInventoryItem(it model.OrderItem) (*model.InventoryItem, error) {

    items, ok := self.inventory.Items[it.CategoryId]
    if !ok {
        return nil, errors.New("Category does not exist")
    }

    item, ok := items[it.ItemType]
    if !ok {
        return nil, errors.New("Item with type '" + string(it.ItemType) + "' does not exist")
    }
    return &item, nil
}
