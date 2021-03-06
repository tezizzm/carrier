// Code generated by go-swagger; DO NOT EDIT.

package security_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// RemoveSpaceFromSecurityGroupHandlerFunc turns a function with the right signature into a remove space from security group handler
type RemoveSpaceFromSecurityGroupHandlerFunc func(RemoveSpaceFromSecurityGroupParams) middleware.Responder

// Handle executing the request and returning a response
func (fn RemoveSpaceFromSecurityGroupHandlerFunc) Handle(params RemoveSpaceFromSecurityGroupParams) middleware.Responder {
	return fn(params)
}

// RemoveSpaceFromSecurityGroupHandler interface for that can handle valid remove space from security group params
type RemoveSpaceFromSecurityGroupHandler interface {
	Handle(RemoveSpaceFromSecurityGroupParams) middleware.Responder
}

// NewRemoveSpaceFromSecurityGroup creates a new http.Handler for the remove space from security group operation
func NewRemoveSpaceFromSecurityGroup(ctx *middleware.Context, handler RemoveSpaceFromSecurityGroupHandler) *RemoveSpaceFromSecurityGroup {
	return &RemoveSpaceFromSecurityGroup{Context: ctx, Handler: handler}
}

/*RemoveSpaceFromSecurityGroup swagger:route DELETE /security_groups/{guid}/spaces/{space_guid} securityGroups removeSpaceFromSecurityGroup

Remove Space from the Security Group

curl --insecure -i %s/v2/security_groups/{guid}/spaces/{space_guid} -X DELETE -H 'Authorization: %s'

*/
type RemoveSpaceFromSecurityGroup struct {
	Context *middleware.Context
	Handler RemoveSpaceFromSecurityGroupHandler
}

func (o *RemoveSpaceFromSecurityGroup) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewRemoveSpaceFromSecurityGroupParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
