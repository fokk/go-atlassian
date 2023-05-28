package assets

import (
	"context"
	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
)

// ObjectAssetConnector represents the Assets objects.
// Use it to search, get, create, delete, and change objects.
type ObjectAssetConnector interface {

	// Get loads one object.
	//
	// GET /jsm/assets/workspace/{workspaceId}/v1/object/{id}
	Get(ctx context.Context, workspaceID, objectID string) (*models.ObjectAssetScheme, *models.ResponseScheme, error)

	// Update updates an existing object in Assets.
	//
	// PUT /jsm/assets/workspace/{workspaceId}/v1/object/{id}
	Update(ctx context.Context, workspaceID, objectID string, payload *models.ObjectAssetPayloadScheme) (*models.ObjectAssetScheme, *models.ResponseScheme, error)

	// Delete deletes the referenced object
	//
	// DELETE /jsm/assets/workspace/{workspaceId}/v1/object/{id}
	Delete(ctx context.Context, workspaceID, objectID string) (*models.ResponseScheme, error)

	// Attributes list all attributes for the given object.
	//
	// GET /jsm/assets/workspace/{workspaceId}/v1/object/{id}/attributes
	Attributes(ctx context.Context, workspaceID, objectID string) ([]*models.ObjectAssetAttributeScheme, *models.ResponseScheme, error)

	// History retrieves the history entries for this object.
	//
	// GET /jsm/assets/workspace/{workspaceId}/v1/object/{id}/history
	History(ctx context.Context, workspaceID, objectID string, ascOrder bool) ([]*models.ObjectAssetHistoryScheme, *models.ResponseScheme, error)

	// References finds all references for an object.
	//
	// GET /jsm/assets/workspace/{workspaceId}/v1/object/{id}/referenceinfo
	References(ctx context.Context, workspaceID, objectID string) ([]*models.ObjectAssetReferenceTypeInfoScheme, *models.ResponseScheme, error)

	// Create creates a new object in Assets.
	//
	// POST /jsm/assets/workspace/{workspaceId}/v1/object/create
	Create(ctx context.Context, workspaceID string, payload *models.ObjectAssetPayloadScheme) (*models.ObjectAssetScheme, *models.ResponseScheme, error)
}