package model

type CategoryId string
type ItemType string

//
// Category describes a particular grouping of items in the inventory.
//
// Id [CategoryId]: A unique identifier for this category
// Desc [string]: A description of the category for display purposes.
//
type Category struct {
    Id CategoryId       `json:"id"`
    Desc string         `json:"description"`
}

func MakeCategory(id CategoryId, desc string) Category {
    return Category{Id: id, Desc: desc}
}

//
// InventoryItem describes an abstraction for individual elements
// in the coffee order process.
//
// Type [string]: A unique identifier for the particular type of item
// Name [string]: A display name for the item
// Value [float32]: The value associated with this item
//
type InventoryItem struct {
    Type ItemType      `json:"type"`
    Name string        `json:"name"`
    Value float32      `json:"value"`
}

func MakeInventoryItem(it ItemType, name string, value float32) InventoryItem {
    return InventoryItem{Type: it, Name: name, Value: value}
}

//
// Inventory describes a list of all available coffee types, preps, sizes and
// condiments.
//
// Categories [map]: A map containing the categories in the inventory keyed by their id.
// Items [map[map]]: A container of all the items and their associated category in the inventory.
//
type Inventory struct {

    // Available Categories
    Categories map[CategoryId]Category               `json:"categories"`

    // Available items
    Items map[CategoryId]map[ItemType]InventoryItem  `json:"items"`
}

//
// Make an empty inventory
//
func MakeInventory() Inventory {
    iv := Inventory{
        Categories: make(map[CategoryId]Category),
        Items: make(map[CategoryId]map[ItemType]InventoryItem),
    }
    return iv
}

//
// OrderItem describes an item selection from a given category.
//
// CategoryId [CategoryId]: The id of the category
// ItemType [ItemType]: The type (or identifier) of the selected item
//
type OrderItem struct {
    CategoryId CategoryId   `json:"categoryid"`
    ItemType ItemType       `json:"choice"`
}

//
// OrderItemList describes 1 or more item selections from a given category.
//
// CategoryId [CategoryId]: The id of the category
// ItemTypes [[]ItemType]: A list of types (or identifiers) of selected items.
//
type OrderItemList struct {
    CategoryId CategoryId   `json:"categoryid"`
    ItemTypes []ItemType    `json:"choices"`
}
