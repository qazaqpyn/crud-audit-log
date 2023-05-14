package main

import (
	"fmt"
	"log"
	"time"

	"github.com/qazaqpyn/crud-audit-log/internal/config"
	"github.com/qazaqpyn/crud-audit-log/internal/repository"
	"github.com/qazaqpyn/crud-audit-log/internal/repository/mongo"
	server "github.com/qazaqpyn/crud-audit-log/internal/server/grpc"
	service "github.com/qazaqpyn/crud-audit-log/internal/service/grpc"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}
	db, err := mongo.NewMongo(cfg.DB.Database, cfg.DB.URI)
	if err != nil {
		log.Fatal(err)
	}

	auditRepo := repository.NewAudit(db)
	auditService := service.NewAudit(auditRepo)
	auditServer := server.NewAuditServer(auditService)
	srv := server.New(auditServer)

	fmt.Println("SERVER STARTED", time.Now())

	if err := srv.ListenAndServe(cfg.Server.Port); err != nil {
		log.Fatal(err)
	}
}
