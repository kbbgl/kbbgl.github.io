# Slices and Loops

Go has two data structures:

- `Array` is a fixed length list.
- `Slice` is an array that can grow or shrink. All elements inside the Slice needs to have the same data type.

## Declare a Slice

```go
cards := []string{}
// e.g.
cards := []string{newCard(), "Ten of Leafs", "Ace of Diamonds"}
```

## Add Element to Slice

```go
cards := []string{newCard(), "Ten of Leafs", "Ace of Diamonds"}
cards = append(cards, "Six of Spades")
```

## Iterating over Slice Elements

```go
cards := []string{newCard(), "Ten of Leafs", "Ace of Diamonds"}
for i, card := range cards {
  fmt.Println(i, card)
}
```

If we want to ignore the indices, we name them `_`:

```go
for _, suit := range cardSuits {
  for _, value := range cardValues {
    cards = append(cards, value+" of "+suit)
  }
}
```

## Slice Structure

When we declare a slice, Go initializes 2 data structures:

1) A new slice containing **reference types**:

- `length` (how many elements are in the array),
- `capacity` (how many elements are allowed in array)
- a pointer to the memory address of the first element of array.

2) An array with the elements and values

When we pass a slice as an argument into the function, Go:

1) Creates a copy of the slice into a different memory address.

2) The copied slice Pointer to the memory address of the first element of the array is the same as the original.
