# nqrycafe
## NQRY Cafe PoS

NQRY Cafe runs a simple JSON API server over HTTP. Its purpose is to demonstrate
a simple Point of Sale (PoS) domain model for ordering coffee from the cafe. It
has been designed to be flexible so that extra coffee types, preparation methods,
sizes and condiments can be easily added to the system.

It is assumed that the UI and elements of its implementation is being handled elsewhere
with the state of the system being accessed via the API.

## Build and Run
Make sure that Go has been installed.
`> cd cafe/
> go build
> ./cafe`

Open a browser and navigate to **http://localhost:8080/api/v1/inventory**

## Design Details
The design is broken down into 4 modules plus the top level main function (cafe/main.go)
to provide the entry point to the application. This design was adopted as it breaks down the
solution into smaller chunks that are easier to maintain, extend and test. By providing
"minimally complete" interfaces to each module, we can ensure the data flow between them is
consistent and optimized.

The modules can be found in the associated folders and are described below.

### Model (model/model.go)
The Model module describes the models used by the system. Models at this
level are generic in that products can be built up using one or more of these models.
Structures defining inventory **Categories** and their contained **InventoryItems** are
provided as well as functions to be able to create them when needed. Structures **OrderItem**
and **OrderItemList** are also provided to describe one or more item selections from a
category.

### DataStore (datastore/datastore.go)
The DataStore module provides a convenient generic interface that allows us to store
persistent model data. It contains an abtract **DataStore** interface and can include any
number of concrete implementations of persistent data storage.
For example, we are able to provide implementations to support databases (PostgreSQL, MongoDB etc) or a simple file based data store depending on the customer requirements. Implementations
of these is beyond the scope of this task. However, a **TestStore** object conforming to the **DataStore** interface is provided for demonstrating its usage as well as allowing us to easily test the system.

### Point of Sale (pos/\*.go)
The PoS module contains the business logic for our application. It includes structures
describing available **Products** and **Transactions** that allows our application
to receive orders, calculate the total cost and return it to the client. This module
also contains the unique products sold by our cafe. In our example, we define a **Coffee**
object (pos/coffee.go) which conforms to the **Product** interface and allows us to
calculate the price of a coffee.
This design allows us to be able to simply add another product to the PoS module by providing its structure, implementing the `GetTotal()` function for the product and ensuring its required
items are available in the DataStore.
A **Controller** object also resides in this module and is responsible for coordinating
the requests from the API to the PoS module for transaction handling and returning the
total cost back to the client.

### API (api/api.go)
The API module provides the JSON/HTTP **APIEndpoint** for our application. It sets up the
desired URL handlers and upon a request, decodes the data and forwards it on to any configured
object conforming to the **APIHandler** interface. In our application, we simply implement
the **APIHandler** interface in **Controller** so that it can receive transactions from the API,
calculate the total and return it back to the **APIEndpoint** for encoding before being sent
back to the client.

## Testing  
This project contains three levels of tests that can be run to demonstrate that the system works as expected and that the calculated costs are correct.

### Level 1 - Unit Testing
Level 1 tests are used to test individual objects (units) to ensure integrity. This is the
foundation of our testing because any faults found in the code at this level can either be exacerbated or hidden by upper layers in the system.
In our project, we define a test that tests the integrity of our datastore so that we are confident that get out precisely what we put into it. To run the datastore test (datastore/datastore_test.go):
`> cd datastore
> go test`

### Level 2 - Component/System Testing
Level 2 tests are used to test components comprised of a number units working as a system. In
our project, we want to ensure that the business logic that calculates the total cost of a coffee
available in our datastore (inventory) is free of any errors. To test the costs of a series
of coffees (pos/coffee_test.go):
`> cd pos
> go test`

### Level 3 - Integration/API Testing
Given that we have high confidence in our code by extensive testing at levels 1 and 2, we can
now test our interactions with our service by testing our API. This will require starting
the application as a service and connecting up a simple client to generate requests.
This is considered beyond the scope of this task.
