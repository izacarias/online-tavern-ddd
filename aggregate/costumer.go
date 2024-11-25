// This package hold our aggregates that combine many
// entitities into a full object

package aggregate

import (
	"errors"

	"github.com/google/uuid"
	"github.com/izacarias/online-tavern-ddd/entity"
	"github.com/izacarias/online-tavern-ddd/valueobject"
)

var (
	ErrInvalidPerson = errors.New("A customer has to have a valid name")
)

type Customer struct {
	// person is the root entity of Customer
	// which means person.ID is the main identifier fo the customer
	person *entity.Person
	// products that the person is buying
	products []*entity.Item
	// Transaction is not a pointer because it does not change
	transactions []valueobject.Transaction
}

// This is an ordinary factory design pattern implementation
// The funcion creates Customers handling internal logic
// and recives the name as parameter
// Here is only validates that name is not empty
func NewCustomer(name string) (Customer, error) {
	if name == "" {
		return Customer{}, ErrInvalidPerson
	}

	person := &entity.Person{
		Name: name,
		ID:   uuid.New(),
	}

	return Customer{
		person:       person,
		products:     make([]*entity.Item, 0),
		transactions: make([]valueobject.Transaction, 0),
	}, nil
}
