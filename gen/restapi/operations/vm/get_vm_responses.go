// Code generated by go-swagger; DO NOT EDIT.

package vm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/girikuncoro/godc/gen/models"
)

// GetVMOKCode is the HTTP code returned for type GetVMOK
const GetVMOKCode int = 200

/*GetVMOK Successful operation

swagger:response getVmOK
*/
type GetVMOK struct {

	/*
	  In: Body
	*/
	Payload *models.VM `json:"body,omitempty"`
}

// NewGetVMOK creates GetVMOK with default headers values
func NewGetVMOK() *GetVMOK {

	return &GetVMOK{}
}

// WithPayload adds the payload to the get Vm o k response
func (o *GetVMOK) WithPayload(payload *models.VM) *GetVMOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get Vm o k response
func (o *GetVMOK) SetPayload(payload *models.VM) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetVMOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}