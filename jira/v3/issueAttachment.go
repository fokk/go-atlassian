package v3

import (
	"bytes"
	"context"
	"fmt"
	models2 "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"io"
	"mime/multipart"
	"net/http"
)

type AttachmentService struct{ client *Client }

// Settings returns the attachment settings, that is, whether attachments are enabled and the maximum attachment size allowed.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/attachments#get-jira-attachment-settings
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-attachments/#api-rest-api-3-attachment-meta-get
func (a *AttachmentService) Settings(ctx context.Context) (result *models2.AttachmentSettingScheme,
	response *ResponseScheme, err error) {

	var endpoint = "rest/api/3/attachment/meta"
	request, err := a.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = a.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Metadata returns the metadata for an attachment. Note that the attachment itself is not returned.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/attachments#get-attachment-metadata
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-attachments/#api-rest-api-3-attachment-id-get
func (a *AttachmentService) Metadata(ctx context.Context, attachmentID string) (result *models2.AttachmentMetadataScheme,
	response *ResponseScheme, err error) {

	if len(attachmentID) == 0 {
		return nil, nil, models2.ErrNoAttachmentIDError
	}

	var endpoint = fmt.Sprintf("rest/api/3/attachment/%v", attachmentID)
	request, err := a.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = a.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Delete deletes an attachment from an issue.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/attachments#delete-attachment
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-attachments/#api-rest-api-3-attachment-id-delete
func (a *AttachmentService) Delete(ctx context.Context, attachmentID string) (response *ResponseScheme, err error) {

	if len(attachmentID) == 0 {
		return nil, models2.ErrNoAttachmentIDError
	}

	var endpoint = fmt.Sprintf("rest/api/3/attachment/%v", attachmentID)
	request, err := a.client.newRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return
	}

	response, err = a.client.call(request, nil)
	if err != nil {
		return
	}

	return
}

// Human returns the metadata for the contents of an attachment, if it is an archive, and metadata for the attachment itself.
// For example, if the attachment is a ZIP archive, then information about the files in the archive is returned and metadata for the ZIP archive.
// Currently, only the ZIP archive format is supported.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/attachments#get-all-metadata-for-an-expanded-attachment
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-attachments/#api-rest-api-3-attachment-id-expand-human-get
// NOTE: Experimental Endpoint
func (a *AttachmentService) Human(ctx context.Context, attachmentID string) (result *models2.AttachmentHumanMetadataScheme,
	response *ResponseScheme, err error) {

	if len(attachmentID) == 0 {
		return nil, nil, models2.ErrNoAttachmentIDError
	}

	var endpoint = fmt.Sprintf("rest/api/3/attachment/%v/expand/human", attachmentID)
	request, err := a.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = a.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Add adds one attachment to an issue. Attachments are posted as multipart/form-data (RFC 1867).
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/attachments#add-attachment
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-attachments/#api-rest-api-3-issue-issueidorkey-attachments-post
func (a *AttachmentService) Add(ctx context.Context, issueKeyOrID, fileName string, file io.Reader) (result []*models2.AttachmentScheme, response *ResponseScheme, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, nil, models2.ErrNoIssueKeyOrIDError
	}

	if len(fileName) == 0 {
		return nil, nil, models2.ErrNoAttachmentNameError
	}

	if file == nil {
		return nil, nil, models2.ErrNoReaderError
	}

	var (
		endpoint         = fmt.Sprintf("rest/api/3/issue/%v/attachments", issueKeyOrID)
		body             = &bytes.Buffer{}
		attachmentWriter = multipart.NewWriter(body)
	)

	// create the attachment form row
	part, _ := attachmentWriter.CreateFormFile("file", fileName)

	// add the attachment metadata
	_, err = io.Copy(part, file)
	if err != nil {
		return nil, nil, err
	}

	attachmentWriter.Close()

	request, err := a.client.newRequest(ctx, http.MethodPost, endpoint, body)
	if err != nil {
		return
	}

	request.Header.Add("Content-Type", attachmentWriter.FormDataContentType())
	request.Header.Add("Accept", "application/json")
	request.Header.Set("X-Atlassian-Token", "no-check")

	response, err = a.client.call(request, &result)
	if err != nil {
		return
	}

	return
}
