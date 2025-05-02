package services

import (
	"context"
	"fmt"

	"go.uber.org/zap"
)

type WorkflowService struct {
	logger *zap.SugaredLogger
	store  Store
}

func (s *WorkflowService) Create(ctx context.Context, req *CreateWorkspaceRequest) (*Workspace, error) {
	ws, err := s.store.Save(ctx, req)
	if err != nil {
		s.logger.Errorf("store save failed: %v", err)
		return nil, fmt.Errorf("create workspace: %w", err)
	}
	return ws, nil
}

func (s *WorkflowService) Get(ctx context.Context, id string) (*Workspace, error) {
	ws, err := s.store.Load(ctx, id)
	if err != nil {
		s.logger.Errorf("load workspace failed: %v", err)
		return nil, err
	}
	return ws, nil
}
