package httpserver

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	gorillaHandlers "github.com/gorilla/handlers"
	httputils "github.com/khulnasoft/terrasec/pkg/utils/http"
	"go.uber.org/zap"

	"github.com/gorilla/mux"
	"github.com/khulnasoft/terrasec/pkg/logging"
)

// Start initializes api routes and starts http server
func Start(port, certFile, privateKeyFile string) {

	// create a new API server
	server := NewAPIServer()

	// get all routes
	routes := server.Routes()

	// get port
	if port == GatewayDefaultPort && os.Getenv(TerrasecServerPort) != "" {
		port = os.Getenv(TerrasecServerPort)
	}

	// register routes and start the http server
	server.start(routes, port, certFile, privateKeyFile)
}

// start http server
func (g *APIServer) start(routes []*Route, port, certFile, privateKeyFile string) {

	var (
		logger = logging.GetDefaultLogger() // new logger
		err    error
		router = mux.NewRouter() // new router
	)

	logWriter := getLogWriter(logger)

	logger.Info("registering routes...")

	if privateKeyFile != "" || certFile != "" {
		logger.Debugf("certfile is %s, privateKeyFile is %s", certFile, privateKeyFile)

		if err := g.validateFiles(privateKeyFile, certFile); err != nil {
			logger.Fatal(err)
		}
	}

	// register all routes
	for _, v := range routes {
		logger.Info("Route ", v.verb, " - ", v.path)
		handler := gorillaHandlers.LoggingHandler(logWriter, http.HandlerFunc(v.fn))
		router.Methods(v.verb).Path(v.path).Handler(handler)
	}

	router.NotFoundHandler = gorillaHandlers.LoggingHandler(logWriter, http.HandlerFunc(httputils.NotFound))
	router.MethodNotAllowedHandler = gorillaHandlers.LoggingHandler(logWriter, http.HandlerFunc(httputils.NotAllowed))

	// Add a route for all static templates / assets. Currently used for the Webhook logs views
	// go/terrasec/asset is the path where the assets files are located inside the docker container
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("/go/terrasec/assets"))))

	// start http server
	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	message := make(chan string)
	go func() {
		var err error
		if certFile != "" && privateKeyFile != "" {
			// In case a certificate file is specified, the server support TLS
			message <- "https server listening at port %v"
			err = server.ListenAndServeTLS(certFile, privateKeyFile)
		} else {
			message <- "http server listening at port %v"
			err = server.ListenAndServe()
		}
		if err != nil && err != http.ErrServerClosed {
			logger.Fatal(err)
		}
	}()

	logger.Infof(<-message, port)

	close(message)

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	// try to stop the server gracefully with default 5 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = server.Shutdown(ctx)
	if err != nil {
		logger.Fatalf("server failed to exit gracefully. error: '%v'", err)
	}
	logger.Info("server exiting gracefully")
}

type logWriter struct {
	logger *zap.SugaredLogger
}

// GetLogWriter creates and fetches a LogWriter object
func getLogWriter(logger *zap.SugaredLogger) logWriter {
	return logWriter{
		logger: logger,
	}
}

// Write writes the byte array to the logger object
func (l logWriter) Write(p []byte) (n int, err error) {
	l.logger.Info(string(p))
	return len(p), nil
}
