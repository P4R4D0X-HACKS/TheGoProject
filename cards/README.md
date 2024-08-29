# Deck of Cards in Go

This project implements a simple deck of cards in Go, showcasing the use of custom types, methods, and slices.

## Overview

The main goal of this project is to demonstrate:

- How to create custom types in Go
- The use of receiver functions to add methods to types
- Basic operations on slices
- Creating and using packages in Go
- Writing unit tests to verify the functionality of Go code


## Usage

The project provides a `deck` package that can be used to create and manipulate a deck of cards. The main file `main.go` demonstrates how to use the package.

To run the project:

```bash

go run main.go

```

To run tests:

```bash

go test

```

## Project Structure

- `deck.go`: Contains the `deck` type and methods associated with it.
- `main.go`: Demonstrates the use of the `deck` type.
- `deck_test.go`: Contains unit tests for the `deck` type.

## Key Features

### 1. Custom Type `deck`

The `deck` type is a custom type that represents a collection of playing cards. It is essentially a slice of strings (`[]string`) but with additional methods that operate on the deck.

### 2. Methods

### `func newDeck() deck`

Creates and returns a new deck of cards.

```go

d := newDeck()

```

### `func (d deck) print()`

Prints all the cards in the deck.

```go

d.print()

```

### `func (d deck) shuffle()`

Shuffles the cards in the deck.

```go

d.shuffle()

```

### `func (d deck) saveToFile(filename string) error`

Saves the deck to a file.

```go

d.saveToFile("my_deck")

```

### `func newDeckFromFile(filename string) deck`

Creates a new deck from a file.

```go

d := newDeckFromFile("my_deck")

```

### `func deal(d deck, handSize int) (deck, deck)`

Deals a hand of cards of the specified size from the deck and returns two decks: the dealt hand and the remaining deck.

```go

hand, remainingDeck := deal(d, 5)

```

## Testing

Unit tests are included in the `deck_test.go` file. These tests cover the creation of a new deck, shuffling, saving to a file, and loading from a file.

To run the tests, execute: