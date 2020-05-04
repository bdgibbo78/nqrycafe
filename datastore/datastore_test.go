package datastore

import (
    "testing"
    "../model"
)

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Fatalf("%s != %s", a, b)
	}
}

func TestDataStore(t *testing.T) {

    // Create our datastore
    db := MakeTestStore()

    // Test Category insertion
    cat := model.MakeCategory("protein", "Schnitzel Protein")
    db.NewCategory(cat)
    inv, _ := db.GetInventory()
    assertEqual(t, inv.Categories["protein"].Desc, "Schnitzel Protein")

    // Test Item insertion
    db.NewItem(cat.Id, model.MakeInventoryItem("beef", "Beef", 15.0))
    db.NewItem(cat.Id, model.MakeInventoryItem("chicken", "Chicken", 13.2))

    // Retieve the data using order items
    order1 := model.OrderItem{CategoryId: cat.Id, ItemType: "beef"}
    order2 := model.OrderItem{CategoryId: cat.Id, ItemType: "chicken"}
    order3 := model.OrderItem{CategoryId: cat.Id, ItemType: "fish"}

    ii1, err1 := db.GetInventoryItem(order1)
    assertEqual(t, ii1.Name, "Beef")
    assertEqual(t, ii1.Value, float32(15.0))
    assertEqual(t, err1, nil)

    ii2, err2 := db.GetInventoryItem(order2)
    assertEqual(t, ii2.Name, "Chicken")
    assertEqual(t, ii2.Value, float32(13.2))
    assertEqual(t, err2, nil)

    _, err3 := db.GetInventoryItem(order3)
    assertEqual(t, err3.Error(), "Item with type 'fish' does not exist")
}
