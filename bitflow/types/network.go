package types

import "github.com/tyrenix/goerr"

// NetworkCode represents a network code.
type NetworkCode string

// Network codes.
const (
	NetworkTON  NetworkCode = "TON"
	NetworkTRON NetworkCode = "TRON"
)

// String returns the string representation of the network code.
func (n NetworkCode) String() string {
	return string(n)
}

// Validate validates the network code.
func (n NetworkCode) Validate() error {
	switch n {
	case NetworkTON, NetworkTRON:
		return nil
	default:
		return ErrInvalidNetworkCode
	}
}

// Network code error constants.
var (
	ErrInvalidNetworkCode = goerr.New("invalid network code")
)
