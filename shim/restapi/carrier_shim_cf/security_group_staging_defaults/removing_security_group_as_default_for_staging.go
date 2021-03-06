// Code generated by go-swagger; DO NOT EDIT.

package security_group_staging_defaults

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// RemovingSecurityGroupAsDefaultForStagingHandlerFunc turns a function with the right signature into a removing security group as default for staging handler
type RemovingSecurityGroupAsDefaultForStagingHandlerFunc func(RemovingSecurityGroupAsDefaultForStagingParams) middleware.Responder

// Handle executing the request and returning a response
func (fn RemovingSecurityGroupAsDefaultForStagingHandlerFunc) Handle(params RemovingSecurityGroupAsDefaultForStagingParams) middleware.Responder {
	return fn(params)
}

// RemovingSecurityGroupAsDefaultForStagingHandler interface for that can handle valid removing security group as default for staging params
type RemovingSecurityGroupAsDefaultForStagingHandler interface {
	Handle(RemovingSecurityGroupAsDefaultForStagingParams) middleware.Responder
}

// NewRemovingSecurityGroupAsDefaultForStaging creates a new http.Handler for the removing security group as default for staging operation
func NewRemovingSecurityGroupAsDefaultForStaging(ctx *middleware.Context, handler RemovingSecurityGroupAsDefaultForStagingHandler) *RemovingSecurityGroupAsDefaultForStaging {
	return &RemovingSecurityGroupAsDefaultForStaging{Context: ctx, Handler: handler}
}

/*RemovingSecurityGroupAsDefaultForStaging swagger:route DELETE /config/staging_security_groups/{guid} securityGroupStagingDefaults removingSecurityGroupAsDefaultForStaging

Removing a Security Group as a default for staging

curl --insecure -i %s/v2/config/staging_security_groups/{guid} -X DELETE -H 'Authorization: %s'

*/
type RemovingSecurityGroupAsDefaultForStaging struct {
	Context *middleware.Context
	Handler RemovingSecurityGroupAsDefaultForStagingHandler
}

func (o *RemovingSecurityGroupAsDefaultForStaging) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewRemovingSecurityGroupAsDefaultForStagingParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
