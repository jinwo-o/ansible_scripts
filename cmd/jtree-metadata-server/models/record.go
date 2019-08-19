package models

import "github.com/go-openapi/swag"

//Record is the super struct
type Record struct {

	// patient
	Patient
	// sample
	Sample
	// experiment
	Experiment
	// result
	Result
	// resultdetails
	Resultdetails

}

// MarshalBinary interface implementation
func (m *Record) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Record) UnmarshalBinary(b []byte) error {
	var res Record
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
