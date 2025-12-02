package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"strconv"
	"time"

	"example.com/fiber-pg-rmq-jwt/internal/db"
	"example.com/fiber-pg-rmq-jwt/internal/mq"
	"example.com/fiber-pg-rmq-jwt/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env if present
	_ = godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Initialize Postgres
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL is required")
	}
	pg, err := db.NewPool(context.Background(), dbURL)
	if err != nil {
		log.Fatalf("pg connect: %v", err)
	}
	defer pg.Close()

	// Auto-create a minimal users table if not exists
	if err := db.Migrate(context.Background(), pg); err != nil {
		log.Fatalf("migrate: %v", err)
	}

	// Initialize RabbitMQ
	rmqURL := os.Getenv("RABBITMQ_URL")
	if rmqURL == "" {
		log.Fatal("RABBITMQ_URL is required")
	}
	rmq, err := mq.New(rmqURL)
	if err != nil {
		log.Fatalf("rabbitmq connect: %v", err)
	}
	defer rmq.Close()

	// Start consumer (optional)
	if err := rmq.StartConsumer("example_queue", func(body []byte) {
		log.Printf("consumer received: %s", string(body))
	}); err != nil {
		log.Fatalf("start consumer: %v", err)
	}

	// Build app and inject dependencies
	app := fiber.New()
	h := handlers.NewHandler(pg, rmq)

	h.RegisterRoutes(app)

	// graceful shutdown
	go func() {
		if err := app.Listen(":" + port); err != nil {
			log.Printf("listen error: %v", err)
		}
	}()
	log.Printf("server listening on :%s", port)

	// Wait for interrupt
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_ = app.ShutdownWithContext(ctx)
}