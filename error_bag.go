package shark

// ErrorBag is a error container
type ErrorBag struct {
	errors map[string][]string
}

// Append add new error for key
func (b ErrorBag) Append(key string, value string) {
	b.errors[key] = append(b.errors[key], value)
}

// GetErrors returns all errors of all keys
func (b ErrorBag) GetErrors() map[string][]string {
	return b.errors
}
