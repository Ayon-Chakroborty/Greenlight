package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func (app *application) serve() error {

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.config.port),
		Handler:      app.routes(),
		IdleTimeout:  5 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(app.logger.Handler(), slog.LevelError),
	}

	shutdownError := make(chan error)

	// USE CTRL \ to skip gracefull shutdown
	// go routine with gracefull shutdown of server
	go func ()  {
		quit := make(chan os.Signal, 1)

		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		s := <-quit

		app.logger.Info("shutting down server", "signal", s.String())

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		shutdownError <- srv.Shutdown(ctx) 
	}()
	

	app.logger.Info("starting server", "addr", srv.Addr, "env", app.config.env)

	err := srv.ListenAndServe() // start server and wait for shutdown
	if !errors.Is(err, http.ErrServerClosed){ // if graceful shutdown occured successfully then it will return ErrServerClosed.
		return err							  // So we check for other errors because that means something went wrong
	}

	err = <-shutdownError // check for errors again
	if err != nil{
		return err
	}

	app.logger.Info("stopped server", "addr", srv.Addr) // successful graceful shutdown

	return nil
}
