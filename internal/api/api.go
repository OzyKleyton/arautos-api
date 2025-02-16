package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/OzyKleyton/arautos-api/config"
	"github.com/OzyKleyton/arautos-api/config/db"
	"github.com/OzyKleyton/arautos-api/internal/api/handler"
	"github.com/OzyKleyton/arautos-api/internal/api/router"
	"github.com/OzyKleyton/arautos-api/internal/model"
	"github.com/OzyKleyton/arautos-api/internal/repository"
	"github.com/OzyKleyton/arautos-api/internal/service"
)

func Run(host, port string) error {
	address := fmt.Sprintf("%s:%s", host, port)
	log.Println("Listen app in port ", address)

	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
		Prefork:     config.GetConfig().Prefork,
		ProxyHeader: fiber.HeaderXForwardedFor,
	})

	db, err := db.ConnectDB(config.GetConfig().DBURL)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	db = db.WithContext(ctx)

	if err := db.AutoMigrate(
		&model.User{},
	); err != nil {
		return err
	}

	userRepo := repository.NewUserRepository(db)

	userService := service.NewUserService(userRepo)

	userHandler := handler.NewUserHandler(userService)

	router.SetupRouter(app, userHandler.Routes())

	c := make(chan os.Signal, 1)
	errc := make(chan error, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	go func() {
		<-c
		fmt.Println("Gracefully shutting down...")
		cancel()
		errc <-app.Shutdown()
	}()

	if err := app.Listen(address); err != nil {
		return err
	}

	err = <-errc

	return err
}