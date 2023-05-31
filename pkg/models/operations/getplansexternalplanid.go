// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Orb/pkg/models/shared"
	"net/http"
)

type GetPlansExternalPlanIDRequest struct {
	Plan           *shared.Plan `request:"mediaType=application/json"`
	ExternalPlanID string       `pathParam:"style=simple,explode=false,name=external_plan_id"`
}

type GetPlansExternalPlanIDResponse struct {
	ContentType string
	// OK
	Plan        *shared.Plan
	StatusCode  int
	RawResponse *http.Response
}
