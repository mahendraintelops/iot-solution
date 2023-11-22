package models

type Device struct {
	Id int64 `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	NoOfWheels int16 `json:"noOfWheels,omitempty"`

	Type string `json:"type,omitempty"`
}
