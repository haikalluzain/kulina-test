# Basic Concept


## Coding

### A. Example of Code Smell in OOP
**Feature Envy**

Feature envy occurs when a method accesses another object's data more than its own. This leads to increased coupling and can hinder maintainability.

**Example**

```Go
type Order struct {
    // Order details
	BasePrice float64
}

type Customer struct {
    // Customer details
}

func (order *Order) CalculateTotal(customer *Customer) float64 {
    // Uses a lot of customer properties and methods
    discount := customer.GetDiscount()
    // Calculate total based on customer info
    total := order.BasePrice - discount
    return total
}
```
**Preventive Action**

Move the functionality related to `Customer` into the `Customer` class itself. Encapsulate the logic within the class that holds the data, promoting better encapsulation and maintainability.
```Go
type Customer struct {
    // Customer details
}

func (customer *Customer) CalculateDiscount() float64 {
    // Calculate discount based on customer info
}

type Order struct {
    // Order details
    Customer Customer
	BasePrice float64
}

func (order *Order) CalculateTotal() float64 {
    // Uses encapsulated customer functionality
    discount := order.Customer.CalculateDiscount()
    total := order.BasePrice - discount
    return total
}

```
### Dependency Injection

*Dependency Injection* is a programming technique that deals with how components or services obtain their dependencies. In other words, it's a technique where the dependencies of a class or a module are injected from the outside rather than being created within the class itself.

In a typical software application, components often depend on other components or services to perform certain tasks. Dependency Injection helps manage these dependencies in a more flexible and decoupled manner. Instead of a class creating its dependencies, these dependencies are "injected" into the class from an external source.

**Why is it important?**
-   **Decoupling:** Dependency injection decouples components, making it easier to replace or change dependencies without modifying the core class.

-   **Testability:** It facilitates easier unit testing by allowing the injection of mock or test implementations of dependencies.

-   **Flexibility:** With Dependency Injection, components are more flexible as dependencies can be changed at runtime or during configuration.



## REST API

You can rename the current file by clicking the file name in the navigation bar or by clicking the **Rename** button in the file explorer.

### A. Do's and Don'ts for Handling Requests

### POST
**Do:** Use POST to create a new resource on the server.
```http request
POST /api/users
Body: { "name": "Haikal", "email": "haikal@gmail.com" }
```
**Don't:** Use POST for idempotent operations. It should not modify the resource with each request.
```http request
POST /api/users/update  (Avoid modifying a resource with each request)
```
### GET

**Do:** Use GET to retrieve information. It should be idempotent and safe, meaning it doesn't modify the server's state.
```http request
GET /api/users
```
**Don't:** Use GET for operations that have side effects or modify server state.
```http request
GET /api/user/123/delete  (Avoid deleting resources with a GET request)
```