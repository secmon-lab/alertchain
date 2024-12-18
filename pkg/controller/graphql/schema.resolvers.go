package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"
	"fmt"

	"github.com/secmon-lab/alertchain/pkg/domain/model"
	"github.com/secmon-lab/alertchain/pkg/domain/types"
	"github.com/secmon-lab/alertchain/pkg/utils"
)

// Workflows is the resolver for the workflows field.
func (r *queryResolver) Workflows(ctx context.Context, offset *int, limit *int) ([]*model.WorkflowRecord, error) {
	results, err := r.svc.Workflow.Get(ctx, offset, limit)
	if err != nil {
		return nil, err
	}

	return utils.ToPtrSlice(results), nil
}

// Workflow is the resolver for the workflow field.
func (r *queryResolver) Workflow(ctx context.Context, id string) (*model.WorkflowRecord, error) {
	return r.svc.Workflow.Lookup(ctx, types.WorkflowID(id))
}

// Actions is the resolver for the actions field.
func (r *workflowRecordResolver) Actions(ctx context.Context, obj *model.WorkflowRecord) ([]*model.ActionRecord, error) {
	panic(fmt.Errorf("not implemented: Actions - actions"))
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// WorkflowRecord returns WorkflowRecordResolver implementation.
func (r *Resolver) WorkflowRecord() WorkflowRecordResolver { return &workflowRecordResolver{r} }

type queryResolver struct{ *Resolver }
type workflowRecordResolver struct{ *Resolver }
