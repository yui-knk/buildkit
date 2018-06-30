// +build !windows

package shellutil

// EqualEnvKeys compare two strings and returns true if they are equal. On
// Windows this comparison is case insensitive.
func EqualEnvKeys(from, to string) bool {
	return from == to
}
