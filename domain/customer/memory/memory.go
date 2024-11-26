// Package memory is a in-memory implementation of Customer Repository

package memory

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
	"github.com/izacarias/online-tavern-ddd/aggregate"
	"github.com/izacarias/online-tavern-ddd/domain/customer"
)

type MemoryRepositry struct {
	customers map[uuid.UUID]aggregate.Customer
	sync.Mutex
}

func New() *MemoryRepositry {
	return &MemoryRepositry{
		customers: make(map[uuid.UUID]aggregate.Customer),
	}
}

func (mr *MemoryRepositry) Get(id uuid.UUID) (aggregate.Customer, error) {
	if customer, ok := mr.customers[id]; ok {
		return customer, nil
	}
	return aggregate.Customer{}, customer.ErrCustomerNotFound
}

func (mr *MemoryRepositry) Add(c aggregate.Customer) error {
	if mr.customers == nil {
		mr.Lock()
		mr.customers = make(map[uuid.UUID]aggregate.Customer)
		mr.Unlock()
	}
	// Check if me customer is already in the repository
	if _, ok := mr.customers[c.GetID()]; ok {
		return fmt.Errorf("customer already in repository :%w", customer.ErrFaileToAddCustomer)
	}
	mr.Lock()
	mr.customers[c.GetID()] = c
	mr.Unlock()
	return nil
}

func (mr *MemoryRepositry) Update(c aggregate.Customer) error {
	if _, ok := mr.customers[c.GetID()]; !ok {
		return fmt.Errorf("customer does not exist: %w", customer.ErrFailedToUpdateCustomer)
	}
	mr.Lock()
	mr.customers[c.GetID()] = c
	mr.Unlock()
	return nil
}
