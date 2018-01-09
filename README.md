# cabimarkt

![Toronto's Flower Market](https://static1.squarespace.com/static/5657197fe4b0f0c1a00f9f7e/t/5701b1e54d088e77810684de/1459728872953/)

It's a small coding exercise in Golang whose main goal is trying to emulate a online shop at a very, very and very small scale. It does not feature any persistance system, webserver or message queue. Instead, it features concurrency and desing patterns.

## Solution

In this section, you can find enumerated the most important types I've added to  implement the exercise.

### Checkout

A checkout represents an order, which is mainly made out of many `CheckoutLines`. This type contains methods to add into the "shopping cart" as well as functions to start a concurrent computation of the final cost of the order after promotions.   

### Catalog

It acts as the source of truth for any kind of product or promotions information. Think of it as a replacement in memory for any kind of database connection. Futhermore, it also implements CRUD operations and a Singelton pattern, to avoid inconsistencies.

### Item

Not much to comment of this one, it's just a DTO or naked struct without any logic that holds product's information such as their codes, prices and visibility.

### Handler

It finds the best `Promotion` for a given item under certain conditions. 

### Promotion

An interface for promotions, so therefore different types of promotions can be implemented in a standarized way.

### Subtotal

Another DTO which contains the result the optimal cost found by a `Handler` for a given item.

## Possible improvements

Some improvements I thought of:

- Bundle promotions.
- Generating shopping tickets.
- `main` function to read user input.
