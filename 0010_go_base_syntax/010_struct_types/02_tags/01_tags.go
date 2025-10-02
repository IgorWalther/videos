package main

// A field declaration may be followed by an optional string literal tag,
// which becomes an attribute for all the fields in the corresponding field declaration.
// An empty tag string is equivalent to an absent tag. The tags are made visible through a
// reflection interface and take part in type identity for structs but are otherwise ignored.

type T struct {
	microsec  uint64 `json:"1"`
	serverIP6 uint64 `protobuf:"2"`
}
