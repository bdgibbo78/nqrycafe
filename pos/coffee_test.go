package pos

import (
    "testing"
    "fmt"
    "../model"
    "../datastore"
)

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Fatalf("%s != %s", a, b)
	}
}

func TestCoffeeOrders(t *testing.T) {

    db := datastore.MakeTestStore()
    _ = db.Init()

    c1 := MakeCoffee("arabica", "latte", "large", []model.ItemType{"sugar"})
    cost, _ := c1.GetTotal(&db)
    assertEqual(t, cost, float32(5.875))

    c2 := MakeCoffee("robusta", "espresso", "addict", []model.ItemType{"milk", "coco"})
    cost, _ = c2.GetTotal(&db)
    assertEqual(t, cost, float32(7.1))

    c3 := MakeCoffee("house_blend", "macchiato", "standard", []model.ItemType{"coco"})
    cost, _ = c3.GetTotal(&db)
    assertEqual(t, cost, float32(2.85))
}

func TestNewCoffee(t *testing.T) {
    db := datastore.MakeTestStore()
    _ = db.Init()

    db.NewItem("coffeetype", model.MakeInventoryItem("instant", "Instant Coffee", 0.5))
    db.NewItem("coffeeprep", model.MakeInventoryItem("kettle", "Mixed with boiling water", 0.2))
    db.NewItem("coffeesize", model.MakeInventoryItem("thermos", "Thermos", 2.5))
    db.NewItem("coffeeconds", model.MakeInventoryItem("honey", "Honey", 1.25))

    c := MakeCoffee("instant", "kettle", "thermos", []model.ItemType{"honey", "milk"})
    cost, err := c.GetTotal(&db)
    if err != nil {
        fmt.Printf("Error: %e\n", err)
    }
    assertEqual(t, cost, float32(4.0))
}

func TestInvalidCoffee(t *testing.T) {
    db := datastore.MakeTestStore()
    _ = db.Init()

    c := MakeCoffee("instant", "kettle", "thermos", []model.ItemType{"honey", "milk"})
    cost, err := c.GetTotal(&db)
    assertEqual(t, cost, float32(0.0))
    assertEqual(t, err.Error(), "Item with type 'instant' does not exist")
}

/*
func TestCoffeeTransaction(t *testing.T) {

    db := datastore.MakeTestStore()
    _ = db.Init()
    c, _ := MakeController(&db)

    trans := MakeTransaction()

    sale, err := c.HandleTransaction(trans)
    if err != nil {
        fmt.Printf("Error: %e\n", err)
    }

    fmt.Printf("Cost: %f\n", sale.Cost)
}
*/
