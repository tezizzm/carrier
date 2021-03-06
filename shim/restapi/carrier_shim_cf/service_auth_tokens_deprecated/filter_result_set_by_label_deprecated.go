// Code generated by go-swagger; DO NOT EDIT.

package service_auth_tokens_deprecated

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// FilterResultSetByLabelDeprecatedHandlerFunc turns a function with the right signature into a filter result set by label deprecated handler
type FilterResultSetByLabelDeprecatedHandlerFunc func(FilterResultSetByLabelDeprecatedParams) middleware.Responder

// Handle executing the request and returning a response
func (fn FilterResultSetByLabelDeprecatedHandlerFunc) Handle(params FilterResultSetByLabelDeprecatedParams) middleware.Responder {
	return fn(params)
}

// FilterResultSetByLabelDeprecatedHandler interface for that can handle valid filter result set by label deprecated params
type FilterResultSetByLabelDeprecatedHandler interface {
	Handle(FilterResultSetByLabelDeprecatedParams) middleware.Responder
}

// NewFilterResultSetByLabelDeprecated creates a new http.Handler for the filter result set by label deprecated operation
func NewFilterResultSetByLabelDeprecated(ctx *middleware.Context, handler FilterResultSetByLabelDeprecatedHandler) *FilterResultSetByLabelDeprecated {
	return &FilterResultSetByLabelDeprecated{Context: ctx, Handler: handler}
}

/*FilterResultSetByLabelDeprecated swagger:route GET /service_auth_tokens serviceAuthTokensDeprecated filterResultSetByLabelDeprecated

Filtering the result set by label (deprecated)

curl --insecure -i %s/v2/service_auth_tokens -X GET -H 'Authorization: %s'

*/
type FilterResultSetByLabelDeprecated struct {
	Context *middleware.Context
	Handler FilterResultSetByLabelDeprecatedHandler
}

func (o *FilterResultSetByLabelDeprecated) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewFilterResultSetByLabelDeprecatedParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
