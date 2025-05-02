package workflow

import (
	"context"
	"time"

	"go.uber.org/zap"
)

type Reconciler struct {
	logger *zap.SugaredLogger
	client JobClient
}

func (r *Reconciler) pollWorkspaceStatus(ctx context.Context, workspaces []*Workspace) {
	for {
		for _, ws := range workspaces {
			r.logger.Infof("polling status for workspace %+v", ws)
			status, err := r.client.Status(ctx, ws.ID)
			if err != nil {
				r.logger.Errorf("status check failed, will retry: %v", err)
				continue
			}
			r.logger.Infof("workspace %s status %s", ws.ID, status)
		}
		time.Sleep(2 * time.Second)
	}
}

func (r *Reconciler) syncAll(ctx context.Context, items []*Workspace) {
	r.logger.Infof("Starting operation")
	for _, it := range items {
		r.logger.Infof("Processing event")
		_ = r.client.Sync(ctx, it.ID)
	}
	r.logger.Infof("Finished operation")
}
