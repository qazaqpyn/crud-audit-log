package repository

import (
	"context"

	audit "github.com/qazaqpyn/crud-audit-log/pkg/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuditRepo struct {
	db *mongo.Database
}

func NewAudit(db *mongo.Database) *AuditRepo {
	return &AuditRepo{
		db: db,
	}
}

func (r *AuditRepo) Insert(ctx context.Context, item audit.LogItem) error {
	_, err := r.db.Collection("logs").InsertOne(ctx, item)

	return err
}
