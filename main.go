package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	opentracing "github.com/opentracing/opentracing-go"
	"log"
	"os"
	"os/signal"
	"syscall"
	"vandyahmad24/maxsol/app/config"
	"vandyahmad24/maxsol/app/router"
	"vandyahmad24/maxsol/app/util"
	"vandyahmad24/maxsol/migration"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("Failed to load config: %s", err)
		os.Exit(1)
	}

	fmt.Println("Maxsol Test By Vandy Ahmad")
	if cfg.Rest.Port == 0 {
		log.Fatal("Port env is requeird")
	}

	fmt.Println("Menjalankan Migration")
	err = migration.Up()
	if err != nil {
		log.Fatal(err)
	}
	tracer, closer := util.Init("Cake Service")
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	app := fiber.New(fiber.Config{
		BodyLimit: 8 * 1024 * 1024, // this is the default limit of 4MB
	})
	app.Use(recover.New())
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON("welcome to Cake Service")
	})

	app.Get("/panic", func(c *fiber.Ctx) error {
		panic("handle panic")
	})

	db := config.InitDb(cfg)
	router.CakeRouter(app, db)
	router.OrderRouter(app, db)

	go func() {
		app.Listen(fmt.Sprintf(":%d", cfg.Rest.Port))
	}()
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	signal := <-c
	log.Fatalf("Server Mati : %v\n", signal.String())

}
