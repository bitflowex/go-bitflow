package types

// Address represents an address.
type Address string

// String returns the string representation of the address.
func (a Address) String() string {
	return string(a)
}
