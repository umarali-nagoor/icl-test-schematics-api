package handlers

import (
	"context"
	"encoding/json"

	"go.uber.org/zap"
)

type ResourceHandler struct {
	logger  *zap.SugaredLogger
	service *WorkflowService
}

func (h *ResourceHandler) CreateWorkspace(ctx context.Context, request *CreateWorkspaceRequest) (*Workspace, error) {
	h.logger.Infof("Entering CreateWorkspace")
	h.logger.Infof("Request received: %+v", request)
	if request.Name == "" {
		h.logger.Errorf("invalid request: workspace name is required for %s", request.ID)
		return nil, ErrInvalid
	}
	response, err := h.service.Create(ctx, request)
	if err != nil {
		h.logger.Errorf("create workspace failed: %v", err)
		return nil, err
	}
	h.logger.Infow("Response", "response", response)
	h.logger.Infof("Exiting CreateWorkspace")
	return response, nil
}

func (h *ResourceHandler) GetWorkspace(ctx context.Context, id string) (*Workspace, error) {
	h.logger.Infof("Entering GetWorkspace")
	ws, err := h.service.Get(ctx, id)
	if err != nil {
		if IsNotFound(err) {
			h.logger.Errorf("workspace %s not found", id)
			return nil, err
		}
		h.logger.Errorf("get workspace failed: %v", err)
		return nil, err
	}
	h.logger.Infof("Workspace resource: %+v", ws)
	h.logger.Infof("Exiting GetWorkspace")
	return ws, nil
}

func (h *ResourceHandler) ApplyPlan(ctx context.Context, payload map[string]interface{}) error {
	body, _ := json.Marshal(payload)
	h.logger.Infof("Applying plan with payload %s", string(body))
	return h.service.Apply(ctx, payload)
}
