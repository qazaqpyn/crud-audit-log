package server

import (
	"context"

	audit "github.com/qazaqpyn/crud-audit-log/pkg/domain"
)

type AuditService interface {
	Insert(ctx context.Context, req *audit.LogRequest) error
}

type AuditServer struct {
	service AuditService
	audit.UnimplementedAuditServiceServer
}

func NewAuditServer(service AuditService) *AuditServer {
	return &AuditServer{
		service: service,
	}
}

func (h *AuditServer) Log(ctx context.Context, req *audit.LogRequest) (*audit.Empty, error) {
	err := h.service.Insert(ctx, req)

	return &audit.Empty{}, err
}
