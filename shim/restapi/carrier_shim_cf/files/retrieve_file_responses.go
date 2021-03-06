// Code generated by go-swagger; DO NOT EDIT.

package files

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// RetrieveFileFoundCode is the HTTP code returned for type RetrieveFileFound
const RetrieveFileFoundCode int = 302

/*RetrieveFileFound successful response

swagger:response retrieveFileFound
*/
type RetrieveFileFound struct {
}

// NewRetrieveFileFound creates RetrieveFileFound with default headers values
func NewRetrieveFileFound() *RetrieveFileFound {

	return &RetrieveFileFound{}
}

// WriteResponse to the client
func (o *RetrieveFileFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(302)
}
