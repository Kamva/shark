package shark

// ErrorBagMapper is a function type that accept errors of a field, process
// theme, and then return the processed errors.
type ErrorBagMapper func([]string) []string

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

// Map accepts a mapper function and run it in every errors slice and return
// new error bag.
func (b ErrorBag) Map(function ErrorBagMapper) ErrorBag {
	for key, value := range b.errors {
		b.errors[key] = function(value)
	}

	return b
}
