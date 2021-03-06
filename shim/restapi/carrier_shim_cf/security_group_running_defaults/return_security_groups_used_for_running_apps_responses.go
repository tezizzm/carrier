// Code generated by go-swagger; DO NOT EDIT.

package security_group_running_defaults

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/suse/carrier/shim/models"
)

// ReturnSecurityGroupsUsedForRunningAppsOKCode is the HTTP code returned for type ReturnSecurityGroupsUsedForRunningAppsOK
const ReturnSecurityGroupsUsedForRunningAppsOKCode int = 200

/*ReturnSecurityGroupsUsedForRunningAppsOK successful response

swagger:response returnSecurityGroupsUsedForRunningAppsOK
*/
type ReturnSecurityGroupsUsedForRunningAppsOK struct {

	/*
	  In: Body
	*/
	Payload *models.ReturnSecurityGroupsUsedForRunningAppsResponsePaged `json:"body,omitempty"`
}

// NewReturnSecurityGroupsUsedForRunningAppsOK creates ReturnSecurityGroupsUsedForRunningAppsOK with default headers values
func NewReturnSecurityGroupsUsedForRunningAppsOK() *ReturnSecurityGroupsUsedForRunningAppsOK {

	return &ReturnSecurityGroupsUsedForRunningAppsOK{}
}

// WithPayload adds the payload to the return security groups used for running apps o k response
func (o *ReturnSecurityGroupsUsedForRunningAppsOK) WithPayload(payload *models.ReturnSecurityGroupsUsedForRunningAppsResponsePaged) *ReturnSecurityGroupsUsedForRunningAppsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the return security groups used for running apps o k response
func (o *ReturnSecurityGroupsUsedForRunningAppsOK) SetPayload(payload *models.ReturnSecurityGroupsUsedForRunningAppsResponsePaged) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReturnSecurityGroupsUsedForRunningAppsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
