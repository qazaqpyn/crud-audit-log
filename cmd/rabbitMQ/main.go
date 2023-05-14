package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/qazaqpyn/crud-audit-log/internal/config"
	"github.com/qazaqpyn/crud-audit-log/internal/repository"
	"github.com/qazaqpyn/crud-audit-log/internal/repository/mongo"
	server "github.com/qazaqpyn/crud-audit-log/internal/server/rabbitMQ"
	service "github.com/qazaqpyn/crud-audit-log/internal/service/rabbitMQ"
)

func main() {
	ctx := context.Background()

	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	db, err := mongo.NewMongo(cfg.DB.Database, cfg.DB.URI)
	if err != nil {
		log.Fatal(err)
	}

	auditRepo := repository.NewAudit(db)
	auditService := service.NewAuditService(auditRepo)

	srv := server.NewClient(auditService)

	if err = srv.Listening(cfg.Rabbit.URI); err != nil {
		log.Fatal(err)
	}

	defer func(srv *server.Client) {
		if err := srv.Close(); err != nil {
			log.Fatal(err)
		}
	}(srv)

	go func() {
		if err := srv.Serve(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	log.Println("server started")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Println("server stopped")
}
