package app

import (
	"context"
	"fmt"
	"github.com/elvin-tacirzade/kubernetes-example/pkg/config"
	"github.com/elvin-tacirzade/kubernetes-example/pkg/db"
	"github.com/elvin-tacirzade/kubernetes-example/pkg/handlers"
	"github.com/elvin-tacirzade/kubernetes-example/pkg/repositories"
	"github.com/elvin-tacirzade/kubernetes-example/pkg/services"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type App interface {
	Start()
	Shutdown()
}

type app struct {
	Config *config.Config
	DB     *db.DB
	Echo   *echo.Echo
	WG     sync.WaitGroup
}

func Init() (App, error) {
	conf, confErr := config.Load()
	if confErr != nil {
		return nil, confErr
	}

	database, databaseErr := db.Connect(conf)
	if databaseErr != nil {
		return nil, databaseErr
	}

	e := echo.New()

	return &app{
		Config: conf,
		DB:     database,
		Echo:   e,
	}, nil
}

func (a *app) Start() {
	a.WG.Add(1)
	a.routes()
	go a.runServer()
}

func (a *app) routes() {
	users := initUsers(a.DB.Postgres)
	a.Echo.POST("/api/users", users.Create)
	a.Echo.GET("/api/users", users.Get)
}

func initUsers(db *sqlx.DB) handlers.Users {
	usersRepository := repositories.NewUsers(db)
	usersService := services.NewUsers(usersRepository)
	usersHandler := handlers.NewUsers(usersService)
	return usersHandler
}

func (a *app) runServer() {
	addr := fmt.Sprintf(":%s", a.Config.ADDR)
	log.Println(a.Echo.Start(addr))
	a.WG.Done()
}
func (a *app) Shutdown() {
	done := make(chan struct{})
	signals := make(chan os.Signal)

	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-signals

		echoShutdownErr := a.Echo.Shutdown(context.Background())
		if echoShutdownErr != nil {
			log.Printf("failed to shutdown echo server: %v\n", echoShutdownErr)
		}
		a.WG.Wait()

		log.Printf("received signal: %v", sig)
		done <- struct{}{}
	}()

	<-done
	log.Print("exited the application...")
}
