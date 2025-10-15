package types

// ResponseStatus represents a response status.
type ResponseStatus string

// Response statuses.
const (
	ResponseSuccess ResponseStatus = "success"
	ResponseError   ResponseStatus = "error"
)

// Represents a response.
type Response[T any] struct {
	StatusCode int            `json:"status_code"`
	Status     ResponseStatus `json:"status"`
	Message    string         `json:"message,omitempty"`
	Error      string         `json:"error,omitempty"`
	Data       T              `json:"data,omitempty"`
}
