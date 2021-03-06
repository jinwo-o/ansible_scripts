// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetSampleColumnsOKCode is the HTTP code returned for type GetSampleColumnsOK
const GetSampleColumnsOKCode int = 200

/*GetSampleColumnsOK columns

swagger:response getSampleColumnsOK
*/
type GetSampleColumnsOK struct {

	/*
	  In: Body
	*/
	Payload [][]string `json:"body,omitempty"`
}

// NewGetSampleColumnsOK creates GetSampleColumnsOK with default headers values
func NewGetSampleColumnsOK() *GetSampleColumnsOK {
	return &GetSampleColumnsOK{}
}

// WithPayload adds the payload to the get sample columns o k response
func (o *GetSampleColumnsOK) WithPayload(payload [][]string) *GetSampleColumnsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get sample columns o k response
func (o *GetSampleColumnsOK) SetPayload(payload [][]string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSampleColumnsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make([][]string, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// GetSampleColumnsBadRequestCode is the HTTP code returned for type GetSampleColumnsBadRequest
const GetSampleColumnsBadRequestCode int = 400

/*GetSampleColumnsBadRequest bad input parameter

swagger:response getSampleColumnsBadRequest
*/
type GetSampleColumnsBadRequest struct {
}

// NewGetSampleColumnsBadRequest creates GetSampleColumnsBadRequest with default headers values
func NewGetSampleColumnsBadRequest() *GetSampleColumnsBadRequest {
	return &GetSampleColumnsBadRequest{}
}

// WriteResponse to the client
func (o *GetSampleColumnsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}
