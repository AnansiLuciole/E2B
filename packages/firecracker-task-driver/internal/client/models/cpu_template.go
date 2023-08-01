// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// CPUTemplate The CPU Template defines a set of flags to be disabled from the microvm so that the features exposed to the guest are the same as in the selected instance type. Works only on Intel.
//
// swagger:model CpuTemplate
type CPUTemplate string

func NewCPUTemplate(value CPUTemplate) *CPUTemplate {
	return &value
}

// Pointer returns a pointer to a freshly-allocated CPUTemplate.
func (m CPUTemplate) Pointer() *CPUTemplate {
	return &m
}

const (

	// CPUTemplateC3 captures enum value "C3"
	CPUTemplateC3 CPUTemplate = "C3"

	// CPUTemplateT2 captures enum value "T2"
	CPUTemplateT2 CPUTemplate = "T2"

	// CPUTemplateNone captures enum value "None"
	CPUTemplateNone CPUTemplate = "None"
)

// for schema
var cpuTemplateEnum []interface{}

func init() {
	var res []CPUTemplate
	if err := json.Unmarshal([]byte(`["C3","T2","None"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cpuTemplateEnum = append(cpuTemplateEnum, v)
	}
}

func (m CPUTemplate) validateCPUTemplateEnum(path, location string, value CPUTemplate) error {
	if err := validate.EnumCase(path, location, value, cpuTemplateEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this Cpu template
func (m CPUTemplate) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateCPUTemplateEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this Cpu template based on context it is used
func (m CPUTemplate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}