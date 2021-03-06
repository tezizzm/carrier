// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// RetrievingOrganizationMemoryUsageHandlerFunc turns a function with the right signature into a retrieving organization memory usage handler
type RetrievingOrganizationMemoryUsageHandlerFunc func(RetrievingOrganizationMemoryUsageParams) middleware.Responder

// Handle executing the request and returning a response
func (fn RetrievingOrganizationMemoryUsageHandlerFunc) Handle(params RetrievingOrganizationMemoryUsageParams) middleware.Responder {
	return fn(params)
}

// RetrievingOrganizationMemoryUsageHandler interface for that can handle valid retrieving organization memory usage params
type RetrievingOrganizationMemoryUsageHandler interface {
	Handle(RetrievingOrganizationMemoryUsageParams) middleware.Responder
}

// NewRetrievingOrganizationMemoryUsage creates a new http.Handler for the retrieving organization memory usage operation
func NewRetrievingOrganizationMemoryUsage(ctx *middleware.Context, handler RetrievingOrganizationMemoryUsageHandler) *RetrievingOrganizationMemoryUsage {
	return &RetrievingOrganizationMemoryUsage{Context: ctx, Handler: handler}
}

/*RetrievingOrganizationMemoryUsage swagger:route GET /organizations/{guid}/memory_usage organizations retrievingOrganizationMemoryUsage

Retrieving organization memory usage

curl --insecure -i %s/v2/organizations/{guid}/memory_usage -X GET -H 'Authorization: %s'

*/
type RetrievingOrganizationMemoryUsage struct {
	Context *middleware.Context
	Handler RetrievingOrganizationMemoryUsageHandler
}

func (o *RetrievingOrganizationMemoryUsage) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewRetrievingOrganizationMemoryUsageParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
