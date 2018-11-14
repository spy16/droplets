package graceful

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
)

// LogFunc can be set on the server to customize the message printed when the
// server is shutting down.
type LogFunc func(msg string, args ...interface{})

// NewServer creates a wrapper around the given handler.
func NewServer(handler http.Handler, signals ...os.Signal) *Server {
	gss := &Server{}
	gss.server = &http.Server{
		Handler: handler,
	}
	gss.signals = signals
	gss.Log = log.Printf
	return gss
}

// Server is a wrapper around an http handler. It provides methods
// to start the server with graceful-shutdown enabled.
type Server struct {
	Addr string
	Log  LogFunc

	server  *http.Server
	signals []os.Signal
	err     error
}

// Serve starts the http listener with the registered http.Handler and
// then blocks until a interrupt signal is received.
func (gss *Server) Serve(l net.Listener) error {
	go func() {
		if err := gss.server.Serve(l); err != nil {
			gss.err = err
		}
	}()
	return gss.waitForShutdown()
}

// ServeTLS starts the http listener with the registered http.Handler and
// then blocks until a interrupt signal is received.
func (gss *Server) ServeTLS(l net.Listener, certFile, keyFile string) error {
	go func() {
		if err := gss.server.ServeTLS(l, certFile, keyFile); err != nil {
			gss.err = err
		}
	}()
	return gss.waitForShutdown()
}

// ListenAndServe serves the requests on a listener bound to interface
// specified by Addr
func (gss *Server) ListenAndServe() error {
	go func() {
		gss.server.Addr = gss.Addr
		if err := gss.server.ListenAndServe(); err != http.ErrServerClosed {
			gss.err = err
		}
	}()
	return gss.waitForShutdown()
}

// ListenAndServeTLS serves the requests on a listener bound to interface
// specified by Addr
func (gss *Server) ListenAndServeTLS(certFile, keyFile string) error {
	go func() {
		if err := gss.server.ListenAndServeTLS(certFile, keyFile); err != http.ErrServerClosed {
			gss.err = err
		}
	}()
	return gss.waitForShutdown()
}

func (gss *Server) waitForShutdown() error {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, gss.signals...)
	_ = <-sig

	if gss.Log != nil {
		gss.Log("received interrupt. shutting down..")
	}
	gss.server.Shutdown(context.Background())
	return gss.err
}
