// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// AndroidForWorkEnrollmentProfileRequestBuilder is request builder for AndroidForWorkEnrollmentProfile
type AndroidForWorkEnrollmentProfileRequestBuilder struct{ BaseRequestBuilder }

// Request returns AndroidForWorkEnrollmentProfileRequest
func (b *AndroidForWorkEnrollmentProfileRequestBuilder) Request() *AndroidForWorkEnrollmentProfileRequest {
	return &AndroidForWorkEnrollmentProfileRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AndroidForWorkEnrollmentProfileRequest is request for AndroidForWorkEnrollmentProfile
type AndroidForWorkEnrollmentProfileRequest struct{ BaseRequest }

// Get performs GET request for AndroidForWorkEnrollmentProfile
func (r *AndroidForWorkEnrollmentProfileRequest) Get(ctx context.Context) (resObj *AndroidForWorkEnrollmentProfile, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AndroidForWorkEnrollmentProfile
func (r *AndroidForWorkEnrollmentProfileRequest) Update(ctx context.Context, reqObj *AndroidForWorkEnrollmentProfile) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AndroidForWorkEnrollmentProfile
func (r *AndroidForWorkEnrollmentProfileRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}