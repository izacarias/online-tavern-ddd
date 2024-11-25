package valueobject

import (
	"time"

	"github.com/google/uuid"
)

// Transaction is a value object for two reasons
// - It has no identifier
// - It cannot change
type Transaction struct {
	amount    int       // the amount of the transaction
	from      uuid.UUID // from which person
	to        uuid.UUID // to which person
	createdAt time.Time // When the transaction occured
}
