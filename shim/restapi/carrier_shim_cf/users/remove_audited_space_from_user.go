// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// RemoveAuditedSpaceFromUserHandlerFunc turns a function with the right signature into a remove audited space from user handler
type RemoveAuditedSpaceFromUserHandlerFunc func(RemoveAuditedSpaceFromUserParams) middleware.Responder

// Handle executing the request and returning a response
func (fn RemoveAuditedSpaceFromUserHandlerFunc) Handle(params RemoveAuditedSpaceFromUserParams) middleware.Responder {
	return fn(params)
}

// RemoveAuditedSpaceFromUserHandler interface for that can handle valid remove audited space from user params
type RemoveAuditedSpaceFromUserHandler interface {
	Handle(RemoveAuditedSpaceFromUserParams) middleware.Responder
}

// NewRemoveAuditedSpaceFromUser creates a new http.Handler for the remove audited space from user operation
func NewRemoveAuditedSpaceFromUser(ctx *middleware.Context, handler RemoveAuditedSpaceFromUserHandler) *RemoveAuditedSpaceFromUser {
	return &RemoveAuditedSpaceFromUser{Context: ctx, Handler: handler}
}

/*RemoveAuditedSpaceFromUser swagger:route DELETE /users/{guid}/audited_spaces/{audited_space_guid} users removeAuditedSpaceFromUser

Remove Audited Space from the User

curl --insecure -i %s/v2/users/{guid}/audited_spaces/{audited_space_guid} -X DELETE -H 'Authorization: %s'

*/
type RemoveAuditedSpaceFromUser struct {
	Context *middleware.Context
	Handler RemoveAuditedSpaceFromUserHandler
}

func (o *RemoveAuditedSpaceFromUser) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewRemoveAuditedSpaceFromUserParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
