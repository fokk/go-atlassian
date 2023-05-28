package assets

import (
	"context"
	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
)

// AQLAssetConnector represents the assets AQL search endpoints.
// Use it to execute AQL searches.
type AQLAssetConnector interface {

	// Filter retrieves a list of objects based on an AQL. Note that the preferred endpoint is /aql.
	//
	// POST /jsm/assets/workspace/{workspaceId}/v1/object/navlist/aql
	Filter(ctx context.Context, workspaceID string, payload *models.AQLAssetSearchParamsScheme) (*models.ObjectAssetPageScheme, *models.ResponseScheme, error)
}