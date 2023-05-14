package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/qazaqpyn/crud-audit-log/internal/config"
	"github.com/qazaqpyn/crud-audit-log/internal/repository"
	"github.com/qazaqpyn/crud-audit-log/internal/repository/mongo"
	rabbitmq "github.com/qazaqpyn/crud-audit-log/internal/server/rabbitMQ"
	"github.com/qazaqpyn/crud-audit-log/internal/service"
	audit "github.com/qazaqpyn/crud-audit-log/pkg/domain"
	"google.golang.org/protobuf/proto"
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

	rabbit, err := rabbitmq.NewClient(cfg.Rabbit.URI)
	if err != nil {
		log.Fatal(err)
	}

	msgs, err := rabbit.StartListening()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("SERVER STARTED", time.Now())

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			ans := &audit.LogRequest{}

			err := proto.Unmarshal(d.Body, ans)
			if err != nil {
				log.Fatal(err)
			}
			err = auditService.Insert(context.Background(), ans)
			if err != nil {
				log.Fatal(err)
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
