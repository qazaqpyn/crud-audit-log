package service

import (
	"context"
	"log"

	audit "github.com/qazaqpyn/crud-audit-log/pkg/domain"
	"google.golang.org/protobuf/proto"
)

type AuditRepo interface {
	Insert(ctx context.Context, msg audit.LogItem) error
}

type AuditService struct {
	repo AuditRepo
}

func NewAuditService(repo AuditRepo) *AuditService {
	return &AuditService{
		repo: repo,
	}
}

func (s *AuditService) Insert(ctx context.Context, msg []byte) error {
	ans := &audit.LogRequest{}

	err := proto.Unmarshal(msg, ans)
	if err != nil {
		log.Fatal(err)
	}
	item := audit.LogItem{
		Action:    ans.GetAction().String(),
		Entity:    ans.GetEntity().String(),
		EntityID:  ans.GetEntityId(),
		Timestamp: ans.GetTimestamp().AsTime(),
	}

	return s.repo.Insert(ctx, item)
}
