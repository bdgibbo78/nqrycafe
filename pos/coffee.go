package pos

import (
    "../model"
    "../datastore"
)

//
// Coffee defines the structure of a coffee order. That is, it
// requires a coffee type, a preparation, a size multiplier and any additional
// condiments.
//
type Coffee struct {
    Type model.OrderItem
    Prep model.OrderItem
    Size model.OrderItem
    Cond model.OrderItemList
}

func MakeCoffee(brand model.ItemType, prep model.ItemType, size model.ItemType, conds []model.ItemType) Coffee {
    coffee := Coffee{
        Type: model.OrderItem{CategoryId: "coffeetype", ItemType: brand},
        Prep: model.OrderItem{CategoryId: "coffeeprep", ItemType: prep},
        Size: model.OrderItem{CategoryId: "coffeesize", ItemType: size},
        Cond: model.OrderItemList{CategoryId: "coffeeconds", ItemTypes: conds},
    }
    return coffee
}

//
// Conforms to the Product interface
//
func (self *Coffee) GetTotal(ds datastore.DataStore) (float32, error) {

    // The cost for this order
    var orderCost float32 = 0.0

    // Coffee Type
    coffeeType, err := ds.GetInventoryItem(self.Type)
    if err != nil {
        return 0.0, err
    }
    orderCost += coffeeType.Value

    // Coffee preparation
    coffeePrep, err := ds.GetInventoryItem(self.Prep)
    if err != nil {
        return 0.0, err
    }
    orderCost += coffeePrep.Value

    // Coffee size
    coffeeSize, err := ds.GetInventoryItem(self.Size)
    if err != nil {
        return 0.0, err
    }
    orderCost *= coffeeSize.Value // multiplier

    // Coffee condiments
    for _, cond := range self.Cond.ItemTypes {
        oItem := model.OrderItem{
            CategoryId: self.Cond.CategoryId,
            ItemType: cond}
        coffeeCond, err := ds.GetInventoryItem(oItem)
        if err != nil {
            return 0.0, err
        }
        orderCost += coffeeCond.Value
    }
    return orderCost, nil
}
