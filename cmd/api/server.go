package main

import (
	"fmt"
	"net/http"
)

func (app *application) serve() error {
	fmt.Println("Logged in")
	srv := &http.Server{
		Handler: nil,
		Addr:    ":4000",
	}
	err := srv.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}
