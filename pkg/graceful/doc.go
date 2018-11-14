// Package graceful provides a simple wrapper for http.Handler which
// handles graceful shutdown based on registered signals. Server in
// this package closely follows the http.Server struct but can not be
// used as a drop-in replacement.
package graceful
