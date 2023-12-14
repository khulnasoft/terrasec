package version

import "fmt"

// Terrasec The Terrasec version
const Terrasec = "1.18.8"

// Get returns the terrasec version
func Get() string {
	return fmt.Sprintf("v%s", Terrasec)
}

// GetNumeric returns the numeric terrasec version
func GetNumeric() string {
	return Terrasec
}