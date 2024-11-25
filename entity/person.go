package entity

import (
	"github.com/google/uuid"
)

// Represent a person in all domains
type Person struct {
	// ID is the identifier of the entity
	ID   uuid.UUID
	Name string
	Age  int
}
