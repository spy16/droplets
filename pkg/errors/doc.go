// Package errors provides common error definitions and tools for dealing
// with errors. The API of this package is drop-in replacement for standard
// errors package except for one significant difference:
//   Since Error type used in this package embeds map inside, errors created
// by this package are not comparable and  hence cannot  be used to create
// Sentinel Errors.
package errors
