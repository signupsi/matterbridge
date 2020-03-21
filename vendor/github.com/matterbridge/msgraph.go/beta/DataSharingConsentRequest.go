// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// DataSharingConsentRequestBuilder is request builder for DataSharingConsent
type DataSharingConsentRequestBuilder struct{ BaseRequestBuilder }

// Request returns DataSharingConsentRequest
func (b *DataSharingConsentRequestBuilder) Request() *DataSharingConsentRequest {
	return &DataSharingConsentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DataSharingConsentRequest is request for DataSharingConsent
type DataSharingConsentRequest struct{ BaseRequest }

// Get performs GET request for DataSharingConsent
func (r *DataSharingConsentRequest) Get(ctx context.Context) (resObj *DataSharingConsent, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DataSharingConsent
func (r *DataSharingConsentRequest) Update(ctx context.Context, reqObj *DataSharingConsent) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DataSharingConsent
func (r *DataSharingConsentRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}