package types

import (
	"github.com/google/uuid"
	"github.com/tyrenix/goerr"
)

// ID represents an ID.
type ID uuid.UUID

// String returns the string representation of the id.
func (id ID) String() string {
	return uuid.UUID(id).String()
}

// ParseID parses the id from string.
func ParseID(id string) (ID, error) {
	// parse uuid
	uuid, err := uuid.Parse(id)
	if err != nil {
		return ID(uuid), goerr.New(ErrInvalidID, err)
	}

	// return id and success
	return ID(uuid), nil
}

// Validate validates the id.
func (id ID) Validate() error {
	// if id is empty
	if id.UUID() == uuid.Nil {
		return ErrInvalidID
	}

	// return success
	return nil
}

// UUID returns the uuid.
func (id ID) UUID() uuid.UUID {
	return uuid.UUID(id)
}

// MarshalJSON marshals the id.
func (id ID) MarshalJSON() ([]byte, error) {
	return []byte(`"` + id.String() + `"`), nil
}

// UnmarshalJSON unmarshals the id.
func (id *ID) UnmarshalJSON(data []byte) error {
	// parse uuid
	i, err := ParseID(string(data))
	if err != nil {
		return goerr.New(ErrInvalidID, err)
	}

	// set id
	*id = i

	// return success
	return nil
}

// ID error constants.
var (
	ErrInvalidID = goerr.New("invalid id")
)
