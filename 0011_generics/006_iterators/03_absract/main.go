package main

// As a matter of convention, we encourage all container types to provide an All method that returns an iterator,
// so that programmers don’t have to remember whether to range over All directly or whether to call
// All to get a value they can range over. They can always do the latter

// The iterator method on a collection type is conventionally named All,
// because it iterates a sequence of all the values in the collection.
// For a type containing multiple possible sequences, the iterator's name can indicate which sequence is being provided
