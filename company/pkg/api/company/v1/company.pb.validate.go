// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: company.proto

package company

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// define the regex for a UUID once up-front
var _company_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on GetCompanyRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetCompanyRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetCompanyRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetCompanyRequestMultiError, or nil if none found.
func (m *GetCompanyRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetCompanyRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if err := m._validateUuid(m.GetCompanyId()); err != nil {
		err = GetCompanyRequestValidationError{
			field:  "CompanyId",
			reason: "value must be a valid UUID",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetCompanyRequestMultiError(errors)
	}

	return nil
}

func (m *GetCompanyRequest) _validateUuid(uuid string) error {
	if matched := _company_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// GetCompanyRequestMultiError is an error wrapping multiple validation errors
// returned by GetCompanyRequest.ValidateAll() if the designated constraints
// aren't met.
type GetCompanyRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetCompanyRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetCompanyRequestMultiError) AllErrors() []error { return m }

// GetCompanyRequestValidationError is the validation error returned by
// GetCompanyRequest.Validate if the designated constraints aren't met.
type GetCompanyRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetCompanyRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetCompanyRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetCompanyRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetCompanyRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetCompanyRequestValidationError) ErrorName() string {
	return "GetCompanyRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetCompanyRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetCompanyRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetCompanyRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetCompanyRequestValidationError{}

// Validate checks the field values on GetCompanyResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetCompanyResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetCompanyResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetCompanyResponseMultiError, or nil if none found.
func (m *GetCompanyResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetCompanyResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetCompany()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetCompanyResponseValidationError{
					field:  "Company",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetCompanyResponseValidationError{
					field:  "Company",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCompany()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetCompanyResponseValidationError{
				field:  "Company",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetCompanyResponseMultiError(errors)
	}

	return nil
}

// GetCompanyResponseMultiError is an error wrapping multiple validation errors
// returned by GetCompanyResponse.ValidateAll() if the designated constraints
// aren't met.
type GetCompanyResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetCompanyResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetCompanyResponseMultiError) AllErrors() []error { return m }

// GetCompanyResponseValidationError is the validation error returned by
// GetCompanyResponse.Validate if the designated constraints aren't met.
type GetCompanyResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetCompanyResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetCompanyResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetCompanyResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetCompanyResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetCompanyResponseValidationError) ErrorName() string {
	return "GetCompanyResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetCompanyResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetCompanyResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetCompanyResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetCompanyResponseValidationError{}

// Validate checks the field values on PatchCompanyRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *PatchCompanyRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PatchCompanyRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PatchCompanyRequestMultiError, or nil if none found.
func (m *PatchCompanyRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *PatchCompanyRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetCompany() == nil {
		err := PatchCompanyRequestValidationError{
			field:  "Company",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetCompany()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, PatchCompanyRequestValidationError{
					field:  "Company",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, PatchCompanyRequestValidationError{
					field:  "Company",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCompany()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PatchCompanyRequestValidationError{
				field:  "Company",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return PatchCompanyRequestMultiError(errors)
	}

	return nil
}

// PatchCompanyRequestMultiError is an error wrapping multiple validation
// errors returned by PatchCompanyRequest.ValidateAll() if the designated
// constraints aren't met.
type PatchCompanyRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PatchCompanyRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PatchCompanyRequestMultiError) AllErrors() []error { return m }

// PatchCompanyRequestValidationError is the validation error returned by
// PatchCompanyRequest.Validate if the designated constraints aren't met.
type PatchCompanyRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PatchCompanyRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PatchCompanyRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PatchCompanyRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PatchCompanyRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PatchCompanyRequestValidationError) ErrorName() string {
	return "PatchCompanyRequestValidationError"
}

// Error satisfies the builtin error interface
func (e PatchCompanyRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPatchCompanyRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PatchCompanyRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PatchCompanyRequestValidationError{}

// Validate checks the field values on DeleteCompanyRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteCompanyRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteCompanyRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteCompanyRequestMultiError, or nil if none found.
func (m *DeleteCompanyRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteCompanyRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if err := m._validateUuid(m.GetCompanyId()); err != nil {
		err = DeleteCompanyRequestValidationError{
			field:  "CompanyId",
			reason: "value must be a valid UUID",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return DeleteCompanyRequestMultiError(errors)
	}

	return nil
}

func (m *DeleteCompanyRequest) _validateUuid(uuid string) error {
	if matched := _company_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// DeleteCompanyRequestMultiError is an error wrapping multiple validation
// errors returned by DeleteCompanyRequest.ValidateAll() if the designated
// constraints aren't met.
type DeleteCompanyRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteCompanyRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteCompanyRequestMultiError) AllErrors() []error { return m }

// DeleteCompanyRequestValidationError is the validation error returned by
// DeleteCompanyRequest.Validate if the designated constraints aren't met.
type DeleteCompanyRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteCompanyRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteCompanyRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteCompanyRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteCompanyRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteCompanyRequestValidationError) ErrorName() string {
	return "DeleteCompanyRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteCompanyRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteCompanyRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteCompanyRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteCompanyRequestValidationError{}

// Validate checks the field values on CreateCompanyRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateCompanyRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateCompanyRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateCompanyRequestMultiError, or nil if none found.
func (m *CreateCompanyRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateCompanyRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetName()); l < 1 || l > 15 {
		err := CreateCompanyRequestValidationError{
			field:  "Name",
			reason: "value length must be between 1 and 15 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetDescription()) > 3000 {
		err := CreateCompanyRequestValidationError{
			field:  "Description",
			reason: "value length must be at most 3000 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetAmountOfEmployees() <= 0 {
		err := CreateCompanyRequestValidationError{
			field:  "AmountOfEmployees",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Registered

	if _, ok := CompanyType_name[int32(m.GetType())]; !ok {
		err := CreateCompanyRequestValidationError{
			field:  "Type",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return CreateCompanyRequestMultiError(errors)
	}

	return nil
}

// CreateCompanyRequestMultiError is an error wrapping multiple validation
// errors returned by CreateCompanyRequest.ValidateAll() if the designated
// constraints aren't met.
type CreateCompanyRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateCompanyRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateCompanyRequestMultiError) AllErrors() []error { return m }

// CreateCompanyRequestValidationError is the validation error returned by
// CreateCompanyRequest.Validate if the designated constraints aren't met.
type CreateCompanyRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateCompanyRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateCompanyRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateCompanyRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateCompanyRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateCompanyRequestValidationError) ErrorName() string {
	return "CreateCompanyRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateCompanyRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateCompanyRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateCompanyRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateCompanyRequestValidationError{}

// Validate checks the field values on CreateCompanyResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateCompanyResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateCompanyResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateCompanyResponseMultiError, or nil if none found.
func (m *CreateCompanyResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateCompanyResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetCompany()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateCompanyResponseValidationError{
					field:  "Company",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateCompanyResponseValidationError{
					field:  "Company",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCompany()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateCompanyResponseValidationError{
				field:  "Company",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateCompanyResponseMultiError(errors)
	}

	return nil
}

// CreateCompanyResponseMultiError is an error wrapping multiple validation
// errors returned by CreateCompanyResponse.ValidateAll() if the designated
// constraints aren't met.
type CreateCompanyResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateCompanyResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateCompanyResponseMultiError) AllErrors() []error { return m }

// CreateCompanyResponseValidationError is the validation error returned by
// CreateCompanyResponse.Validate if the designated constraints aren't met.
type CreateCompanyResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateCompanyResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateCompanyResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateCompanyResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateCompanyResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateCompanyResponseValidationError) ErrorName() string {
	return "CreateCompanyResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateCompanyResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateCompanyResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateCompanyResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateCompanyResponseValidationError{}

// Validate checks the field values on Company with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Company) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Company with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in CompanyMultiError, or nil if none found.
func (m *Company) ValidateAll() error {
	return m.validate(true)
}

func (m *Company) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if err := m._validateUuid(m.GetId()); err != nil {
		err = CompanyValidationError{
			field:  "Id",
			reason: "value must be a valid UUID",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetName()); l < 1 || l > 15 {
		err := CompanyValidationError{
			field:  "Name",
			reason: "value length must be between 1 and 15 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetDescription()) > 3000 {
		err := CompanyValidationError{
			field:  "Description",
			reason: "value length must be at most 3000 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetAmountOfEmployees() <= 0 {
		err := CompanyValidationError{
			field:  "AmountOfEmployees",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Registered

	if _, ok := CompanyType_name[int32(m.GetType())]; !ok {
		err := CompanyValidationError{
			field:  "Type",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return CompanyMultiError(errors)
	}

	return nil
}

func (m *Company) _validateUuid(uuid string) error {
	if matched := _company_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// CompanyMultiError is an error wrapping multiple validation errors returned
// by Company.ValidateAll() if the designated constraints aren't met.
type CompanyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CompanyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CompanyMultiError) AllErrors() []error { return m }

// CompanyValidationError is the validation error returned by Company.Validate
// if the designated constraints aren't met.
type CompanyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CompanyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CompanyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CompanyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CompanyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CompanyValidationError) ErrorName() string { return "CompanyValidationError" }

// Error satisfies the builtin error interface
func (e CompanyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCompany.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CompanyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CompanyValidationError{}
