package service

import (
	"context"

	audit "github.com/qazaqpyn/crud-audit-log/pkg/domain"
)

type AuditRepo interface {
	Insert(ctx context.Context, item audit.LogItem) error
}

type AuditService struct {
	repo AuditRepo
}

func NewAudit(repo AuditRepo) *AuditService {
	return &AuditService{
		repo: repo,
	}
}

func (s *AuditService) Insert(ctx context.Context, req *audit.LogRequest) error {
	item := audit.LogItem{
		Action:    req.GetAction().String(),
		Entity:    req.GetEntity().String(),
		EntityID:  req.GetEntityId(),
		Timestamp: req.GetTimestamp().AsTime(),
	}

	return s.repo.Insert(ctx, item)
}
