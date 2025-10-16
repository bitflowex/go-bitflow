package types

import "github.com/tyrenix/goerr"

// Username represents a username.
type Username string

// Validate validates the username.
func (u Username) Validate() error {
	// if username is empty
	if u == "" {
		return ErrEmptyUsername
	}

	// return success
	return nil
}

// User represents a user.
type User struct {
	ID       *ID       `json:"id"`
	Username *Username `json:"username"`
}

// Username error constants.
var (
	ErrEmptyUsername = goerr.New("empty username")
)
