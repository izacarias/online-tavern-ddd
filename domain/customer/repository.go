// Here we handle the logic for storing the the customer aggregate
// It basically builds upon the Repository Design Pattern

package customer

import (
	"errors"

	"github.com/google/uuid"
	"github.com/izacarias/online-tavern-ddd/aggregate"
)

var (
	ErrCustomerNotFound       = errors.New("The customer was not found")
	ErrFaileToAddCustomer     = errors.New("Failed to add the customer")
	ErrFailedToUpdateCustomer = errors.New("Failed to update the customer")
)

type CustomerRepository interface {
	Get(uuid.UUID) (aggregate.Customer, error)
	Add(aggregate.Customer) error
	Update(aggregate.Customer) error
}
