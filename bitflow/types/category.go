package types

// CategoryCode represents a category code.
type CategoryCode string

// String returns the string representation of the category code.
func (c CategoryCode) String() string {
	return string(c)
}
