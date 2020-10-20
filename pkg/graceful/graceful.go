package graceful

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// LogFunc can be set on the server to customize the message printed when the
// server is shutting down.
type LogFunc func(msg string, args ...interface{})

// NewServer creates a wrapper around the given handler.
func NewServer(handler http.Handler, timeout time.Duration, signals ...os.Signal) *Server {
	gss := &Server{}
	gss.server = &http.Server{Handler: handler}
	gss.signals = signals
	gss.Log = log.Printf
	gss.timeout = timeout
	return gss
}

// Server is a wrapper around an http handler. It provides methods
// to start the server with graceful-shutdown enabled.
type Server struct {
	Addr string
	Log  LogFunc

	server   *http.Server
	signals  []os.Signal
	timeout  time.Duration
	startErr error
}

// Serve starts the http listener with the registered http.Handler and
// then blocks until a interrupt signal is received.
func (gss *Server) Serve(l net.Listener) error {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		if err := gss.server.Serve(l); err != nil {
			gss.startErr = err
			cancel()
		}
	}()
	return gss.waitForInterrupt(ctx)
}

// ServeTLS starts the http listener with the registered http.Handler and
// then blocks until a interrupt signal is received.
func (gss *Server) ServeTLS(l net.Listener, certFile, keyFile string) error {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		if err := gss.server.ServeTLS(l, certFile, keyFile); err != nil {
			gss.startErr = err
			cancel()
		}
	}()
	return gss.waitForInterrupt(ctx)
}

// ListenAndServe serves the requests on a listener bound to interface
// specified by Addr
func (gss *Server) ListenAndServe() error {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		gss.server.Addr = gss.Addr
		if err := gss.server.ListenAndServe(); err != http.ErrServerClosed {
			gss.startErr = err
			cancel()
		}
	}()
	return gss.waitForInterrupt(ctx)
}

// ListenAndServeTLS serves the requests on a listener bound to interface
// specified by Addr
func (gss *Server) ListenAndServeTLS(certFile, keyFile string) error {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		if err := gss.server.ListenAndServeTLS(certFile, keyFile); err != http.ErrServerClosed {
			gss.startErr = err
			cancel()
		}
	}()
	return gss.waitForInterrupt(ctx)
}

func (gss *Server) waitForInterrupt(ctx context.Context) error {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, gss.signals...)

	select {
	case sig := <-sigCh:
		if gss.Log != nil {
			gss.Log("shutting down (signal=%s)...", sig)
		}
		break

	case <-ctx.Done():
		return gss.startErr
	}

	return gss.shutdown()
}

func (gss *Server) shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), gss.timeout)
	defer cancel()
	return gss.server.Shutdown(ctx)
}
