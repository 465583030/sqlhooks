package sqlhooks

import "database/sql"

type ExampleHooks struct{}

func ExampleNewDriver() {
	NewDriver("mysql", nil)
}

func ExampleRegister() {
	// Register the driver under `hooked-mysql` name
	Register("hooked-mysql", NewDriver("mysql", nil))

	// Open a connection
	sql.Open("hooked-mysql", "/db")
}
