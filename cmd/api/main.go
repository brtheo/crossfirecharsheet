package main

import (
	"crossfirecharsheet/internal/server"
	"fmt"
)

func main() {
	app := server.NewPocketBaseApp()
	// server := server.NewServer()

	// server.ListenAndServe()
	if err := app.App.Start(); err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
