// Code generated by go-swagger; DO NOT EDIT.

package device

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetDeviceNameAddrOKCode is the HTTP code returned for type GetDeviceNameAddrOK
const GetDeviceNameAddrOKCode int = 200

/*
GetDeviceNameAddrOK Success

swagger:response getDeviceNameAddrOK
*/
type GetDeviceNameAddrOK struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewGetDeviceNameAddrOK creates GetDeviceNameAddrOK with default headers values
func NewGetDeviceNameAddrOK() *GetDeviceNameAddrOK {

	return &GetDeviceNameAddrOK{}
}

// WithPayload adds the payload to the get device name addr o k response
func (o *GetDeviceNameAddrOK) WithPayload(payload string) *GetDeviceNameAddrOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get device name addr o k response
func (o *GetDeviceNameAddrOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDeviceNameAddrOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetDeviceNameAddrNotFoundCode is the HTTP code returned for type GetDeviceNameAddrNotFound
const GetDeviceNameAddrNotFoundCode int = 404

/*
GetDeviceNameAddrNotFound Not Found

swagger:response getDeviceNameAddrNotFound
*/
type GetDeviceNameAddrNotFound struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewGetDeviceNameAddrNotFound creates GetDeviceNameAddrNotFound with default headers values
func NewGetDeviceNameAddrNotFound() *GetDeviceNameAddrNotFound {

	return &GetDeviceNameAddrNotFound{}
}

// WithPayload adds the payload to the get device name addr not found response
func (o *GetDeviceNameAddrNotFound) WithPayload(payload string) *GetDeviceNameAddrNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get device name addr not found response
func (o *GetDeviceNameAddrNotFound) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDeviceNameAddrNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}