// Code generated by go-swagger; DO NOT EDIT.

package spaces

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"
)

// AssociateAuditorWithSpaceURL generates an URL for the associate auditor with space operation
type AssociateAuditorWithSpaceURL struct {
	AuditorGUID string
	GUID        string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *AssociateAuditorWithSpaceURL) WithBasePath(bp string) *AssociateAuditorWithSpaceURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *AssociateAuditorWithSpaceURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *AssociateAuditorWithSpaceURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/spaces/{guid}/auditors/{auditor_guid}"

	auditorGUID := o.AuditorGUID
	if auditorGUID != "" {
		_path = strings.Replace(_path, "{auditor_guid}", auditorGUID, -1)
	} else {
		return nil, errors.New("auditorGuid is required on AssociateAuditorWithSpaceURL")
	}

	guid := o.GUID
	if guid != "" {
		_path = strings.Replace(_path, "{guid}", guid, -1)
	} else {
		return nil, errors.New("guid is required on AssociateAuditorWithSpaceURL")
	}

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/v2"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *AssociateAuditorWithSpaceURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *AssociateAuditorWithSpaceURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *AssociateAuditorWithSpaceURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on AssociateAuditorWithSpaceURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on AssociateAuditorWithSpaceURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *AssociateAuditorWithSpaceURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
