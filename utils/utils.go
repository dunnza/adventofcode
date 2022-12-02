package utils

import "os"

// Check panics if there is an error. In a real app you would try not to panic
// but for these little puzzles this provides quick feedback while iterating.
func Check(err error) {
	if err != nil {
		panic(err)
	}
}

// ReadFile reads the file contents as a string. You should pass a relative
// path with the starting point being the main.go file. So the general pattern
// is "./adventXXXX/dayXX/<your file>.
func ReadFile(fileName string) string {
	data, err := os.ReadFile(fileName)
	Check(err)
	return string(data)
}
