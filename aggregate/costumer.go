// This package hold our aggregates that combine many
// entitities into a full object

package aggregate

import (
	"github.com/izacarias/online-tavern-ddd/entity"
	"github.com/izacarias/online-tavern-ddd/valueobject"
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
