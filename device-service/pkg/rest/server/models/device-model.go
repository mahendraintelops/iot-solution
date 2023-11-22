package models

type Device struct {
	Id int64 `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	Type string `json:"type,omitempty"`
}
