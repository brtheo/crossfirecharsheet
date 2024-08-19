package server

import (
	_ "github.com/joho/godotenv/autoload"

	"github.com/pocketbase/pocketbase"
)

// type Server struct {
// 	port int
// }

// func NewServer() *http.Server {
// 	port, _ := strconv.Atoi(os.Getenv("PORT"))
// 	NewServer := &Server{
// 		port: port,
// 	}

// 	// Declare Server config
// 	server := &http.Server{
// 		Addr:         fmt.Sprintf(":%d", NewServer.port),
// 		Handler:      NewServer.RegisterRoutes(),
// 		IdleTimeout:  time.Minute,
// 		ReadTimeout:  10 * time.Second,
// 		WriteTimeout: 30 * time.Second,
// 	}

// 	return server
// }

type PBApp struct {
	App *pocketbase.PocketBase
}

func NewPocketBaseApp() *PBApp {
	pbapp := &PBApp{
		App: pocketbase.New(),
	}
	pbapp.App.OnBeforeServe().Add(pbapp.PocketBaseRoutes)
	return pbapp
}
