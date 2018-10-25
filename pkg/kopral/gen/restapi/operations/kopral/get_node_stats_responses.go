// Code generated by go-swagger; DO NOT EDIT.

package kopral

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/girikuncoro/godc/pkg/kopral/gen/models"
)

// GetNodeStatsOKCode is the HTTP code returned for type GetNodeStatsOK
const GetNodeStatsOKCode int = 200

/*GetNodeStatsOK Successful operation

swagger:response getNodeStatsOK
*/
type GetNodeStatsOK struct {

	/*
	  In: Body
	*/
	Payload *models.NodeStats `json:"body,omitempty"`
}

// NewGetNodeStatsOK creates GetNodeStatsOK with default headers values
func NewGetNodeStatsOK() *GetNodeStatsOK {

	return &GetNodeStatsOK{}
}

// WithPayload adds the payload to the get node stats o k response
func (o *GetNodeStatsOK) WithPayload(payload *models.NodeStats) *GetNodeStatsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get node stats o k response
func (o *GetNodeStatsOK) SetPayload(payload *models.NodeStats) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetNodeStatsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetNodeStatsDefault Unexpected Error

swagger:response getNodeStatsDefault
*/
type GetNodeStatsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetNodeStatsDefault creates GetNodeStatsDefault with default headers values
func NewGetNodeStatsDefault(code int) *GetNodeStatsDefault {
	if code <= 0 {
		code = 500
	}

	return &GetNodeStatsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get node stats default response
func (o *GetNodeStatsDefault) WithStatusCode(code int) *GetNodeStatsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get node stats default response
func (o *GetNodeStatsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get node stats default response
func (o *GetNodeStatsDefault) WithPayload(payload *models.Error) *GetNodeStatsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get node stats default response
func (o *GetNodeStatsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetNodeStatsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
