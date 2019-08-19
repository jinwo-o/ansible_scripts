// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// AddExperimentOKCode is the HTTP code returned for type AddExperimentOK
const AddExperimentOKCode int = 200

/*AddExperimentOK id

swagger:response addExperimentOK
*/
type AddExperimentOK struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewAddExperimentOK creates AddExperimentOK with default headers values
func NewAddExperimentOK() *AddExperimentOK {
	return &AddExperimentOK{}
}

// WithPayload adds the payload to the add experiment o k response
func (o *AddExperimentOK) WithPayload(payload string) *AddExperimentOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add experiment o k response
func (o *AddExperimentOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddExperimentOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// AddExperimentBadRequestCode is the HTTP code returned for type AddExperimentBadRequest
const AddExperimentBadRequestCode int = 400

/*AddExperimentBadRequest invalid input, object invalid

swagger:response addExperimentBadRequest
*/
type AddExperimentBadRequest struct {
}

// NewAddExperimentBadRequest creates AddExperimentBadRequest with default headers values
func NewAddExperimentBadRequest() *AddExperimentBadRequest {
	return &AddExperimentBadRequest{}
}

// WriteResponse to the client
func (o *AddExperimentBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// AddExperimentConflictCode is the HTTP code returned for type AddExperimentConflict
const AddExperimentConflictCode int = 409

/*AddExperimentConflict an existing item already exists

swagger:response addExperimentConflict
*/
type AddExperimentConflict struct {
}

// NewAddExperimentConflict creates AddExperimentConflict with default headers values
func NewAddExperimentConflict() *AddExperimentConflict {
	return &AddExperimentConflict{}
}

// WriteResponse to the client
func (o *AddExperimentConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(409)
}
